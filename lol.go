package lol

import (
	cryptorand "crypto/rand"
	"io"
	"math"
	"math/big"
)

const (
	suchIntExponent = int64(1 << 6)
	primalityRounds = 1 << 11
)

var suchInt = new(big.Int).Exp(big.NewInt(math.MaxInt64), big.NewInt(suchIntExponent), nil)

// Woot returns a random *big.Int.
func Woot(rand io.Reader) (*big.Int, bool, error) {
	i, err := cryptorand.Int(rand, suchInt)
	if err != nil {
		return nil, false, err
	}

	return i, i.ProbablyPrime(primalityRounds), nil
}
