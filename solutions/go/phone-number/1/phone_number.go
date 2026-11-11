package phonenumber

import (
	"errors"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	phone := strings.Map(func(r rune) rune {
		if strings.ContainsRune("()+-. ", r) {
			return -1
		}
		return r
	}, phoneNumber)
	switch len(phone) {
	case 10:

	case 11:
		if phone[0] != '1' {
			return "", errors.New("More than 10 digits must begin with country code 1")
		}
		phone = phone[1:]
	default:
		return "", errors.New("incorrect number of digits")
	}
	for _, r := range phone {
		if r < '0' || r > '9' {
			return "", errors.New("letters not permitted")
		}
	}
	if phone[0] < '2' || phone[0] > '9' {
		return "", errors.New("First Verification Digit out of bound")
	}
	if phone[3] < '2' || phone[3] > '9' {
		return "", errors.New("Second Verification Digit out of bound")
	}
	// fmt.Printf("phone number = %v ", phone)
	return phone, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phone, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phone[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phone, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + phone[:3] + ") " + phone[3:6] + "-" + phone[6:], nil
}
