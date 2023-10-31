package main

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
)

func main() {
	p := "0x0000000100000001000000020000000000000000000000000000000000000000835f0aa692b8ebcdea8e64742e5bce30303669c2000000000000000000000000798d4ba9baf0064ec19eb4f0a1a45785ae9d6dfc00000001000000000000000000000000c0a4b9e04fb55b1b498c634faeeb7c8dd5895b53000000000000000000000000798d4ba9baf0064ec19eb4f0a1a45785ae9d6dfc0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000835f0aa692b8ebcdea8e64742e5bce30303669c2"
	bin, err := hexutil.Decode(p)
	if err != nil {
		panic(err)
	}

	// version uint32 - uint32 is 4 bytes
	// local domain - uint32
	// remote domain - uint32
	// nonce - uint64
	// sender - bytes32
	// recipient - bytes32
	// burn message - bytes

	messageVersion := binary.BigEndian.Uint32(bin[0:4])
	fmt.Println(messageVersion)

	localDomain := binary.BigEndian.Uint32(bin[4:8])
	fmt.Println(localDomain)

	remoteDomain := binary.BigEndian.Uint32(bin[8:12])
	fmt.Println(remoteDomain)

	nonce := binary.BigEndian.Uint64(bin[12:20])
	fmt.Println(nonce)

	sender := bin[20:52]
	senderHex := hexutil.Encode(sender[12:32])
	fmt.Println(senderHex)

	recipient := bin[52:84]
	fmt.Println(hexutil.Encode(recipient[12:32]))

	burnMessage := bin[84:]

	// Now decompose burn message
	// version uint32
	// burnToken - bytes32
	// mintRecipient - bytes32
	// amount - uint256
	// sender - bytes32

	fmt.Printf("\nBurn message fields\n")

	burnMessageVersion := binary.BigEndian.Uint32(burnMessage[0:4])
	fmt.Println(burnMessageVersion)

	burnToken := burnMessage[4:36]
	fmt.Println(hexutil.Encode(burnToken[12:32]))

	mintRecipient := burnMessage[36:68]
	fmt.Println(hexutil.Encode(mintRecipient[12:32]))

	//amount := burnMessage[68:100]
	//hexAmount := hexutil.Encode(amount[12:32])
	//hexAmount := "0x0000000000000000000000000000000000000000"
	hexAmount := "0x0000000000000000000000000000000000000005"
	//hexAmount := "0x9999999999999999999999999999999999999999"
	fmt.Println("starting hex amount", hexAmount)

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

	fmt.Println("new hex amount", hexAmount)

	amountDec, err := uint256.FromHex(hexAmount)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(amountDec)
	}

	sender2 := burnMessage[100:132]
	fmt.Println(hexutil.Encode(sender2[12:32]))

}
