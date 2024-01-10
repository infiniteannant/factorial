package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter a non negative to find its factorial -> ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input := scan.Text()
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result, err := fact(num)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Factorial of %d is:\n%s\n", num, result)
}

func fact(n int) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("input must be a non -ve no.")
	}

	res := new(big.Int).SetInt64(1)

	for i := 2; i <= n; i++ {
		res.Mul(res, new(big.Int).SetInt64(int64(i)))
	}

	return res, nil
}
