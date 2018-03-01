package main

import (
    "fmt"
    "strings"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    count := make(map[string]int)
    tokens := strings.Split(s, " ")
    for _, token := range tokens {
        count[token]++
    }
    return count
}

func main() {
    wc.Test(WordCount)
}
