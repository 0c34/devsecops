package main

import (
    "fmt"
    "devsecops/module"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
    result = x + 2
    return result
}

func main() {
    fmt.Println("Hello World")
    module.Mo()
}