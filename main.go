package main

import (
    "fmt"
    "log"
    "github.com/elken/geddit/geddit"
)

func main() {
    items, err := geddit.Get("golang")
    if err != nil {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Println(item)
    }
}
