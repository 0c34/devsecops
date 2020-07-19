package main

import (
    "fmt"
    "devsecops/module"
    "devsecops/module/submodule"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
    result = x + 2
    return result
}

func main() {
    fmt.Println("Hello World")
    module.Mo()
    submodule.SubMo()
}