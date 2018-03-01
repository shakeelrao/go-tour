package main

import "golang.org/x/tour/reader"
import "errors"
import "fmt"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
    for i := range b {
        b[i] = 'A'
    }
    // attempt to replicate the ternary operator xD
    err := errors.New("error: buffer empty")
    if len(b) != 0 {
        err = nil
    }
    return len(b), err
}

func main() {
    reader.Validate(MyReader{})
    // test the error case
    n, err := MyReader{}.Read(make([]byte, 0))
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("read %d bytes", n)
    }
}
