# goperflint

Linters with micro-optimizations making code faster without spoiling the quality of the code itself.

## arrcopy

The `range` operator creates a copy of the array before iterating over it.
In case of heavyweight arrays it may lead to significant performance degradation.
One can see explanation with clear benchmarks [here](https://medium.com/@haaawk/i-thought-i-understood-how-iteration-over-an-array-works-but-apparently-not-in-golang-441a7abd6540)

