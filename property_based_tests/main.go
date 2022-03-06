package main

import "strings"

type romanNumber struct {
	Value  uint16
	Symbol string
}

type romanNumbers []romanNumber

var AllRomanNumbers = romanNumbers{
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

func ConvertToRoman(number uint16) string {
	var result strings.Builder

	for _, numeral := range AllRomanNumbers {
		for number >= numeral.Value {
			result.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var total uint16 = 0
	length := uint16(len(roman))

	for i := uint16(0); i < length; i++ {
		symbol := roman[i]

		if i+1 < length && isSubtractive(symbol) {
			nextSymbol := roman[i+1]

			if value := AllRomanNumbers.ValueOf(symbol, nextSymbol); value != 0 {
				total += value
				i++
			} else {
				total += AllRomanNumbers.ValueOf(symbol)
			}
		} else {
			total += AllRomanNumbers.ValueOf(symbol)
		}
	}

	return total
}

func (r romanNumbers) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if symbol == s.Symbol {
			return s.Value
		}
	}
	return 0
}

func isSubtractive(symbol byte) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
