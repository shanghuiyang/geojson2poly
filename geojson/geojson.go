package geojson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	ErrEmptyGeojson = errors.New("empty geojson")
	emptyPoly       = "NoName\nEND\n"
)

// Geojson ...
type Geojson struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

type Feature struct {
	Geometry   *Geometry   `json:"geometry"`
	Properties *Properties `json:"properties"`
	Type       string      `json:"type"`
}

// Geometry ...
type Geometry struct {
	Coordinates [][][]float64 `json:"coordinates"`
	Type        string        `json:"type"`
}

// Properties ...
type Properties struct {
	Name string `json:"name"`
}

// NewGeojson ...
func NewGeojson() *Geojson {
	return &Geojson{}
}

// Load ...
func (g *Geojson) Load(filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, g); err != nil {
		return err
	}

	return nil
}

// ToPoly ...
func (g *Geojson) ToPoly(filePath string) error {
	if g == nil {
		return ErrEmptyGeojson
	}
	if len(g.Features) == 0 {
		if err := ioutil.WriteFile(filePath, []byte(emptyPoly), 0644); err != nil {
			return err
		}
		return nil
	}

	s := "NoName\n"
	for _, fea := range g.Features {
		if fea.Properties != nil && fea.Properties.Name != "" {
			s = fea.Properties.Name
		}
		if !strings.Contains(fea.Geometry.Type, "Polygon") {
			continue
		}
		for i, linears := range fea.Geometry.Coordinates {
			s += fmt.Sprintf("Polygon%v\n", i+1)
			for _, points := range linears {
				s += fmt.Sprintf("\t%v %v\n", points[0], points[1])
			}

			s += "END\n"
		}
		s += "END\n"
	}

	if err := ioutil.WriteFile(filePath, []byte(s), 0644); err != nil {
		return err
	}

	return nil
}
