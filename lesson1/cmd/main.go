package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	fibo "pkg/fibonacci"
)

func main() {
	numb, err := askForNumb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numb, "fibonacci number is", fibo.Calculate(numb))
}

func askForNumb() (int, error) {
	var numb int
	fmt.Println("Enter the ordinal number of the fibonacci number (from 0 to 20):")

	fmt.Fscan(os.Stdin, &numb)
	if numb > 20 || numb < 0 {
		return 0, errors.New("Input mistake! Please enter a number between [0..20]")
	}

	return numb, nil
}
