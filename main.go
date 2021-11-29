package main

import (
	"fmt"
	"io/ioutil"
	"os"

	geojson "github.com/paulmach/go.geojson"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("ERROR! invalid args")
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

	rawJSON, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Printf("ERROR! %v\n", err)
		os.Exit(1)

	}

	f, err := geojson.UnmarshalFeature(rawJSON)
	if err != nil {
		fmt.Printf("ERROR! %v\n", err)
		os.Exit(1)
	}

	if f.Type != "Feature" {
		fmt.Println("ERROR! unsupported geojson")
		os.Exit(1)
	}

	result := "no_name"
	if name, err := f.PropertyString("name"); err == nil {
		result = name
	}
	result += "\n"

	switch f.Geometry.Type {
	case geojson.GeometryPolygon:
		result += "1\n"
		result += toPoly(f.Geometry.Polygon)
		result += "END\n"
	case geojson.GeometryMultiPolygon:
		for i, p := range f.Geometry.MultiPolygon {
			result += fmt.Sprintf("%v\n", i+1)
			result += toPoly(p)
			result += "END\n"
		}
	}
	result += "END\n"
	if err := ioutil.WriteFile(output, []byte(result), 0644); err != nil {
		fmt.Printf("ERROR! %v\n", err)
		os.Exit(1)
	}
	fmt.Println("success")
	os.Exit(0)
}

func toPoly(polygon [][][]float64) string {
	if len(polygon) == 0 {
		return ""
	}

	res := ""
	for _, line := range polygon {
		for _, pt := range line {
			res += fmt.Sprintf("\t%v %v\n", pt[0], pt[1])
		}
	}
	return res
}
