package main

import (
	"crypto/ecdsa"
	"database/sql"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
	"github.com/holiman/uint256"
	"github.com/mattn/go-sqlite3"
	"github.com/urfave/negroni"
)

var attestorPrivateKey *ecdsa.PrivateKey
var db *sql.DB

func init() {

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

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/attestor/attest", storeAttestation).Methods("POST")
	r.HandleFunc("/api/v1/attestor/receipts/{sourceDomain}/{recipient}", listReceipts).Methods("GET")
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
		Sender:         strings.ToLower(senderHex),
		Recipient:      strings.ToLower(recipientHex),
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
		MintRecipient: strings.ToLower(hexutil.Encode(mintRecipient[12:32])),
		Amount:        amountDec,
		Sender:        strings.ToLower(hexutil.Encode(sender[12:32])),
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

func listReceipts(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	sourceDomain := params["sourceDomain"]
	recipient := strings.ToLower(params["recipient"])

	if (sourceDomain == "") || (recipient == "") {
		http.Error(w, "sourceDomain and recipient must be specified", 400)
		return
	}

	sdVal, err := strconv.ParseUint(sourceDomain, 10, 32)
	if err != nil {
		http.Error(w, "sourceDomain must be a number", 400)
		return
	}

	rows, err := db.Query("SELECT nonce, sender, receiver, source_domain, dest_domain, amount, message, signature FROM attestations WHERE source_domain = ? AND receiver = ?", sdVal, recipient)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var nonce uint64
		var sender string
		var receiver string
		var sourceDomain uint32
		var destDomain uint32
		var amount uint64
		var message string
		var signature string

		err = rows.Scan(&nonce, &sender, &receiver, &sourceDomain, &destDomain, &amount, &message, &signature)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprintf(w, "%d %s %s %d %d %d %s %s\n", nonce, sender, receiver, sourceDomain, destDomain, amount, message, signature)
	}

}
