package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

func Generate(str string) (string, error) {
	_h := sha256.New()
	_, err := _h.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(_h.Sum(nil)), nil
}
