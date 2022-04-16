package jwt

import (
	"crypto/ecdsa"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func ParseSSHPrivateKey(key string) (*ecdsa.PrivateKey, error) {
	blob := []byte(key)
	rawKey, err := ssh.ParseRawPrivateKey(blob)
	if err != nil {
		return nil, err
	}

	privateKey, ok := rawKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse it to the EDSA private key struct")
	}

	return privateKey, nil
}
