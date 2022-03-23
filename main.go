package main

import "fmt"
import "os"
import "strconv"
import "image/png"

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
    layers := Generate(n)

    fmt.Printf("creating image....\n")
    img := Visualize(layers, 256, 256)

    file, err := os.Create("image.png")
    if err != nil {
        fmt.Println("Problems creating image file:", err)
    }
    defer file.Close()

    png.Encode(file, img)
}
