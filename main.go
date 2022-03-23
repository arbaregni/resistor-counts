package main

import "fmt"
import "os"
import "log"
import "strconv"
import "image/png"

import "github.com/arbaregni/resistor-counts/rationals"

func main() {
	log.SetFlags(log.Lshortfile)

	n := func() int {
		if len(os.Args) < 2 {
			fmt.Println("Missing required argument <n>")
			return 10
		}
		arg := os.Args[1]
		n, _ := strconv.ParseInt(arg, 0, 64)
		return int(n)
	}()

	fmt.Printf("generating R^%v\n", n)

	dp := NewDP(n)
	layers := dp.Generate(n)

	r := rationals.MakeRational(3, 4)
	formula, _ := dp.Construct(r)
	fmt.Printf("%v = %v\n", r, formula)

	fmt.Printf("creating image....\n")
	img := Visualize(layers, 256, 256)

	file, err := os.Create("image.png")
	if err != nil {
		fmt.Println("Problems creating image file:", err)
	}
	defer file.Close()

	png.Encode(file, img)
}
