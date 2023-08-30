package mini

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strconv"
)

func CheckSignature(signature, timestamp, nonce, token string) bool {

	tempArr := []string{token, timestamp, nonce}
	sort.Strings(tempArr)
	tempStr := strconv.Quote(implode(tempArr))

	hash := sha1.New()
	hash.Write([]byte(tempStr))
	hashed := hash.Sum(nil)
	hashedStr := hex.EncodeToString(hashed)

	if hashedStr == signature {
		return true
	} else {
		return false
	}

}

func implode(arr []string) string {
	var result string
	for _, str := range arr {
		result += str
	}
	return result
}
