package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimesInRange(start, end int, primeChan chan int, doneChan chan bool) {
	for i := start; i*i<= end; i++ {
		if isPrime(i) && end % i == 0{
			primeChan <- i
		}
	}
	doneChan <- true
}

func fPrime(n int, allPrime []int) []int {
	result := []int{}
	i:=0
	for {
		if n % allPrime[i] == 0 {
			result = append(result, allPrime[i])
			n /= allPrime[i]
		} else if n == 1 {
			break
		} else {
			i++
		}
	}
	return result
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 1 {
		return
	}

	n, err := strconv.Atoi(arguments[0])
	if err != nil {
		return
	}

	primeChan := make(chan int)
	doneChan := make(chan bool)


	go findPrimesInRange(2, n, primeChan, doneChan)

	go func() {
		<-doneChan
		close(primeChan)
		close(doneChan)
	}()

	_result := []int{}

	for prime := range primeChan {
		_result = append(_result, prime)
	}

	result := fPrime(n, _result)

	fmt.Println(result)
}
