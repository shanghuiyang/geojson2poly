[![Build Status](https://travis-ci.org/shanghuiyang/geojson2poly.svg?branch=master)](https://travis-ci.org/shanghuiyang/geojson2poly)


# geojson2poly
Converts geojson polygons to openstreetmap(osm) poly format file.

## What is geojson?
http://geojson.org/


## What is osm poly format file
https://wiki.openstreetmap.org/wiki/Osmosis/Polygon_Filter_File_Format

## Install
build and install from source codes,
```shell
$ go get -u github.com/shanghuiyang/geojson2poly
$ cd $GOPATH/src/github.com/shanghuiyang/geojson2poly
$ go install
```

or download the binary from [here](https://github.com/shanghuiyang/geojson2poly/releases)

## Usage
```shell
$ geojson2poly -i input.geojson -o output.poly
```
