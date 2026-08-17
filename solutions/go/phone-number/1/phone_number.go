package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	var builder strings.Builder
	err := errors.New("invalid phone number")

	builder.Grow(11)

	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			builder.WriteRune(r)
		}
	}

	str := builder.String()

	if len(str) < 10 || len(str) > 11 {
		return "", err
	}

	if len(str) == 11 {
		if str[0] == '1' {
			str = str[1:]
		} else {
			return "", err
		}
	}

	if str[0] < '2' || str[0] > '9' || str[3] < '2' || str[3] > '9' {
		return "", err
	}

	return str, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return phoneNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%v) %v-%v", phoneNumber[0:3], phoneNumber[3:6], phoneNumber[6:]), nil
}
