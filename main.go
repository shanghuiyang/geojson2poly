package main

import (
	"fmt"
	"io/ioutil"
	"os"

	geojson "github.com/paulmach/go.geojson"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	input := kingpin.Flag("input", "a .geojson file as input").Short('i').PlaceHolder("GEOJSON").Required().String()
	output := kingpin.Flag("output", "a .poly file as output").Short('o').PlaceHolder("POLY").Required().String()
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	rawJSON, err := ioutil.ReadFile(*input)
	if err != nil {
		fmt.Printf("ERROR! %v\n", err)
		os.Exit(1)

	}

	// assume the rawJSON is a geojson Feature
	f, err := geojson.UnmarshalFeature(rawJSON)
	if err != nil {
		fmt.Printf("ERROR! %v\n", err)
		os.Exit(1)
	}

	result := "no_name"
	if name, err := f.PropertyString("name"); err == nil {
		result = name
	}
	result += "\n"

	switch f.Type {
	case "Feature":
		result += fromFeature(f)
	case "FeatureCollection":
		// try to treat rawJSON as FeatureCollection since it isn't a Feature.
		fc, err := geojson.UnmarshalFeatureCollection(rawJSON)
		if err != nil {
			fmt.Printf("ERROR! %v\n", err)
			os.Exit(1)
		}
		result += fromFeatureCollection(fc)
	default:
		fmt.Println("ERROR! unsupported feature type")
		os.Exit(1)
	}
	result += "END\n"

	if err := ioutil.WriteFile(*output, []byte(result), 0644); err != nil {
		fmt.Printf("ERROR! %v\n", err)
		os.Exit(1)
	}
	fmt.Println("success")
	os.Exit(0)
}

func fromFeature(f *geojson.Feature) string {
	result := ""
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
	default:
		return ""
	}

	return result
}

func fromFeatureCollection(fc *geojson.FeatureCollection) string {
	result := ""
	for _, f := range fc.Features {
		result += fromFeature(f)
	}
	return result
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
