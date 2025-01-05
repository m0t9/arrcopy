# arrcopy

[![coverage](https://raw.githubusercontent.com/m0t9/arrcopy/badges/.badges/master/coverage.svg)](https://github.com/m0t9/arrcopy/actions/workflows/.testcoverage.yml)
[![build](https://github.com/m0t9/arrcopy/actions/workflows/go.yml/badge.svg)](https://github.com/m0t9/arrcopy/actions/workflows/go.yml)

Simple Go linter prohibiting `for-range` iteration over array copy.

## Details

The `range` operator creates a copy of the array before iterating over it.
In case of heavyweight arrays it may lead to significant performance degradation.
One can see explanation with clear benchmarks [here](https://medium.com/@haaawk/i-thought-i-understood-how-iteration-over-an-array-works-but-apparently-not-in-golang-441a7abd6540)

## Usage

### Installation

`go install github.com/m0t9/arrcopy/cmd/...@latest`

### Running

Linter is compatible with `go vet`

`go vet -vettool=$(which arrcopy) ./...` 

## Optimization example 

The inefficiency of array copying can be shown in the benchmark below

```go
func BenchmarkArrayIteration(b *testing.B) {
	const Size = 1000000
	type Item struct {
		s      string
		subarr [5]int64
		sl     []int
		p      *int
	}

	b.Run("over copy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := [Size]Item{}
			b.StartTimer()
			for idx, item := range arr {
				_, _ = idx, item
			}
		}
	})

	b.Run("over reference", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := [Size]Item{}
			b.StartTimer()
			for idx, item := range &arr { // added & before arr
				_, _ = idx, item
			}
		}
	})
}
```

```
goos: darwin
goarch: arm64
pkg: github.com/m0t9/arrcopy/cmd/arrcopy
cpu: Apple M1
BenchmarkArrayIteration/over_copy-8                   52          21382312 ns/op
BenchmarkArrayIteration/over_reference-8            3085            382052 ns/op
PASS
ok      github.com/m0t9/arrcopy/cmd/arrcopy     9.923s
```
