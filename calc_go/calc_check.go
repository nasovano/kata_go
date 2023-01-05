package main

import ("os"
		"fmt"
		"strings")

//check how many parts are in Math Expression
func CheckParts(text *string) ([]string) {
	slice := strings.Fields(*text)
	if len(slice) != 3 {
		fmt.Println("Error: incorrect number of components\n",
				"of the Mathematical Expression.\n",
				"Please, input Math Expression in three parts\n",
				"each separated by whitespace")
		os.Exit(1)}
	return slice	
}	   

//chesk every byte of the text for possible trash
func check_trash(some byte, partNumber int) bool {
		not_arab := false 
		not_roma := false 
		not_oper := false
//check if not Arabic numeral		
	if (partNumber == 0 || partNumber == 2) && 
		(some < 48 || some > 57) {
		if !not_arab {
		not_arab = true }
	}
//check if not Roman numeral		
	if (partNumber == 0 || partNumber == 2) && 
		(some != 73 && some != 86 && some != 88) {
		if !not_roma {
		not_roma = true } 
	}
//check if not opeartor		
	if partNumber == 1 && some != 42 && 
		some != 43 && some != 45 && some != 47 {
		if !not_oper {
		not_oper = true }
	}
//check if not opeartor or if not any approved numeral		
	if not_oper || (not_arab && not_roma) {
		return true }
	return false
}

//check trash in the text
func CheckTrash(slice []string, f *Flags) {
	for i := 0; i < len(slice); i++ {
			f.lenPart[i] = len(slice[i])
		for j := 0; j < len(slice[i]); j++ {
			f.trash[i][j] = check_trash(slice[i][j], i)
			if f.trash[i][j] { 
				fmt.Println("Error: some symbols are out of range\n",
						"Please, input Math Expression with:\n",
						"either Arabic Numearls from 1 to 10\n",
						"or Roman Numerals from I to X\n",
						"and math operators such as: + - * /")
				os.Exit(1) 
			}
		}
	}
}


	
