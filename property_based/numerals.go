package main

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var ArabicToRoman = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, num := range ArabicToRoman {
		for arabic >= num.Value {
			result.WriteString(num.Symbol)
			arabic -= num.Value
		}
	}

	return result.String()
}

var RomanToArabic = map[RomanNum]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func ConvertToArabic(roman string) (arabic int) {
	romanNums := GetRomanNumeralsReversed(roman)

	var prevNum RomanNum

	for _, r := range romanNums {
		if prevNum == "" {
			arabic += r.ToArabic()
			prevNum = r
			continue
		}

		if prevNum.ToArabic() > r.ToArabic() {
			arabic -= r.ToArabic()
			prevNum = r
			continue
		}

		arabic += r.ToArabic()
		prevNum = r
	}

	return arabic
}

type RomanNumerals []RomanNum
type RomanNum string

func (r RomanNum) ToArabic() int {
	return RomanToArabic[r]
}

func GetRomanNumeralsReversed(roman string) []RomanNum {
	reversed := Reverse(roman)
	numerals := strings.Split(reversed, "")

	result := make([]RomanNum, len(reversed))
	for i, r := range numerals {
		result[i] = RomanNum(r)
	}

	return result
}

func Reverse(s string) (reversed string) {
	for _, r := range s {
		reversed += string(r)
	}
	return reversed
}
