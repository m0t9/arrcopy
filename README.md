# arrcopy

[![coverage](https://raw.githubusercontent.com/m0t9/arrcopy/badges/.badges/master/coverage.svg)](https://github.com/m0t9/arrcopy/actions/workflows/.testcoverage.yml)
[![build](https://github.com/m0t9/arrcopy/actions/workflows/go.yml/badge.svg)](https://github.com/m0t9/arrcopy/actions/workflows/go.yml)

Simple Go linter prohibiting `for-range` iteration over array copy.

## Details

The `range` operator creates a copy of the array before iterating over it.
In case of heavyweight arrays it may lead to significant performance degradation.
One can see explanation with clear benchmarks [here](https://medium.com/@haaawk/i-thought-i-understood-how-iteration-over-an-array-works-but-apparently-not-in-golang-441a7abd6540)

## Usage

`go install github.com/m0t9/arrcopy/cmd/...@latest`

Linter is compatible with `go vet`
