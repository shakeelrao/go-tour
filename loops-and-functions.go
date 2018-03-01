package main

import (
    "fmt"
    "math"
)

const eps float64 = 1e-10

func Sqrt(x float64) float64 {
    z := 1.0
    for math.Abs(z - math.Sqrt(x)) > eps {
        z -= (z*z - x) / (2*z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
