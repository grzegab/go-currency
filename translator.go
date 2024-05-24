package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func Translate(a float64) string {
	var sb strings.Builder

	value, _ := math.Modf(a)
	friction := int(a*100) % 100

	mainText, _ := countMain(int(value))
	decimalText, _ := countDecimal(friction)

	sb.WriteString(mainText)

	sb.WriteString("i")
	sb.WriteString(decimalText)

	return sb.String()
}

func countDecimal(d int) (string, error) {
	var sb strings.Builder

	text, err := numberToText(d)
	if err != nil {
		return "", err
	}

	sb.WriteString(text)

	last := lastLetterFromText(text)

	sb.WriteString("grosz")

	if d == 1 {
		return sb.String(), nil
	}

	if last == 'a' || last == 'y' {
		sb.WriteString("e")
		return sb.String(), nil
	}

	sb.WriteString("y")

	return sb.String(), nil
}

func countMain(num int) (string, error) {
	var sb strings.Builder

	strNum := strconv.Itoa(num)
	revNum := reverseString(strNum)
	parts := []int{}

	for i := 0; i < len(revNum); i += 3 {
		end := i + 3
		if end > len(revNum) {
			end = len(revNum)
		}
		partStr := reverseString(revNum[i:end])
		partNum, _ := strconv.Atoi(partStr)
		parts = append(parts, partNum)
	}

	for j := len(parts) - 1; j >= 0; j-- {
		text, _ := numberToText(parts[j])
		sb.WriteString(text)
		sb.WriteString(" ")

		if j > 0 {
			sb.WriteString(multiThousand(j, text))
			sb.WriteString(" ")
		}

		if j == 0 {
			ll := lastLetterFromText(text)

			if parts[j] == 1 {
				sb.WriteString("złoty")
			} else if ll == 'a' {
				sb.WriteString("złote")
			} else {
				sb.WriteString("złotych")
			}

			sb.WriteString(" ")
		}
	}

	return sb.String(), nil
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
	if n > 100 && n < 500 {
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

		sb.WriteString(smallest(secondLastDigit))
		if secondLastDigit == 2 {
			sb.WriteString("dzieścia")
			sb.WriteString(" ")
		} else if secondLastDigit >= 3 && secondLastDigit <= 4 {
			sb.WriteString("dzieści")
			sb.WriteString(" ")
		} else if secondLastDigit > 5 {
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
