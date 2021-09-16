package pbt_test

import (
	"fmt"
	"learngowithtests/pbt"
	"log"
	"testing"
	"testing/quick"
)

var allRomanSymbols = []byte{
	'M',
	'D',
	'C',
	'L',
	'X',
	'V',
	'I',
}

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1998, Roman: "MCMXCVIII"},
	{Arabic: 1999, Roman: "MCMXCIX"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, err := pbt.ConvertToRoman(test.Arabic)

			if err != nil {
				t.Error("didn't expect an error but got one", err)
			}

			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := pbt.ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	t.Run("assert both functions returns same n", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}
			roman, _ := pbt.ConvertToRoman(arabic)
			fromRoman := pbt.ConvertToArabic(roman)
			return fromRoman == arabic
		}

		if err := quick.Check(assertion, &quick.Config{
			MaxCount: 1000,
		}); err != nil {
			t.Error("failed checks", err)
		}
	})

	t.Run("assert values greater than 3999 returns err", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic < 3999 {
				return true
			}
			_, err := pbt.ConvertToRoman(arabic)
			return err != nil
		}

		if err := quick.Check(assertion, &quick.Config{
			MaxCount: 1000,
		}); err != nil {
			t.Error("failed checks", err)
		}
	})

	t.Run("assert no more than 3 consecutive symbols", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}

			roman, _ := pbt.ConvertToRoman(arabic)
			log.Println(roman)
			log.Println(assertNoMoreThan3Consecutive(roman))
			return assertNoMoreThan3Consecutive(roman)
		}

		if err := quick.Check(assertion, &quick.Config{
			MaxCount: 1000,
		}); err != nil {
			t.Error("failed checks", err)
		}
	})
}

func assertNoMoreThan3Consecutive(roman string) bool {
	// MMMCDXXXIX
	var lastRune rune
	var lastRuneCount = 0
	for _, c := range roman {
		if c == lastRune {
			lastRuneCount++
			if lastRuneCount > 3 {
				return false
			}
		} else {
			lastRune = c
			lastRuneCount = 1
		}
	}

	return true
}
