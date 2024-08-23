package main

import (
	"testing"
)

func Test_unit_decimal(t *testing.T) {
	c := []struct {
		name     string
		number   float64
		expected string
	}{
		{"fractional zero", 0.99, "zero złotych i dziewięćdziesiąt dziewięć groszy"},
		{"fractional one", 1.11, "jeden złoty i jedenaście groszy"},
		{"fractional two", 1.23, "jeden złoty i dwadzieścia trzy grosze"},
		{"fractional tree", 1.35, "jeden złoty i trzydzieści pięć groszy"},
		{"fractional four", 1.49, "jeden złoty i czterdzieści dziewięć groszy"},
		{"fractional five", 1.52, "jeden złoty i pięćdziesiąt dwa grosze"},
		{"fractional six", 1.64, "jeden złoty i sześćdziesiąt cztery grosze"},
		{"fractional seven", 1.76, "jeden złoty i siedemdziesiąt sześć groszy"},
		{"fractional eight", 1.88, "jeden złoty i osiemdziesiąt osiem groszy"},
		{"fractional nine", 1.97, "jeden złoty i dziewięćdziesiąt siedem groszy"},
		{"fractional zero", 1.01, "jeden złoty i jeden grosz"},
		{"main simple", 3, "trzy złote i zero groszy"},
		{"main one", 10, "dziesięć złotych i zero groszy"},
		{"main two", 45, "czterdzieści pięć złotych i zero groszy"},
		{"main tree", 121, "sto dwadzieścia jeden złotych i zero groszy"},
		{"main four", 234, "dwieście trzydzieści cztery złote i zero groszy"},
		{"main five", 332, "trzysta trzydzieści dwa złote i zero groszy"},
		{"main six", 496, "czterysta dziewięćdziesiąt sześć złotych i zero groszy"},
		{"main seven", 507, "pięćset siedem złotych i zero groszy"},
		{"main eight", 618, "sześćset osiemnaście złotych i zero groszy"},
		{"main nine", 799, "siedemset dziewięćdziesiąt dziewięć złotych i zero groszy"},
		{"main zero", 800, "osiemset złotych i zero groszy"},
		{"main nine", 973, "dziewięćset siedemdziesiąt trzy złote i zero groszy"},
		{"complex big", 2_317_468_234.32987, "dwa miliardy trzysta siedemnaście milionów czterysta sześćdziesiąt osiem tysięcy dwieście trzydzieści cztery złote i trzydzieści trzy grosze"},
		{"complex hundred", 100.99, "sto złotych i dziewięćdziesiąt dziewięć groszy"},
		{"complex thousand", 123_456.11, "sto dwadzieścia trzy tysiące czterysta pięćdziesiąt sześć złotych i jedenaście groszy"},
		{"complex thousand simple", 19_000, "dziewiętnaście tysięcy złotych i zero groszy"},
		{"complex random 1", 23_568.37, "dwadzieścia trzy tysiące pięćset sześćdziesiąt osiem złotych i trzydzieści siedem groszy"},
		{"complex random 2", 35_677.21, "trzydzieści pięć tysięcy sześćset siedemdziesiąt siedem złotych i dwadzieścia jeden groszy"},
		{"complex random 3", 9_844.09, "dziewięć tysięcy osiemset czterdzieści cztery złote i dziewięć groszy"},
		{"complex random 4", 2_000, "dwa tysiące złotych i zero groszy"},
	}

	for _, ov := range c {
		t.Run(ov.name, func(t *testing.T) {
			tr := Translator{}
			output := tr.Translate(ov.number)
			if output != ov.expected {
				t.Errorf("%s: expected %s, got %s", ov.name, ov.expected, output)
			}
		})
	}
}
