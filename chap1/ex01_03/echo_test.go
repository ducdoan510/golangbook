package ex01_03

import (
	"golangbook/gentest"
	"testing"
)

var randomTests10 = gentest.RandomStrings(10, 25)
var randomTests100 = gentest.RandomStrings(100, 25)
var randomTests1000 = gentest.RandomStrings(1000, 25)

func BenchmarkEchoWithJoin10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithJoin(randomTests10)
	}
}

func BenchmarkEchoWithConcat10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithConcat(randomTests10)
	}
}

func BenchmarkEchoWithJoin100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithJoin(randomTests100)
	}
}

func BenchmarkEchoWithConcat100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithConcat(randomTests100)
	}
}

func BenchmarkEchoWithJoin1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithJoin(randomTests1000)
	}
}

func BenchmarkEchoWithConcat1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithConcat(randomTests1000)
	}
}