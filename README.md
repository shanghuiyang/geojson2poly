# geojson2poly
Converts geojson polygons to openstreetmap(osm) poly format file.

## What is geojson?
http://geojson.org/


## What is osm poly format file
https://wiki.openstreetmap.org/wiki/Osmosis/Polygon_Filter_File_Format


## Build
```
$ go build -o geojson2poly main.go
```


## Usage
```
$ geojson2poly -i infile.json -o outfile.poly
```
