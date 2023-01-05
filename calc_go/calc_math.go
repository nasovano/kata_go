package main

import ("os"
		"fmt")

//to convert Numeral to Int
func NumeralToInt(f *Flags, slice []string, ME *MathEx) {
//Arabic numeral conversion to int		
	if f.numeral_type == "arab" {
		for i := 0; i < len(slice); i++ {
			if i == 0 {
				ME.A = ArabToInt(slice[i])
			}
			if i == 2 {
				ME.B = ArabToInt(slice[i])
			}
		}
	}
//Roman numeral conversion to int	
	if f.numeral_type == "roma" {
		for i := 0; i < len(slice); i++ {
			if i == 0 {
				ME.A = RomanToInt(slice[i])
			}
			if i == 2 {
				ME.B = RomanToInt(slice[i])
			}
		}
	}
	f.first_operand = ""
	f.second_operand = ""	
}

//doing Math: + or - or * or /
func MathOperation(ME *MathEx) int {
	if ME.A < 1 || ME.B < 1 ||
	ME.A > 10 || ME.B > 10 {
		fmt.Println("Error: incorrect range of numerals\n",
			"Please, input either Arabic numerals\n", 
			"in the range from 1 to 10\n",
			"or Roman numerals in the range from I to X")
			os.Exit(1)
	}
	if ME.operator == 43 {
		ME.Result = ME.A + ME.B
	} else if ME.operator == 45 {
		ME.Result = ME.A - ME.B
	} else if ME.operator == 42 {
		ME.Result = ME.A * ME.B
	} else if ME.operator == 47 {
		ME.Result = ME.A / ME.B
	}
	return ME.Result
}

//Math
func Math(f *Flags, slice []string, ME *MathEx) int {
	NumeralToInt(f, slice, ME)
	MathOperation(ME)
	if f.numeral_type == "roma" {
		if ME.Result < 0 {
			fmt.Println("Error: the result is empty because it's negative\n",
					"There is no negative Roman numeral\n")
			os.Exit(1)
		} else if ME.Result == 0 {
			fmt.Println(ME.Result)
			return 0	
		} else {	 
		ME.RomaRes = ToRoman(ME.Result)
		fmt.Println(ME.RomaRes)
		return 0	
	}}
	fmt.Println(ME.Result)
	return 0	
} 
