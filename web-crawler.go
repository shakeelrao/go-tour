package main

import (
    "fmt"
    "sync"
)

/* create a concurrent map */
type SafeMap struct {
    m  map[string]bool
    mux sync.Mutex
}

func (s *SafeMap) Add(k string, v bool) {
    s.mux.Lock()
    defer s.mux.Unlock()
    s.m[k] = v
}

func (s *SafeMap) Exists(k string) bool {
    s.mux.Lock()
    defer s.mux.Unlock()
    _, ok := s.m[k]
    return ok
}

var visited = SafeMap{m: make(map[string]bool)}

type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
    defer wg.Done()
    if depth <= 0 {
        return
    }
    /* skip URL */
    if visited.Exists(url) {
        return
    }
    /* add URL to visited */
    visited.Add(url, true)
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        wg.Add(1)
        go Crawl(u, depth-1, fetcher, wg)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    Crawl("https://golang.org/", 4, fetcher, &wg)
    wg.Wait()
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
    "https://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "https://golang.org/pkg/",
            "https://golang.org/cmd/",
        },
    },
    "https://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "https://golang.org/",
            "https://golang.org/cmd/",
            "https://golang.org/pkg/fmt/",
            "https://golang.org/pkg/os/",
        },
    },
    "https://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
    "https://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
}
