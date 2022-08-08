package chain

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"
)

// ChecksumToSarafuAddress returns a lowercased 40 len hex address representation compatible with cic-stack database schemas
//
// 0x02b0DF387A3A68AA3134668752dd82bE70B7dE1C -> 02b0df387a3a68aa3134668752dd82be70b7de1c
func ChecksumToSarafuAddress(address string) (string, error) {
	l := len(address)

	if l < 40 || l > 42 {
		return "", fmt.Errorf("%s is not a valid eth address", address)
	}

	if len(address) == 42 {
		return strings.ToLower(address)[2:], nil
	}

	return strings.ToLower(address), nil
}

// SarafuAddressToChecksum returns a checksumed address from a lowercased sarafu address representations ready to be used with w3.A()
//
// 02b0df387a3a68aa3134668752dd82be70b7de1c -> 0x02b0DF387A3A68AA3134668752dd82bE70B7dE1C
func SarafuAddressToChecksum(address string) string {
	address = strings.ToLower(address)
	address = strings.Replace(address, "0x", "", 1)

	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(address))
	hash := sha.Sum(nil)
	hashstr := hex.EncodeToString(hash)
	result := []string{"0x"}
	for i, v := range address {
		res, _ := strconv.ParseInt(string(hashstr[i]), 16, 64)
		if res > 7 {
			result = append(result, strings.ToUpper(string(v)))
			continue
		}
		result = append(result, string(v))
	}

	return strings.Join(result, "")
}
