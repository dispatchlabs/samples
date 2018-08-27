package util

import (
	"strings"
	"crypto/md5"
	"encoding/hex"
)

func GetUniqueStrings(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	var result []string

	for _, v := range elements {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, strings.TrimSpace(v))
		}
	}
	return result
}

func GetMD5(value string) string {
	hasher := md5.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}