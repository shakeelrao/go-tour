package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

func walk(t *tree.Tree, ch chan int) {
    if t != nil {
        /* in-order traversal */
        walk(t.Left, ch)
        ch <- t.Value
        walk(t.Right, ch)
    }
}

func Walk(t *tree.Tree, ch chan int) {
    walk(t, ch)
    close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
    ch1, ch2 := make(chan int), make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for {
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2
        if (ok1 != ok2) || (v1 != v2) {
            return false
        } else if !ok1 {
            break
        }
    }
    return true
}

func main() {
    /* test Walk */
    ch := make(chan int)
    go Walk(tree.New(1), ch)
    for v := range ch {
        fmt.Println(v)
    }
    /* test Same */
    fmt.Println("test 1:", Same(tree.New(1), tree.New(1)))
    fmt.Println("test 2:", Same(tree.New(1), tree.New(2)))
}
