[![Build Status](https://travis-ci.org/shanghuiyang/geojson2poly.svg?branch=master)](https://travis-ci.org/shanghuiyang/geojson2poly)


# geojson2poly
Converts geojson polygons to openstreetmap(osm) poly format file.

## What is geojson?
http://geojson.org/


## What is osm poly format file
https://wiki.openstreetmap.org/wiki/Osmosis/Polygon_Filter_File_Format

## Install
```shell
$ go get github.com/shanghuiyang/geojson2poly
```

## Build
```shell
$ go build -o geojson2poly main.go
```

## Usage
```shell
$ geojson2poly -i infile.json -o outfile.poly
```
