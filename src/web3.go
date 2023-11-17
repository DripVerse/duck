package src

import (
	"crypto/ecdsa"
	"log"
	"reflect"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

func wifToPrivateKey(wif string) (*ecdsa.PrivateKey, error) {
	decoded, err := base58.Decode(wif)
	if err != nil {
		return nil, err
	}

	// The first byte is the version byte, and the last 4 bytes are the checksum.
	// We skip the version byte and take all but the last 4 bytes.
	privateKeyBytes := decoded[1 : len(decoded)-4]

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func GetAddress(key string) {
	log.Println("key:", key)

	privateKey, err := wifToPrivateKey(key)

	// privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	log.Println("Address: %s\n", address, reflect.TypeOf(address))
}
