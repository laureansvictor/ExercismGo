package isbnverifier

func IsValidISBN(isbn string) bool {
	var digits []int

	for i := 0; i < len(isbn); i++ {
		c := isbn[i]

		if c == '-' || c == ' ' {
			continue
		}

		var digit int
		if c == 'X' {
			if i != len(isbn)-1 {
				return false
			}
			digit = 10
		} else if c >= '0' && c <= '9' {
			digit = int(c - '0')
		} else {
			return false
		}

		digits = append(digits, digit)
	}

	if len(digits) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += digits[i] * (10 - i)
	}

	return sum%11 == 0
}
