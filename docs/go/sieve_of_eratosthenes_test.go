package main

import (
	"testing"
)

func TestNewEratosthenes(t *testing.T) {
	isPrime := NewEratosthenes(100)
	for i := 0; i < len(isPrime); i++ {
		t.Error(i, isPrime[i])
	}
}
