package main

import "fmt"
import "os"
import "strconv"

func main() {
    n := func() int {
        if len(os.Args) < 2 {
            fmt.Println("Missing required argument <n>")
            return 10
        }
        arg := os.Args[1]
        n, _ := strconv.ParseInt(arg,0,64)
        return int(n)
    }()

    fmt.Printf("generating R^%v\n", n)

    Generate(n)

}
