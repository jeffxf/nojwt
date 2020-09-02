package nojwt

// Heavily influenced by @FiloSottile (Filippo Valsorda)'s tweet about JWS:
// https://twitter.com/FiloSottile/status/1288964453065797632

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	// ErrInvalidSignature is returned if the signature of the tokin is invalid
	ErrInvalidSignature = errors.New("Invalid signature")
)

type tokenFormat struct {
	Data      []byte `json:"data"`
	Signature []byte `json:"signature"`
}

// Encode creates a signed token from a private key and interface containing data
func Encode(privateKey []byte, data interface{}) (tokenString string, err error) {
	var token tokenFormat
	token.Data, err = json.Marshal(data)
	if err != nil {
		return "", err
	}

	token.Signature = ed25519.Sign(privateKey, token.Data)

	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	return string(tokenBytes), nil
}

// Decode will verify a token and write the data to the provided interface
func Decode(publicKey []byte, token string, v interface{}) error {
	var t tokenFormat

	err := json.Unmarshal([]byte(token), &t)
	if err != nil {
		return err
	}

	verified := ed25519.Verify(publicKey, t.Data, t.Signature)
	if !verified {
		return fmt.Errorf("%w: %v (public key: %v)", ErrInvalidSignature, token, publicKey)
	}

	err = json.Unmarshal(t.Data, v)
	if err != nil {
		return err
	}
	return nil
}
