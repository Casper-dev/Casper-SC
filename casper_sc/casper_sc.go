package Casper_SC

import (
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"


	//"time"
	//"context"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/casperDev1/Casper-SC/casper"
)

// with no 0x
const ContractAddress = "eb9100d37d8fa8A54608cf0eb20671a3bc802E80"

// with no 0x
const PrivateKey = "674393e0fb1cba8a71be3f1261e7171effb998bc5047ae0eee8b0e49e556e293"

func GetIPv4() (ip string) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if v.IP.To4() != nil && !v.IP.IsLoopback() {
					ip = v.IP.String()
					fmt.Printf("got ipv4: %s, %s\n", ip, addr.Network())
				}
			}
		}
	}
	return
}

var casperClient *casper.Casper = nil
var client *ethclient.Client
var auth *bind.TransactOpts

func initSC() (*casper.Casper, *ethclient.Client, *bind.TransactOpts) {

	/**
	 * Connecting to provider
	 */
	client, err := ethclient.Dial("https://ropsten.infura.io/WTu1LsZIdygEHpUcn2Nd")

	if err != nil {
		log.Fatal(err)
	}

	/**
	 * Connecting to contract at an address
	 */

	contractAddress := common.HexToAddress(ContractAddress)
	casperSClient, err := casper.NewCasper(contractAddress, client)

	key, err := crypto.HexToECDSA(PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(key)
	//auth.GasPrice = big.NewInt(250000000)
	auth.GasLimit = uint64(500000);
	return casperSClient, client, auth
}

func GetSC() (*casper.Casper, *ethclient.Client, *bind.TransactOpts) {
	if casperClient == nil {
		casperClient, client, auth = initSC()
	}
	return casperClient, client, auth
}

func ValidateMineTX(txGet func() (tx *types.Transaction, err error), client *ethclient.Client) (data string) {
	tx, err := txGet()
	for ; err != nil; {
		fmt.Println(err)
		tx, err = txGet()
		time.Sleep(time.Second * 2)
	}
	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())
	data = MineTX(tx, client)
	return
}

func MineTX(tx *types.Transaction, client *ethclient.Client) (data string) {
	fmt.Printf("Gas %d\nGas price %d", tx.Gas(), tx.GasPrice())
	for ; ; {
		rxt, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			println(err)
		} else {
			println("Waiting for TX to mine")
		}
		if (!pending) {
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			fmt.Println(err)
			if err == nil {
				println(rxt.String())
				println("receipt", receipt.Status, receipt.String())
				if len(receipt.Logs) > 0 {
					data = string(receipt.Logs[0].Data)
					fmt.Println(data)
				}
			}

			break
		}
		time.Sleep(2500 * time.Millisecond)

	}
	fmt.Println("Tx succesfully mined")
	return
}

func Usage_example() {

	casperSClient, client, auth := GetSC()
	fmt.Printf("%d\n", auth.GasPrice, client)

	fmt.Printf("wallet %s\n", auth.From.String())

	tx_get, err := casperSClient.GetPeers(nil, big.NewInt(int64(1337)))

	/*tx_msg, errm := hex.DecodeString(tx_get)
	if errm != nil {
		fmt.Println(err)
	}
	fmt.Printf("Got %s\n", string(tx_get))
	*/
	fmt.Printf("From hex %s\n", "Pl0x make GetNPeers ffs")
	/**
	 * Calling contract method
	 */

	/*ip := GetIPv4()


	tx, err := casperSClient.Register(auth, ip, big.NewInt(int64(1337)))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())


	if err != nil {
		log.Fatal(err)
	}
	for ;; {
		rxt, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			println(err)
		}else {
			println("Waiting for TX to mine")
		}
		if (!pending) {
			println(rxt.String())
			break
		}
		time.Sleep(500 * time.Millisecond)
	}*/

	auth.GasLimit = uint64(1500000)
	casperSClient.RemoveAllPeers(auth)
	var tx *types.Transaction
	for i := 0; i < 1; i++ {
		tx, err = casperSClient.RemoveProviderMachine(auth, "")
	}
	fmt.Println(err)
	if tx != nil {
		fmt.Printf("Gas price: %d\n", tx.GasPrice())
		for {
			rxt, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
			if err != nil {
				println(err)
			} else {
				println("Waiting for TX to mine")
			}
			if !pending {
				println(rxt.String())
				break
			}
			time.Sleep(2500 * time.Millisecond)
		}
	}

	//_, err = greeterClient.RemoveIP(auth, ip.String())
	tx_get, err = casperSClient.GetPeers(nil, big.NewInt(int64(1337)))
	fmt.Printf("Got %s\n", tx_get.Ip1)

}

func Crypto_Example() {
	inputFile := "aesexampleinput"
	outputFile := "aesexampleoutput"
	message := []byte("Hello, AES!")
	ioutil.WriteFile(inputFile, message, os.ModePerm)
	f, err := os.Create(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := f.Truncate(1e9); err != nil {
		log.Fatal(err)
	}

	toEncrypt, _ := ioutil.ReadFile(inputFile)

	key := make([]byte, 32)

	read, _ := rand.Read(key)
	if read != 32 {
		fmt.Printf("Read %d, expected 32 bytes", read)
	}
	///Encoding
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	encrypted := make([]byte, aes.BlockSize+len(toEncrypt))
	iv := encrypted[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(aesBlock, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], toEncrypt)
	ioutil.WriteFile(outputFile, encrypted, os.ModePerm)

	///Decoding
	newAESBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	todecrypt, err := ioutil.ReadFile(outputFile)
	if err != nil {
		fmt.Println(err)
	}

	if len(todecrypt) < aes.BlockSize {
		panic("ciphertext too short")
	}

	newiv := todecrypt[:aes.BlockSize]
	todecrypt = todecrypt[aes.BlockSize:]

	newstream := cipher.NewCFBDecrypter(newAESBlock, newiv)
	newstream.XORKeyStream(todecrypt, todecrypt)
	//fmt.Println(string(todecrypt))

}
