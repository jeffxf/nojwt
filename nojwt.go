package nojwt

import (
	"crypto/ed25519"
	"encoding/json"
)

// Heavily influenced by @FiloSottile (Filippo Valsorda)'s tweet about JWS:
// https://twitter.com/FiloSottile/status/1288964453065797632

type tokenFormat struct {
	Data      []byte `json:"data"`
	Signature []byte `json:"signature"`
}

// CreateToken creates a token from a private key and interface containing data
func CreateToken(privateKey []byte, data interface{}) (tokenString string, err error) {
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

// VerifyToken will verify the signature of a token given a public key
func VerifyToken(publicKey []byte, token string) (bool, error) {
	var t tokenFormat
	err := json.Unmarshal([]byte(token), &t)
	if err != nil {
		return false, err
	}

	verified := ed25519.Verify(publicKey, t.Data, t.Signature)
	return verified, nil
}

// ReadToken will read the data from a token into the provided interface
func ReadToken(token string, v interface{}) error {
	var t tokenFormat
	err := json.Unmarshal([]byte(token), &t)
	if err != nil {
		return err
	}
	err = json.Unmarshal(t.Data, v)
	if err != nil {
		return err
	}
	return nil
}
