package utils_test

import (
	"testing"
	"crypto"
	"crypto/sha256"
	"fmt"
	"github.com/coffemanfp/shachat/utils"
)

func TestValidateHash(t *testing.T) {
	// Using a i index for the checksums.
	for i := 0; i < 100; i++ {
		sum := fmt.Sprintf("%x", sha256.Sum256([]byte(string(i))))
		valid, err := utils.ValidateHash(sum, crypto.SHA256)
		if err != nil {
			t.Errorf("unexpected error:\n%s", err)
		}

		if !valid {
			t.Errorf("valid value (%t) invalid, expected (%t)", valid, true)
		}
	}
}