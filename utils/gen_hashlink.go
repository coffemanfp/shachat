package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// GenHashlink - Gen a identity hash based in a prefix, a ID and the current time
func GenHashlink(id int) (hash string) {
	time := time.Now().Unix()
	hash = fmt.Sprintf(
		"%x",
		sha256.Sum256(
			[]byte(
				fmt.Sprintf(
					"%d - %d",
					id,
					time,
				),
			),
		),
	)
	return
}
