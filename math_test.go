package main

import "testing"

func TestGcd(t *testing.T) {
	result := gcd(123456789, 987654321)
	if result != 9 {
		t.Fatal("failed")
	}
}

func TestGcdCoprime(t *testing.T) {
	result := gcd(123456789, 98765432)
	if result != 1 {
		t.Fatal("failed")
	}
}

func TestGcdNegative(t *testing.T) {
	result := gcd(-123456789, -98765432)
	if result != 1 {
		t.Fatal("failed")
	}
}
