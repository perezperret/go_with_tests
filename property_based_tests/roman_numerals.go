package main

import (
	"strings"
)

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = RomanNumerals{
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

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	i := 0

	for i < len(roman)-1 {
		symbol := roman[i]
		nextSymbol := roman[i+1]

		value := allRomanNumerals.ValueOf(symbol)
		nextValue := allRomanNumerals.ValueOf(nextSymbol)

		if value < nextValue {
			total += (nextValue - value)
			i += 2
		} else {
			total += value
			i++
		}
	}

	if i == len(roman)-1 {
		total += allRomanNumerals.ValueOf(roman[len(roman)-1])
	}

	return total
}
