package main

// 0からnまでの整数で、素数かどうかををもつboolのスライスを返す
// O(N * loglogN)らしい
// https://algo-method.com/descriptions/64

// n should be more than 1
func NewEratosthenes(n int) []bool {
	isPrime := make([]bool, n+1)

	// set those nums true that are more than 1
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// set non-prime indexes false
	for i := 2; i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := 2 * i; j <= n; j += i {
			isPrime[j] = false
		}
	}

	return isPrime
}
