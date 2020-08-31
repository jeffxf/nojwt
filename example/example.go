package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/jeffxf/nojwt"
)

// Claims describes the data we want to include in a token
type Claims struct {
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
	tokenClaims := Claims{
		Username:      "jeffxf",
		UserID:        501,
		SomeOtherType: map[string]string{"details": "nada"},
	}
	fmt.Printf("Token Claims: %+v\n\n", tokenClaims)

	// Create the token
	token, err := nojwt.CreateToken(privateKey, tokenClaims)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token: %v\n\n", token)

	// Once the token has been sent to a client and returned back to you, any of
	// your services that have the public key can verify the signature is valid
	verified, err := nojwt.VerifyToken(publicKey, token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Valid Signature: %t\n\n", verified)
	if !verified {
		return
	}

	// Now that we've validated the signature, let's extract the data from the
	// token into a new instance of our Claims struct
	var retrievedTokenData Claims
	err = nojwt.ReadToken(token, &retrievedTokenData)
	if err != nil {
		log.Fatal(err)
	}
	// Print the username field from the token as an example
	fmt.Printf("Username: %s\n\n", retrievedTokenData.Username)
}
