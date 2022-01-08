package utils

import "fmt"

// var lowNames = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
// var tensNames = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
// var bigNames = []string{"thousand", "million", "billion"}

var lowNames = []string{"zero", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove", "dez", "onze", "doze", "treze", "quartoze", "quinze", "dezesseis", "dezesete", "dezoito", "dezenove"}
var tensNames = []string{"vinte", "trinta", "quarenta", "cinquenta", "secenta", "setenta", "oitenta", "noventa"}
var bigNames = []string{"mil", "milhão", "bilhão"}

func convert999(num int) string {
	s1 := lowNames[num/100] + " hundred"
	s2 := convert99(num % 100)
	if num <= 99 {
		return s2
	}
	if num == 0 {
		return s1
	} else {
		return s1 + " " + s2
	}
}

func convert99(num int) string {
	if num < 20 {
		return lowNames[num]
	}
	s := tensNames[num/10-2]
	if num == 0 {
		return s
	}
	return s + "-" + lowNames[num]
}

func ConvertNum2Words(num int) string {
	if num < 0 {
		return "negative " + ConvertNum2Words(-num)
	}

	if num <= 999 {
		return convert999(num)
	}

	s := ""
	t := 0
	for num > 0 {
		if num != 0 {
			s2 := convert999(num % 1000)
			if t > 0 {
				s2 = s2 + " " + bigNames[t-1]
			}
			if s == "" {
				s = s2
			} else {
				s = s2 + ", " + s
			}
		}
		num /= 1000
		t++
	}
	return s
}

func SecondsToMinutes(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("d:d", minutes, seconds)
	return str
}
