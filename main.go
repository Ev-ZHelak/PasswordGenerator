package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"unicode/utf8"
)

const (
	MinPasswordLength = 4
	MinPasswordsCount = 1
	MaxPasswordsCount = 50
)

var (
	ErrPasswordLengthTooLow = errors.New("password length too low")
	ErrTooLowPasswordsCount = errors.New("too low passwords count")
	ErrTooBigPasswordsCount = errors.New("too big passwords count")
)

var (
	upperChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerChars   = []rune("abcdefghijklmnopqrstuvwxyz")
	digitChars   = []rune("0123456789")
	specialChars = []rune("!@#$%^&*")
)

func main() {
	slicePasswords, err := generatePassword(42, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range slicePasswords {
			fmt.Println(v)
		}
		
	}

}

func generatePassword(length int, count int) ([]string, error) {
	switch {
	case length < 4:
		return nil, ErrPasswordLengthTooLow
	case count < MinPasswordsCount:
		return nil, ErrTooLowPasswordsCount
	case count > MaxPasswordsCount:
		return nil, ErrTooBigPasswordsCount
	}

	charSets := [][]rune{
		upperChars,
		lowerChars,
		digitChars,
		specialChars,
	}

	passwordsMap := map[string]struct{}{}

	for len(passwordsMap) != count {
		resPass := ""

		basket := map[int64]struct{}{}
		randomN := make([]int64, 0, len(charSets))

		for len(randomN) != len(charSets) {
			n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSets))))
			if _, ok := basket[n.Int64()]; !ok {
				basket[n.Int64()] = struct{}{}
				randomN = append(randomN, n.Int64())
			}
		}

		for _, i := range randomN {
			index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSets[i]))))
			resPass += string(charSets[i][index.Int64()])
		}

		for utf8.RuneCountInString(resPass) != length {
			indexX, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSets))))
			charX := charSets[indexX.Int64()]
			index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charX))))
			resPass += string(charX[index.Int64()])
		}
		passwordsMap[resPass] = struct{}{}
	}

	resultPasswords := make([]string, 0, count)
	for k := range passwordsMap {
		resultPasswords = append(resultPasswords, k)
	}
	return resultPasswords, nil
}
