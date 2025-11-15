package password

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Params struct {
	timeParam uint32
	memKiB    uint32
	threads   uint8
	bSalt     string
	bHash     string
	keyLen    int
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

func VerifyHashPassword(hashedPassword, inputPassword string) (bool, error) {
	params, err := parseArgon2Params(hashedPassword)
	if err != nil {
		return false, err
	}

	saltBytes, err := base64.RawStdEncoding.DecodeString(params.bSalt)
	if err != nil {
		log.Println("I cant decode")
		return false, err
	}
	hashedInput := argon2.IDKey([]byte(inputPassword), saltBytes, params.timeParam, params.memKiB, params.threads, 32)
	check := base64.RawStdEncoding.EncodeToString(hashedInput) == params.bHash
	if !check {
		return !check, errors.New("Wrong")
	}
	return check, nil
}

func generateSalt(n int) ([]byte, error) {
	s := make([]byte, n)
	_, err := rand.Read(s)
	return s, err
}

func parseArgon2Params(hash string) (Params, error) {
	// parts[1] -> "argon2id"
	// parts[2] -> "v=19"
	// parts[3] -> "m=32768,t=3,p=3"
	// parts[4] -> salt
	// parts[5] -> hashed password
	// where
	// m = MemKiB | t = timeParam | p = threads | v = igonore
	var params Params
	parts := strings.Split(hash, "$")
	params.bSalt = parts[4]
	params.bHash = parts[5]
	params.keyLen = len(params.bHash)
	keyValuePairs := strings.Split(parts[3], ",")

	for _, value := range keyValuePairs {
		kv := strings.Split(value, "=")
		if len(kv) != 2 {
			return Params{}, errors.New("bigger than need")
		}
		switch kv[0] {
		case "m":
			memKiB, err := strconv.Atoi(kv[1])
			if err != nil {
				return Params{}, errors.New("error parsing memory (m) parameter")
			}
			params.memKiB = uint32(memKiB)
		case "t":
			iter, err := strconv.Atoi(kv[1])
			if err != nil {
				return Params{}, errors.New("error parsing time (t) parameter")
			}
			params.timeParam = uint32(iter)
		case "p":
			threads, err := strconv.Atoi(kv[1])
			if err != nil {
				return Params{}, errors.New("error parsing threads (p) parameter")
			}
			params.threads = uint8(threads)
		}
	}

	return params, nil
}
