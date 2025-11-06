package password

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"

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

	timeParam := uint32(3)     // count iteration (time)
	memKB := uint32(32 * 1024) // 64 MB
	threads := uint8(3)        // threads for calculate
	keyLen := uint32(32)       // length hash

	derived := argon2.IDKey([]byte(password), salt, timeParam, memKB, threads, keyLen)

	bSalt := base64.RawStdEncoding.EncodeToString(salt)
	bHash := base64.RawStdEncoding.EncodeToString(derived)

	//"argon2id$t=19,m=%d,p=%d$%s$%s"
	argon2string := "argon2id$v=19" +
		"m=" + strconv.Itoa(int(memKB)) +
		",t=" + strconv.Itoa(int(timeParam)) +
		",p=" + strconv.Itoa(int(threads)) +
		"$" + bSalt +
		"$" + bHash

	return argon2string, nil
}

func CheckHashPassword() {

}
