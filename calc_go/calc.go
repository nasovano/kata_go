package main

import (
		"os"
		"fmt"
		"bufio"
		"strings"
		)

type MathEx struct {
	A			int
	B			int
	Result		int
	RomaRes		string
	operator	byte	
}

type Flags struct {
	parts		int
	lenPart		[3]int
	trash		[10][10]bool
	format		[10][10]bool
	numeral		[10][10]string
	numeral_type	string
	first_operand	string
	second_operand	string
}

func ArabToInt(n string) int {
	out := 0
	ln := len(n)
	for i := 0; i < ln; i++ {
	out = out * 10 + int(byte(n[i]) - 48)
	}
	return out
}		

func main() {
	ME := &MathEx{}
	f := &Flags{}
	s :=  []string {} 
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please, input Math Expression")
		args, _ := reader.ReadString('\n')
		args = strings.TrimSpace(args)
		s = CheckParts(&args)
		CheckTrash(s, f)
		CheckFormat(s, f, ME)
		Math(f, s, ME)
	}
	return	   
}
