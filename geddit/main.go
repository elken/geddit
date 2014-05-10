package main

import (
    "fmt"
    "log"
    "geddit"
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
