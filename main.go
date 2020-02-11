package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("Please set vaild flag")
	}

	switch os.Args[1] {
	case "-n":
		normal()
		return nil
	case "-c":
		concurrency()
		return nil
	}

	return fmt.Errorf("Undexpected flag")
}

func normal() {
	for i := 0; i < 100; i++ {
		fizzbuzz(i + 1)
		fmt.Print(" ")
	}
	fmt.Println()
}

func concurrency() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fizzbuzz(i + 1)
			fmt.Print(" ")
		}
		fmt.Println()
	}()
	wg.Wait()
}

func fizzbuzz(n int) {
	switch {
	case n%15 == 0:
		fmt.Print("FizzBuzz")
	case n%3 == 0:
		fmt.Print("Fizz")
	case n%5 == 0:
		fmt.Print("Buzz")
	default:
		fmt.Print(n)
	}
}
