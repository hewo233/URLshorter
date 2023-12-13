package shortener

import (
	"crypto/sha256"
	"github.com/itchyny/base58-go"
	"math/big"
	"strconv"
)

func sha256Of(s string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(s))
	return algorithm.Sum(nil)
}

func base58Encode(b []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, errEncode := encoding.Encode(b)
	if errEncode != nil {
		panic(errEncode)
	}
	return string(encoded)
}

func GenerateShortURL(originUrl string, userId string) string {
	urlHashBytes := sha256Of(originUrl + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalShortURL := base58Encode([]byte(strconv.FormatUint(generatedNumber, 10)))
	return finalShortURL[:8]
}
