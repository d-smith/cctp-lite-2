package cmd

import (
	"crypto/ecdsa"
	"database/sql"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
	"github.com/holiman/uint256"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"

	"github.com/mattn/go-sqlite3"
)

// runAttestorCmd represents the runAttestor command
var runAttestorCmd = &cobra.Command{
	Use:   "runAttestor",
	Short: "Run the attestor API",
	Long:  `Run the attestor API`,
	Run:   api,
}

var attestorPrivateKey *ecdsa.PrivateKey
var db *sql.DB

func init() {
	rootCmd.AddCommand(runAttestorCmd)

	privateKeyFromEnv := os.Getenv("ETH_ATTESTOR_KEY")
	if privateKeyFromEnv == "" {
		log.Fatal("ETH_ATTESTOR_KEY not set")
	}

	if privateKeyFromEnv[:2] == "0x" {
		privateKeyFromEnv = privateKeyFromEnv[2:]
	}

	var err error
	attestorPrivateKey, err = crypto.HexToECDSA(privateKeyFromEnv)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing DB...")
	v, _, _ := sqlite3.Version()
	log.Println("Opening sqlite with driver version", v)

	db, err = sql.Open("sqlite3", "attestor.db")
	if err != nil {
		log.Fatal(err)
	}
}

func api(cmd *cobra.Command, args []string) {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/attestor/attest", storeAttestation).Methods("POST")
	n := negroni.Classic()
	n.UseHandler(r)

	err := http.ListenAndServe(":3010", n)
	if err != nil {
		log.Fatalln("Error starting server", err)
	}

}

func storeAttestation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called storeAttestation")

	defer r.Body.Close()

	message, err := io.ReadAll(r.Body)
	if err != nil {
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
	}

	parsedMessage, err := parseMessageSent(message)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	printMessage(parsedMessage)

	signature, err := signMessage(message)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	res, err := db.Exec("INSERT INTO attestations (nonce, sender, receiver, source_domain, dest_domain, amount, message, signature)	 VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		parsedMessage.Nonce, parsedMessage.Sender, parsedMessage.Recipient,
		parsedMessage.LocalDomain, parsedMessage.RemoteDomain,
		parsedMessage.BurnMessage.Amount,
		hexutil.Encode(message), hexutil.Encode(signature))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		log.Println("Error getting last insert id", err)
	} else {
		log.Println("Inserted row with id", int(id))
	}
}

func signMessage(message []byte) ([]byte, error) {
	msgHash := crypto.Keccak256Hash(message)
	stamp := []byte("\x19Ethereum Signed Message:\n32")
	signature, err := crypto.Sign(crypto.Keccak256Hash(stamp, msgHash.Bytes()).Bytes(), attestorPrivateKey)
	if err != nil {
		return nil, err
	}

	if signature[crypto.RecoveryIDOffset] == 0 || signature[crypto.RecoveryIDOffset] == 1 {
		signature[crypto.RecoveryIDOffset] += 27
	}
	return signature, nil
}

type MessageSent struct {
	// local domain - uint32
	// remote domain - uint32
	// nonce - uint64
	// sender - bytes32
	// recipient - bytes32
	// burn message - bytes
	MessageVersion uint32
	LocalDomain    uint32
	RemoteDomain   uint32
	Nonce          uint64
	Sender         string
	Recipient      string
	BurnMessage    *BurnMessage
}

func parseMessageSent(messageSent []byte) (*MessageSent, error) {
	if len(messageSent) < 84 {
		return nil, fmt.Errorf("invalid message sent length")
	}

	messageVersion := binary.BigEndian.Uint32(messageSent[0:4])
	localDomain := binary.BigEndian.Uint32(messageSent[4:8])
	remoteDomain := binary.BigEndian.Uint32(messageSent[8:12])
	nonce := binary.BigEndian.Uint64(messageSent[12:20])

	sender := messageSent[20:52]
	senderHex := hexutil.Encode(sender[12:32])

	recipient := messageSent[52:84]
	recipientHex := hexutil.Encode(recipient[12:32])

	burnMessage := messageSent[84:]

	parsedBurnMessage, err := parseBurnMessage(burnMessage)
	if err != nil {
		return nil, err
	}

	return &MessageSent{
		MessageVersion: messageVersion,
		LocalDomain:    localDomain,
		RemoteDomain:   remoteDomain,
		Nonce:          nonce,
		Sender:         senderHex,
		Recipient:      recipientHex,
		BurnMessage:    parsedBurnMessage,
	}, nil
}

type BurnMessage struct {
	Version       uint32
	BurnToken     string
	MintRecipient string
	Amount        *uint256.Int
	Sender        string
}

func parseBurnMessage(burnMessage []byte) (*BurnMessage, error) {
	if len(burnMessage) < 132 {
		return nil, fmt.Errorf("invalid burn message length")
	}

	burnMessageVersion := binary.BigEndian.Uint32(burnMessage[0:4])
	burnToken := burnMessage[4:36]
	mintRecipient := burnMessage[36:68]

	amountBytes := burnMessage[68:100]
	hexAmount := hexutil.Encode(amountBytes[12:32])

	// Convert hexAmount to a uint256

	i := 0
	if hexAmount[0:2] == "0x" {
		i = 2
	}

	for ; i < len(hexAmount); i++ {
		if hexAmount[i] != '0' {
			break
		}
	}

	if i == len(hexAmount) {
		hexAmount = "0x0"
	} else {
		hexAmount = fmt.Sprintf("0x%s", hexAmount[i:])
	}

	amountDec, err := uint256.FromHex(hexAmount)
	if err != nil {
		return nil, err
	}

	sender := burnMessage[100:132]

	return &BurnMessage{
		Version:       burnMessageVersion,
		BurnToken:     hexutil.Encode(burnToken[12:32]),
		MintRecipient: hexutil.Encode(mintRecipient[12:32]),
		Amount:        amountDec,
		Sender:        hexutil.Encode(sender[12:32]),
	}, nil
}

func printMessage(m *MessageSent) {
	fmt.Printf("Message fields\n")
	fmt.Printf("Message Version: %d\n", m.MessageVersion)
	fmt.Printf("Local domain: %d\n", m.LocalDomain)
	fmt.Printf("Remote domain: %d\n", m.RemoteDomain)
	fmt.Printf("Nonce: %d\n", m.Nonce)
	fmt.Printf("Sender: %s\n", m.Sender)
	fmt.Printf("Recipient: %s\n", m.Recipient)
	fmt.Printf("Burn message:\n")
	fmt.Printf("  Version: %d\n", m.BurnMessage.Version)
	fmt.Printf("  Burn token: %s\n", m.BurnMessage.BurnToken)
	fmt.Printf("  Mint recipient: %s\n", m.BurnMessage.MintRecipient)
	fmt.Printf("  Amount: %d\n", m.BurnMessage.Amount)
	fmt.Printf("  Sender: %s\n", m.BurnMessage.Sender)
}
