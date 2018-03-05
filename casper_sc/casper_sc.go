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
	"github.com/Casper-dev/Casper-SC/casper"

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
	"bufio"
	"strings"
	rand2 "math/rand"
)

// with no 0x 		   
const ContractAddress = "8F98d0170f06dbc7b2D577724765F03B96163D7f"

// with no 0x
const PrivateKey = "674393e0fb1cba8a71be3f1261e7171effc992bc5047ae0efe8b0e49e554e293"

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
	client, err := ethclient.Dial("http://94.130.182.144:8545")

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

	fmt.Println(client.PendingBalanceAt(context.Background(), auth.From))
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

func ValidateMineTX(txGet func() (tx *types.Transaction, err error), client *ethclient.Client, auth *bind.TransactOpts) (data string) {
	tx, err := txGet()
	n, _ := rand.Int(rand.Reader, big.NewInt(100000000000))
	rand2.Seed(n.Int64())
	for ; err != nil; {
		fmt.Println(err)
		nonce, _ := client.PendingNonceAt(context.Background(), auth.From)
		auth.Nonce = big.NewInt(int64(nonce + rand2.Uint64()%15))
		fmt.Println("nonce", auth.Nonce)
		tx, err = txGet()
		time.Sleep(time.Millisecond * time.Duration(100+rand2.Int()%3000))
	}
	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())
	data = MineTX(tx, client)
	return
}

func MineTX(tx *types.Transaction, client *ethclient.Client) (data string) {
	defer func() {
		if r:=recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Printf("Gas %d\nGas price %d", tx.Gas(), tx.GasPrice())
	for ; ; {
		rxt, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			println(err)
		} else {
			println("Waiting for TX to mine")
		}
		if (!pending) {
			fmt.Println("Waiting a second for the receipt")
			time.Sleep(1 * time.Second)
			fmt.Println("getting receipt")
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			fmt.Println(err)
			if err == nil {
				println(rxt.String())
				println("receipt", receipt.Status, receipt.String())
				if len(receipt.Logs) > 0 {
					for _, log := range receipt.Logs {
						data += string(log.Data)
					}
					fmt.Println("data", data)
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

	tx_get, err := casperSClient.GetPeers(nil, big.NewInt(int64(13353457)))
	fmt.Println(tx_get)
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
	//var tx *types.Transaction
	fmt.Println("getting peers: ")
	fmt.Println(casperSClient.GetPeers(nil, big.NewInt(213434)))

	fmt.Println(casperSClient.VerifyReplication(nil, "/ip4/192.168.120.55/tcp/4001/ipfs/QmRFXdsPm6d6yr3izZEBK7oqZhPvmGUr5V3XnAKM2J8V8s"))
	removePeersClosure := func(ip string) func() (*types.Transaction, error) {
		//return casperSClient.RemoveAllPeers(auth)
		return func() (*types.Transaction, error) {
			return casperSClient.RemoveAllPeers(auth)
		}
	}
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("> Get peers for hash")
	sentenceo, err := buf.ReadString('\n')
	n, _ := casperSClient.GetStoringPeers(nil, sentenceo)
	fmt.Println("read ", sentenceo, n)

	fmt.Println(err)
	fmt.Println(removePeersClosure)
	fmt.Println("> remove all peers? (Y/N)")
	sentence, err := buf.ReadString('\n')
	fmt.Println("read ", sentence)
	if strings.ContainsAny(sentence, "yY") {
		for ; ; {
			tx, _ := casperSClient.GetPeers(nil, big.NewInt(34234))
			if tx.Ip1 == "" {
				break
			}
			fmt.Println("removing ", tx.Ip1)
			ValidateMineTX(removePeersClosure(tx.Ip1), client, auth)
			fmt.Println("removed ", tx.Ip1)
		}
	}
	fmt.Println("getting peers: ")
	fmt.Println(casperSClient.GetPeers(nil, big.NewInt(213434)))
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
