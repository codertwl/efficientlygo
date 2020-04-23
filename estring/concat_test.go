package estring

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func Benchmark_Native(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = s + strconv.Itoa(i)
	}
	_ = s
}

func Benchmark_Native2(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += strconv.Itoa(i)
	}
	_ = s
}

func Benchmark_Native3(b *testing.B) {
	var a []byte
	for i := 0; i < b.N; i++ {
		a = append(a,[]byte(strconv.Itoa(i))...)
	}
	_ = string(a)
}

func Benchmark_Join(b *testing.B) {
	var a []string
	for i := 0; i < b.N; i++ {
		a = append(a, strconv.Itoa(i))
	}
	_ = strings.Join(a,"")
}

func Benchmark_BytesBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}
	_ = buffer.String()
}

func Benchmark_StringsBuilder(b *testing.B) {
	var s strings.Builder
	for i := 0; i < b.N; i++ {
		s.WriteString(strconv.Itoa(i))
	}
	_ = s.String()
}