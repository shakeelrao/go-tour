package main

import (
    "io"
    "os"
    "strings"
    "unicode"
)

func rot13(char *byte) {
    // convert to lower-case
    c := unicode.ToLower(rune(*char))
    // rotate by 13
    if c >= 'a' && c < 'n' {
        *char += 13
    } else if c >= 'n' && c <= 'z' {
        *char -= 13
    }
}

type rot13Reader struct {
    r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
    n, err := rot.r.Read(b)
    for i := range b {
        rot13(&b[i])
    }
    return n, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)

