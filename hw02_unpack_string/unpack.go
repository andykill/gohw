package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (result string, err error) {
	err = ErrInvalidString
	if str == "" {
		return "", nil
	} else if unicode.IsDigit(rune(str[0])) {
		return
	}

	var builder strings.Builder
	lenStr := len(str)
	for i := 0; i < lenStr; {
		symbol := rune(str[i])

		symbolStr := string(str[i])
		if (i + 1) >= lenStr {
			builder.WriteString(symbolStr)
			break
		}
		nextSymbol := rune(str[i+1])
		if nextSymbol == symbol {
			return "", err
		} else if symbolStr == "\\" {
			i++
			builder.WriteString(string(nextSymbol))
		} else if unicode.IsDigit(nextSymbol) {
			var countString strings.Builder
			countString.WriteString(string(nextSymbol))
			j := i + 2
			for j < lenStr {
				if !unicode.IsDigit(rune(str[j])) {
					break
				}
				countString.WriteString(string(str[j]))
				j++
			}

			countInt, err := strconv.Atoi(countString.String())
			if err != nil {
				return "", err
			}
			strRepeat := strings.Repeat(symbolStr, countInt)
			builder.WriteString(strRepeat)
			i = j
			continue
		} else if symbol < 65 || (symbol > 90 && symbol < 97) || symbol > 122 {
			return "", err
		} else if !unicode.IsDigit(nextSymbol) {
			builder.WriteString(symbolStr)
		}
		i++
	}
	return builder.String(), nil
}
