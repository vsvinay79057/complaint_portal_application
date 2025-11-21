package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func GenerateSecretCode() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	ts := time.Now().UnixNano()
	hexPart := hex.EncodeToString(b)

	code := strings.ToUpper(hexPart) + "-" + fmt.Sprint(ts%1000000)
	return code
}
