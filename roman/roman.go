package roman

import (
	"errors"
	"regexp"
	"strings"
)

func ParseToRoman(i int) (string, error) {
	r, err := NewRomanNumeralFromArabic(i)
	if err != nil {
		return "", err
	}
	return r.toRoman(), nil
}

func ParseToArabic(roman string) (int, error) {
	r, err := NewRomanNumeralFromRoman(roman)
	if err != nil {
		return 0, err
	}
	return r.toArabic()
}

type RomanNumeral struct {
	arabicVal int
	romanVal  string
}

var (
	ErrNumeralMustBePositive      = errors.New("value of RomanNumeral must be positive")
	ErrNumeralMustBe3999OrLess    = errors.New("value of RomanNumeral must be 3999 or less")
	ErrDoesNotDefineARomanNumeral = errors.New("an empty string does not define a Roman numeral")
	ErrNoNegativeRomanNumerals    = errors.New("negative numbers not allowed")
	ErrNumberFormatException      = errors.New("illegal character in Roman numeral")
	ErrInvalidRomanNumeral        = errors.New("invalid Roman numeral")
)

var (
	numbers = [...]int{1000, 900, 500, 400, 100, 90,
		50, 40, 10, 9, 5, 4, 1}
	letters = [...]string{"M", "CM", "D", "CD", "C", "XC",
		"L", "XL", "X", "IX", "V", "IV", "I"}
)

const (
	min = 1
	max = 3999
)

func NewRomanNumeralFromArabic(arabic int) (*RomanNumeral, error) {
	if arabic < min {
		return nil, ErrNumeralMustBePositive
	}
	if arabic > max {
		return nil, ErrNumeralMustBe3999OrLess
	}
	return &RomanNumeral{arabicVal: arabic}, nil
}

var validPattern = regexp.MustCompile(`[IVXLXCDM]+`)

func NewRomanNumeralFromRoman(roman string) (*RomanNumeral, error) {
	if len(roman) == 0 {
		return nil, ErrDoesNotDefineARomanNumeral
	}
	if roman[0] == '-' {
		return nil, ErrNoNegativeRomanNumerals
	}
	if !validPattern.MatchString(roman) {
		return nil, ErrInvalidRomanNumeral
	}

	return &RomanNumeral{romanVal: strings.ToUpper(roman)}, nil
}

func (r *RomanNumeral) toRoman() string {
	if r.romanVal != "" {
		return r.romanVal
	}

	remainingPartToConvert := r.arabicVal
	for i, n := range numbers {
		for remainingPartToConvert >= n {
			r.romanVal += letters[i]
			remainingPartToConvert -= n
		}
	}

	return r.romanVal
}

func (r *RomanNumeral) toArabic() (int, error) {
	if r.arabicVal > 0 {
		return r.arabicVal, nil
	}

	positionInString := 0

	roman := r.romanVal

	for positionInString < len(roman) {
		letter := roman[positionInString]
		number := letterToNumber(letter)
		if number < 0 {
			return 0, ErrNumberFormatException
		}
		positionInString++
		if positionInString == len(roman) {
			r.arabicVal += number
		} else {
			nextLetter := roman[positionInString]
			nextNumber := letterToNumber(nextLetter)
			if nextNumber > number {
				r.arabicVal += nextNumber - number
				positionInString++
			} else {
				r.arabicVal += number
			}
		}
	}

	if r.arabicVal > max {
		return 0, ErrNumeralMustBe3999OrLess
	}

	return r.arabicVal, nil
}

func letterToNumber(letter uint8) int {
	switch letter {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return -1
	}
}
