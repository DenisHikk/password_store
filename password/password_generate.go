package password

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lower  = "abcdefghijklmnopqrstuvwxyz"
	upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digit  = "0123456789"
	symbol = "!@#$%^&*()-_=+[]{};:,.?/\\"
)

type PasswordConfig struct {
	Length    int
	UseLower  bool
	UseUpper  bool
	UseDigit  bool
	UseSymbol bool
}

func GeneratePassword(cfg PasswordConfig) (string, error) {
	if cfg.Length <= 0 {
		return "", fmt.Errorf("Invalid length")
	}

	pools := ""
	if cfg.UseLower {
		pools += lower
	}
	if cfg.UseUpper {
		pools += upper
	}
	if cfg.UseDigit {
		pools += digit
	}
	if cfg.UseSymbol {
		pools += symbol
	}

	out := make([]byte, cfg.Length)
	for i := range out {
		ch, err := pick(pools)
		if err != nil {
			return "", err
		}
		out[i] = ch
	}
	return string(out), nil
}

func pick(pools string) (byte, error) {
	m, err := rand.Int(rand.Reader, big.NewInt(int64(len(pools))))
	if err != nil {
		return 0, err
	}
	return pools[m.Int64()], nil
}
