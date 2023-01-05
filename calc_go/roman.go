package main

import (
		"os"
		"fmt")

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numInv = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func RomanToInt(n string) int {
	if n == "I" { return 1
	} else if n == "II" { return 2
	} else if n == "III" { return 3
	} else if n == "IV" { return 4
	} else if n == "V" { return 5
	} else if n == "VI" { return 6 
	} else if n == "VII" { return 7 
	} else if n == "VIII" { return 8 
	} else if n == "IX" { return 9 
	} else if n == "X" { return 10 
	} else {
		fmt.Println("Error: incorrect Roman numeral\n",
				"Please, input Roman numeral in the range from I to X")
		os.Exit(1)	
	}
	return 0
}

// ToRoman is to convert int to roman numeral
func ToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}
