package brailleprinter

import (
	"math/rand"
	"time"
)

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genKey() string {
	keyLenth := 12
	rand.Seed(time.Now().UnixNano())

	var k int
	key := make([]byte, keyLenth)

	for i := 0; i < keyLenth; i++ {
		k = rand.Intn(len(keyChar))
		key[i] = keyChar[k]
	}

	return string(key[:])
}
