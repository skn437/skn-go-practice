package components

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GetCryptoRandomNumber() {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Printf("Crypto Random Number: %v \n", randomNumber)
}
