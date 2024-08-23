package main

import (
	"errors"
	"log"
	"math"
	"strconv"
	"strings"
)

type lang int

const (
	PL lang = iota
	EN
)

type Translator struct {
	Input    float64
	Language lang
	output   strings.Builder
}

// Translate will turn float number into currency text value
func (t *Translator) Translate(a float64) string {
	i, _ := math.Modf(a)
	fraction := int(math.Round(a*100) - (i * 100))

	intToText, err := translateInt(int(i))
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}
	if len(intToText) > 0 {
		t.output.WriteString(intToText)
	}

	decimalText, err := translateFraction(fraction)
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}

	if len(decimalText) > 0 {
		t.output.WriteString("i")
		t.output.WriteString(decimalText)
	}

	return t.output.String()
}

func translateFraction(d int) (string, error) {
	var sb strings.Builder

	text, err := numberToText(d)
	if err != nil {
		return "", err
	}

	sb.WriteString(text)

	l := lastLetterFromText(text)

	sb.WriteString(" ")
	sb.WriteString("grosz")

	if d == 1 {
		return sb.String(), nil
	}

	if l == 'a' || l == 'y' {
		sb.WriteString("e")
	} else {
		sb.WriteString("y")
	}

	return sb.String(), nil
}

func translateInt(num int) (string, error) {
	var sb strings.Builder
	var p []int

	st := strconv.Itoa(num)
	rt := reverseString(st)

	// Divide number into group of 3 e.g. 123456: 123 & 456
	for i := 0; i < len(rt); i += 3 {
		end := i + 3
		if end > len(rt) {
			end = len(rt)
		}
		partStr := reverseString(rt[i:end])
		partNum, _ := strconv.Atoi(partStr)
		p = append(p, partNum)
	}

	for j := len(p) - 1; j >= 0; j-- {
		text, err := numberToText(p[j])
		if err != nil {
			return "", err
		}

		if !(len(sb.String()) > 1 && strings.Trim(text, " ") == "zero") {
			sb.WriteString(strings.Trim(text, " "))
			sb.WriteString(" ")
		}

		if j > 0 {
			sb.WriteString(multiThousand(j, text))
			sb.WriteString(" ")
		}

		if j == 0 {
			ll := lastLetterFromText(text)

			if p[j] == 1 {
				sb.WriteString("złoty")
			} else if ll == 'a' || ll == 'y' {
				sb.WriteString("złote")
			} else {
				sb.WriteString("złotych")
			}

			sb.WriteString(" ")
		}
	}

	t := sb.String()

	return t, nil
}

func lastLetterFromText(t string) rune {
	runesText := []rune(t)
	return runesText[len(runesText)-1]
}

func multiThousand(i int, t string) string {
	var sb strings.Builder

	ll := lastLetterFromText(t)

	switch i {
	case 1:
		if t == "jeden" {
			sb.WriteString("tysiąc")
		} else if ll == 'a' || ll == 'y' {
			sb.WriteString("tysiące")
		} else {
			sb.WriteString("tysięcy")
		}
	case 2:
		if t == "jeden" {
			sb.WriteString("milion")
		} else if ll == 'a' || ll == 'y' {
			sb.WriteString("miliony")
		} else {
			sb.WriteString("milionów")
		}
	case 3:
		if t == "jeden" {
			sb.WriteString("miliard")
		} else if ll == 'a' || ll == 'y' {
			sb.WriteString("miliardy")
		} else {
			sb.WriteString("miliardów")
		}
	}

	return sb.String()
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func numberToText(n int) (string, error) {
	var sb strings.Builder

	if n > 999 || n < 0 {
		return "", errors.New("wrong length of number")
	}

	firstDigit := n / 100
	if n >= 100 && n < 500 {
		switch firstDigit {
		case 1:
			sb.WriteString("sto")
		case 2:
			sb.WriteString("dwieście")
		case 3:
			sb.WriteString("trzysta")
		case 4:
			sb.WriteString("czterysta")
		}
	} else if n >= 500 {
		sb.WriteString(smallest(firstDigit))
		sb.WriteString("set")
	}

	sb.WriteString(" ")

	if n == 0 {
		sb.WriteString("zero")
	}

	lastTwoDigits := n % 100
	if lastTwoDigits > 0 && lastTwoDigits <= 10 {
		sb.WriteString(smallest(lastTwoDigits))
	} else if lastTwoDigits == 15 {
		sb.WriteString("piętnaście")
	} else if lastTwoDigits >= 12 && lastTwoDigits < 19 {
		sb.WriteString(smallest(lastTwoDigits % 10))
		sb.WriteString("naście")
	} else if lastTwoDigits == 11 {
		sb.WriteString("jedenaście")
	} else if lastTwoDigits == 19 {
		sb.WriteString("dziewiętnaście")
	} else {
		secondLastDigit := lastTwoDigits / 10 % 10

		sb.WriteString(firstForDouble(secondLastDigit))

		if secondLastDigit == 2 {
			sb.WriteString("dzieścia")
			sb.WriteString(" ")
		} else if secondLastDigit >= 3 && secondLastDigit <= 4 {
			sb.WriteString("dzieści")
			sb.WriteString(" ")
		} else if secondLastDigit >= 5 {
			sb.WriteString("dziesiąt")
			sb.WriteString(" ")
		}

		lastDigit := n % 10
		if lastDigit > 0 {
			sb.WriteString(smallest(lastDigit))
		}
	}

	return sb.String(), nil
}

func firstForDouble(num int) string {
	var sb strings.Builder

	switch num {
	case 2:
		sb.WriteString("dwa")
	case 3:
		sb.WriteString("trzy")
	case 4:
		sb.WriteString("czter")
	case 5:
		sb.WriteString("pięć")
	case 6:
		sb.WriteString("sześć")
	case 7:
		sb.WriteString("siedem")
	case 8:
		sb.WriteString("osiem")
	case 9:
		sb.WriteString("dziewięć")
	}

	return sb.String()
}

func smallest(num int) string {
	var sb strings.Builder

	switch num {
	case 1:
		sb.WriteString("jeden")
	case 2:
		sb.WriteString("dwa")
	case 3:
		sb.WriteString("trzy")
	case 4:
		sb.WriteString("cztery")
	case 5:
		sb.WriteString("pięć")
	case 6:
		sb.WriteString("sześć")
	case 7:
		sb.WriteString("siedem")
	case 8:
		sb.WriteString("osiem")
	case 9:
		sb.WriteString("dziewięć")
	case 10:
		sb.WriteString("dziesięć")
	}

	return sb.String()
}
