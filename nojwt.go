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

// Token describes the fields within a token
type Token struct {
	Data      []byte `json:"data"`
	Signature []byte `json:"signature"`
}

// Encode creates a signed token from a private key and interface containing data
func Encode(privateKey []byte, data interface{}) (token Token, err error) {
	token.Data, err = json.Marshal(data)
	if err != nil {
		return Token{}, err
	}
	token.Signature = ed25519.Sign(privateKey, token.Data)
	return token, nil
}

// EncodeToString calls the Encode function and returns the token as a string
func EncodeToString(privateKey []byte, data interface{}) (string, error) {
	token, err := Encode(privateKey, data)
	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return "", err
	}
	return string(tokenBytes), nil
}

// Decode will verify a token and write the data to the provided interface
func Decode(publicKey []byte, token Token, v interface{}) error {
	verified := ed25519.Verify(publicKey, token.Data, token.Signature)
	if !verified {
		return fmt.Errorf("%w: %v (public key: %v)", ErrInvalidSignature, token, publicKey)
	}

	err := json.Unmarshal(token.Data, v)
	if err != nil {
		return err
	}
	return nil
}

// DecodeFromString accepts a string token and calls the Decode function
func DecodeFromString(publicKey []byte, tokenString string, v interface{}) error {
	var token Token
	err := json.Unmarshal([]byte(tokenString), &token)
	if err != nil {
		return err
	}
	return Decode(publicKey, token, v)
}
