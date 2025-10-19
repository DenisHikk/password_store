package password

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

type HashPassword struct {
	Hash      string
	Salt      string
	MemKB     uint32
	TimeParam uint32
	threads   uint8
	KeyLen    uint32
}

func generateSalt(n int) ([]byte, error) {
	s := make([]byte, n)
	_, err := rand.Read(s)
	return s, err
}

func EncodeHashPassword(password string) (HashPassword, error) {
	salt, err := generateSalt(16)
	if err != nil {
		return HashPassword{}, err
	}

	timeParam := uint32(3)     // ?
	memKB := uint32(64 * 1024) // 64 MB
	threads := uint8(4)        // threads for calculate
	keyLen := uint32(32)       // length hash

	derived := argon2.IDKey([]byte(password), salt, timeParam, memKB, threads, keyLen)

	bSalt := base64.RawStdEncoding.EncodeToString(salt)
	bHash := base64.RawStdEncoding.EncodeToString(derived)

	hashPassword := HashPassword{
		Hash:      bHash,
		Salt:      bSalt,
		MemKB:     memKB,
		TimeParam: timeParam,
		threads:   threads,
		KeyLen:    keyLen,
	}

	return hashPassword, nil
}

func CheckHashPassword()