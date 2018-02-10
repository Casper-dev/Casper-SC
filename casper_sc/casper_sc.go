package Casper_SC

import (
	"encoding/hex"
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
	"github.com/Casper-dev/Casper-SC/casper"
)

// with no 0x
const ContractAddress = "B36482457fE521fAcd46AAB5FF860A99D16Da95d"

// use your key with no 0x
const PrivateKey = ""

type Config struct {
	ipv4Addr string

}

func GetConfig() *Config {
	return nil
}

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
	///Using ropsten infura addr
	client, err := ethclient.Dial("https://ropsten.infura.io/**********")

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
	///TODO: gas oracle sometimes fails too hard; for now let's assume all ops are under this gas
	auth.GasLimit = uint64(500000)
	return casperSClient, client, auth
}

func GetSC() (*casper.Casper, *ethclient.Client, *bind.TransactOpts) {
	if casperClient == nil {
		casperClient, client, auth = initSC()
	}
	return casperClient, client, auth
}
func Usage_example() {

	casperSClient, client, auth := GetSC()
	fmt.Printf("%d\n", auth.GasPrice, client)

	fmt.Printf("wallet %s\n", auth.From.String())

	tx_get, err := casperSClient.GetPeers(nil, big.NewInt(int64(1337)))

	tx_msg, errm := hex.DecodeString(tx_get)
	if errm != nil {
		fmt.Println(err)
	}
	fmt.Printf("Got %s\n", string(tx_get))
	fmt.Printf("From hex %s\n", string(tx_msg))
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

	var tx *types.Transaction
	for i := 0; i < 1; i++ {
		tx, err = casperSClient.RemoveProviderMachine(auth, "/ip4/192.168./tcp/4001/ipfs/QmcQsNc79XWhDrHM")
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
			time.Sleep(500 * time.Millisecond)
		}
	}

	//_, err = greeterClient.RemoveIP(auth, ip.String())
	tx_get, err = casperSClient.GetPeers(nil, big.NewInt(int64(1337)))
	fmt.Printf("Got %s\n", string(tx_get))

}

func Crypto_Example() {
	inputFile := "aesexampleinput"
	outputFile := "aesexampleoutput"
	message := []byte("Hello, AES!")
	ioutil.WriteFile(inputFile, message, os.ModePerm)

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
	fmt.Println(string(todecrypt))

}
