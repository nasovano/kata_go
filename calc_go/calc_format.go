package main

import ("os"
		"fmt")

// get numeral type of the operand
func GetOperandType(some byte) string {
	arab := false
	roma := false		  
	if (some >= 48 && some <= 57) {
		if !arab { arab = true 
		return "arab" 
		}
	}
	if (some == 73 || some == 86 || some == 88) {
		if !roma { roma = true 
		return "roma" 
		}
	}
	return "false"	
}

//get numeral type of both operands and check if they are equal
//check correctness of the operator and if yes, then get it
func GetType(some byte, f *Flags, partNumber int, byteNumber int, ME *MathEx) bool {
	oper := false
//check firs operand type
	if partNumber == 0  {
	f.numeral[partNumber][byteNumber] = GetOperandType(some)
	f.first_operand = f.numeral[partNumber][byteNumber]
		if byteNumber > 0 && f.numeral[partNumber][byteNumber] != "" {
			if f.numeral[partNumber][byteNumber] != f.numeral[partNumber][byteNumber - 1] {
				fmt.Println("Error: Different numerals in the same operand\n")
				os.Exit(1)	
			}
		}
	}
//check second operand type
	if partNumber == 2  {
	f.numeral[partNumber][byteNumber] = GetOperandType(some)
	f.second_operand = f.numeral[partNumber][byteNumber]
		if byteNumber > 0 && f.numeral[partNumber][byteNumber] != "" {
			if f.numeral[partNumber][byteNumber] != f.numeral[partNumber][byteNumber - 1] {
				fmt.Println("Error: Different numerals in the same operand\n")
				os.Exit(1)	
			}
		}
	}
// check and get operator			  
	if partNumber == 1 && byteNumber > 0 {
		fmt.Println("Error: incorrect operators' format\n",
				"Please, input math operator with one symbol only")
		os.Exit(1)	
	} else if partNumber == 1 && byteNumber == 0 {
		ME.operator = some
		if !oper { oper = true }
	}
//check different operands for the same numeral type
	if f.first_operand != "" && f.second_operand != "" {
		if f.first_operand == f.second_operand {
		f.numeral_type = f.first_operand
	} else if f.first_operand != f.second_operand {
		fmt.Println("Error: Numerals of different types in the same Math Expression\n")
		os.Exit(1)
	}}
	return true 
}

//check both operands for the same numeral type, check and get operator
func CheckFormat(slice []string, f *Flags, ME *MathEx) {
	for i := 0; i < len(slice); i++ {
			f.lenPart[i] = len(slice[i])
		for j := 0; j < len(slice[i]); j++ {
			f.format[i][j] = GetType(slice[i][j], f, i, j, ME)
			if !f.format[i][j] { 
				fmt.Println("Error: f.format is false\n")
				os.Exit(1) 
			}
		}
	}
}


	
