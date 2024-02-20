package service

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/starshine-bcit/bby-buohub/auth/util"
	"golang.org/x/crypto/argon2"
)

// https://snyk.io/blog/secure-password-hashing-in-go/

type Argon2idHash struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
	saltLen uint32
}

type HashSalt struct {
	Hash []byte
	Salt []byte
}

func NewArgon2idHash(time, saltLen uint32, memory uint32, threads uint8, keyLen uint32) *Argon2idHash {
	return &Argon2idHash{
		time:    time,
		saltLen: saltLen,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

func NewDefaultArgon2idHash() *Argon2idHash {
	return &Argon2idHash{
		time:    2,
		saltLen: 32,
		memory:  64 * 1024,
		threads: 16,
		keyLen:  256,
	}
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (a *Argon2idHash) GenerateHash(password, salt []byte) (*HashSalt, error) {
	var err error
	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}
	if err != nil {
		return nil, err
	}
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)
	return &HashSalt{Hash: hash, Salt: salt}, nil
}

func (a *Argon2idHash) Compare(hash, salt, password []byte) error {
	hashSalt, err := a.GenerateHash(password, salt)
	if err != nil {
		return err
	}
	if !bytes.Equal(hash, hashSalt.Hash) {
		return errors.New("received failed login attempts, hash doesn't match")
	}
	return nil
}

func (hs *HashSalt) Stringify() (string, string) {
	hashstr := base64.StdEncoding.EncodeToString(hs.Hash)
	saltstr := base64.StdEncoding.EncodeToString(hs.Salt)
	return hashstr, saltstr
}

func newHashSaltFromString(hash string, salt string) (*HashSalt, error) {
	hashbyte, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return nil, errors.New("error decoding hash string")
	}
	saltbyte, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return nil, errors.New("error decoding salt string")
	}
	return &HashSalt{Hash: hashbyte, Salt: saltbyte}, nil
}

func CompareHashFromDB(hash, salt, password string) bool {
	hashSalt, err := newHashSaltFromString(hash, salt)
	if err != nil {
		util.InfoLogger.Println("Error base64 decoding hash and/or salt")
		return false
	}
	passbyte := []byte(password)
	a2 := NewDefaultArgon2idHash()
	if err = a2.Compare(hashSalt.Hash, hashSalt.Salt, passbyte); err != nil {
		util.InfoLogger.Println(err.Error())
		return false
	}
	return true
}

func GenerateHashForDB(password string) (string, string, error) {
	a2 := NewDefaultArgon2idHash()
	hs, err := a2.GenerateHash([]byte(password), []byte{})
	if err != nil {
		util.ErrorLogger.Println(err.Error())
		return "", "", err
	}
	hash, salt := hs.Stringify()
	return hash, salt, nil
}
