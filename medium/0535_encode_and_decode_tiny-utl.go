package medium

import (
	"math/rand/v2"
	"strings"
)

type Codec struct {
	urls    map[string]string
	symbols string
}

func Constructor() Codec {
	return Codec{
		urls:    map[string]string{},
		symbols: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	randomCode := generateRandomCode(this.symbols)
	this.urls[randomCode] = longUrl
	return longUrl[:(strings.LastIndex(longUrl, "/") + 1)] + randomCode 
}

func generateRandomCode(code string) string {
	var builder strings.Builder
	builder.Grow(8)
	for range 8 {
		idx := rand.IntN(len(code) + 1)
		builder.WriteByte(code[idx])
	}
	return builder.String()
}

func (this *Codec) decode(shortUrl string) string {
    return this.urls[shortUrl[(strings.LastIndex(shortUrl, "/") + 1):]]
}
