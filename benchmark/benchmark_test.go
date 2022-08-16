package main

import "testing"

//test or benchmark

func TestPrint01(t *testing.T) {
	if print01(100) != "100" {
		t.Fatal("error")
	}
}

func TestPrint02(t *testing.T) {
	if print02(int64(100)) != "100" {
		t.Fatal("error")
	}
}

func TestPrint03(t *testing.T) {
	if print03(100) != "100" {
		t.Fatal("error")
	}
}

func BenchmarkPrint01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print01(100)
	} //default 1 second, see how many times it can run in 1 sec.
}

func BenchmarkPrint02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print02(int64(100))
	} //default 1 second, see how many times it can run in 1 sec.
}

func BenchmarkPrint03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print03(100)
	} //default 1 second, see how many times it can run in 1 sec.
}
