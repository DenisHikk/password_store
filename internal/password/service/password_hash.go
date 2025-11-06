package password

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func generateSalt(n int) ([]byte, error) {
	s := make([]byte, n)
	_, err := rand.Read(s)
	return s, err
}

func EncodeHashPassword(password string) (string, error) {
	salt, err := generateSalt(16)
	if err != nil {
		return "", err
	}

	timeParam := uint32(3)      // count iteration (time)
	memKiB := uint32(32 * 1024) // 32 MiB
	threads := uint8(3)         // threads for calculate
	keyLen := uint32(32)        // length hash

	derived := argon2.IDKey([]byte(password), salt, timeParam, memKiB, threads, keyLen)

	bSalt := base64.RawStdEncoding.EncodeToString(salt)
	bHash := base64.RawStdEncoding.EncodeToString(derived)

	//example output string "$argon2i$v=19$m=16,t=2,p=1$MTIzMTIzMTIz$gOKio+D/siwKq8z8+/Y2EQ"
	argon2string := fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		memKiB, timeParam, threads, bSalt, bHash,
	)

	return argon2string, nil
}
