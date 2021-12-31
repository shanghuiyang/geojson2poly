# geojson2poly
[![CI](https://github.com/shanghuiyang/geojson2poly/actions/workflows/ci.yml/badge.svg)](https://github.com/shanghuiyang/geojson2poly/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/shanghuiyang/geojson2poly/blob/master/LICENSE)

Converts [geojson](http://geojson.org/) Polygon/MultiPolygon to openstreetmap(osm) [poly](https://wiki.openstreetmap.org/wiki/Osmosis/Polygon_Filter_File_Format) format file.

## Usage
```
usage: geojson2poly --input=GEOJSON --output=POLY [<flags>]

Flags:
  -h, --help           Show context-sensitive help (also try --help-long and --help-man).
  -i, --input=GEOJSON  a .geojson file as input
  -o, --output=POLY    a .poly file as output
```

example,
```shell
$ geojson2poly -i test.geojson -o test.poly
```

## Install
**Build from sourcecodes**
```shell
$ go get -u github.com/shanghuiyang/geojson2poly
# from project directory
$ go build .
$ cp cp geojson2poly /usr/local/bin
# test
$ geojson2poly -h
```
**Download from the per-compiled binary**
Download from [Release](https://github.com/shanghuiyang/geojson2poly/releases)
