package main

import "fmt"
import "os"
import "log"
import "strconv"
import "strings"
import "image/png"

import "github.com/arbaregni/resistor-counts/rationals"

type ArgFlags = int
const (
    NoFlags ArgFlags = 1 << iota
    ShowLayers
    DoDerive
    DoVisualize
    Quit
)

type Args struct {
    flags ArgFlags
    n int
    p, q int
    filename string
}

func ParseArgs(cmdargs []string) (Args, bool) {
    args := Args{}

    if len(cmdargs) <= 1 {
        PrintHelp()
        return args, false
    }
    cmd, cmdargs := cmdargs[1], cmdargs[2:]
    switch cmd {
    case "layer":
        args.flags |= ShowLayers
        if len(cmdargs) < 1 {
            fmt.Println("Missing required parameter <n>. How many layers should I generate?")
            return args, false
        }
        args.n, _ = strconv.Atoi(cmdargs[0])
    case "derive":
        args.flags |= DoDerive
        if len(cmdargs) < 2 {
            fmt.Println("Missing required parameter(s) <p> and <q>. What rational number do you want to derive?")
            return args, false
        }
        args.p, _ = strconv.Atoi(cmdargs[0])
        args.q, _ = strconv.Atoi(cmdargs[1])
    case "visual":
        args.flags |= DoVisualize
        if len(cmdargs) < 1 {
            fmt.Println("Missing required parameter <n>. How many layers should I generate?")
            return args, false
        }
        args.n, _ = strconv.Atoi(cmdargs[0])
        cmdargs = cmdargs[1:]
        if len(cmdargs) == 0 {
            args.filename = "image.png"
        } else {
            args.filename = cmdargs[0]
        }
    default:
        fmt.Println("I don't understand this command:",cmd)
        return args, false
    }
    return args, true
}

func PrintHelp() {
    name := os.Args[0]
    {
        if len(name) == 0 {
            name = "This"
        }
        if strings.HasSuffix(name, "\\") || strings.HasSuffix(name, "/") {
            name = name[:len(name)-1]
        }
        i := strings.LastIndexAny(name, "\\/")
        if i != -1 {
            name = name[i+1:]
        }
    }
    fmt.Printf("%v is a tool for the RC problem\n", name)
    fmt.Println("    Usage:")
    fmt.Println("        <command> [arguments]")
    fmt.Println()
    fmt.Println("The commands are:")
    fmt.Println()
    fmt.Println("    layer <n>                              generate and display the n-resistor constructable numbers")
    fmt.Println("    derive <p> <q>                         derive the RC formula for the rational number p/q")
    fmt.Println("    visual <n> [filename] [options...]     generate the n-resistor constructable numbers and create the visualization")
    fmt.Println("                                               writing it to a file named filename (defaults to image.png)")
}

func main() {
	log.SetFlags(log.Lshortfile)

    args, ok := ParseArgs(os.Args)
    if !ok {
        return
    }
	dp := NewDP(args.n)
    layers := dp.Generate(args.n)
	fmt.Printf("generating R^%v\n", args.n)


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
