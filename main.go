package main

import (
	"fmt"
	"os"

	"github.com/shanghuiyang/geojson2poly/geojson"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("error: invalid args")
		fmt.Println("usage: geojson2poly -i infile.json -o outfile.poly")
		os.Exit(1)
	}
	var input string
	var output string
	for i, arg := range os.Args {
		if arg == "-i" || arg == "-I" {
			input = os.Args[i+1]
		}
		if arg == "-o" || arg == "-O" {
			output = os.Args[i+1]
		}
	}

	fmt.Printf("input: %v\n", input)
	fmt.Printf("output: %v\n", output)

	g := geojson.NewGeojson()
	err := g.Load(input)
	if err != nil {
		fmt.Printf("error: fail to load geojson from: %v\n", input)
		fmt.Printf("err msg: %v\n", err)
		os.Exit(1)
	}

	err = g.ToPoly(output)
	if err != nil {
		fmt.Printf("error: fail to save poly file: %v\n", output)
		fmt.Printf("err msg: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("done in success")
	os.Exit(0)
}
