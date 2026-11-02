package luhn

import "strings"

func Valid(id string) bool {
	aux := strings.Fields(id)
	sum := 0
	posFromRight := 0

	for i := len(aux) - 1; i >= 0; i-- {
		for j := len(aux[i]) - 1; j >= 0; j-- {
			c := aux[i][j]

			if c < '0' || c > '9' {
				return false
			}

			digit := int(c - '0')

			if posFromRight%2 == 1 {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}

			sum += digit
			posFromRight++

		}
	}

	if posFromRight <= 1 {
		return false
	}

	return sum%10 == 0
}
