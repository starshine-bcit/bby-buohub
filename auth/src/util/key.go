package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
)

var PrivKey *ecdsa.PrivateKey
var PubKey *ecdsa.PublicKey

func LoadKey() {
	ex, err := os.Executable()
	if err != nil {
		ErrorLogger.Fatalln(err)
	}
	keyPath := filepath.Join(filepath.Dir(filepath.Dir(ex)), "keys")
	privKeyPath := filepath.Join(keyPath, "authjwt")
	pubKeyPath := filepath.Join(keyPath, "authjwt.pub")
	privKeyData, err := os.ReadFile(privKeyPath)
	if err != nil {
		ErrorLogger.Fatalf("Could not open private key file: %v\n", privKeyPath)
	}
	pubKeyData, err := os.ReadFile(pubKeyPath)
	if err != nil {
		ErrorLogger.Fatalf("Could not open public key file: %v\n", pubKeyPath)
	}
	PrivKey, PubKey = decodeByesKey(privKeyData, pubKeyData)
}

func EncodeKey(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func DecodeKey(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

func decodeByesKey(pemEncoded []byte, pemEncodedPub []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := jwt.ParseECPrivateKeyFromPEM(pemEncoded)
	if err != nil {
		ErrorLogger.Fatalln("Failed to parse ECDSA private key from PEM")
	}
	publicKey, err := jwt.ParseECPublicKeyFromPEM(pemEncodedPub)
	if err != nil {
		ErrorLogger.Fatalln("Failed to parse ECDSA public key from PEM")
	}
	return privateKey, publicKey
}

func TestKey() {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	publicKey := &privateKey.PublicKey

	encPriv, encPub := EncodeKey(privateKey, publicKey)

	fmt.Println(encPriv)
	fmt.Println(encPub)

	priv2, pub2 := DecodeKey(encPriv, encPub)

	fmt.Println(priv2)
	fmt.Println(pub2)
}
