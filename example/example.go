package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/jeffxf/nojwt"
)

// Token describes the data within a token
type Token struct {
	Username      string `json:"username"`
	UserID        int    `json:"userid"`
	SomeOtherType interface{}
}

func main() {
	// Generate a new random key pair
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	// You can create a key pair externally and provide it instead which is what
	// you probably want to do. Otherwise, a new key will be genereated
	// everytime your app is restarted which will invalidate all of the
	// previously signed tokens

	//Set the fields of the token we want to create
	token := Token{
		Username:      "jeffxf",
		UserID:        501,
		SomeOtherType: map[string]string{"details": "nada"},
	}
	fmt.Printf("Token data: %+v\n\n", token)

	// Encode and sign the token with the private key
	signedToken, err := nojwt.Encode(privateKey, token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Signed Token: %s\n\n", signedToken)

	// Once the token has been sent to a client and returned back to you, any of
	// your services that have the public key can verify the signature of a
	// token is valid and read the data in one shot via the Decode function.
	// Let's extract the data into a new instance of our Claims struct
	var decodedToken Token
	err = nojwt.Decode(publicKey, signedToken, &decodedToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Print the username field from the decoded token as an example
	fmt.Printf("Username: %s\n\n", decodedToken.Username)
}
