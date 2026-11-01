package hamming

import "errors"

func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	count := 0
	if len(a) != len(b) {
		return 0, errors.New("both sequences should have the same size")
	}
	for i, v := range ar {
		if br[i] != v {
			count++
		}
	}
	return count, nil
}
