package utils

import (
	"crypto"
	"errors"
	"regexp"
)

// ValidateHash - Validates a hash.
func ValidateHash(hash string, algorithm crypto.Hash) (valid bool, err error) {
	if hash == "" || algorithm == 0 {
		return
	}

	switch algorithm {
	case crypto.SHA256:
		valid, _ = regexp.MatchString(`[a-fA-F0-9]{64}`, hash)
	default:
		err = errors.New("Hash algorithm is not supported by ValidatesHash()")
	}

	return
}
