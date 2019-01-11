
# for mac os
echo building binary for mac os ...
go build -o geojson2poly main.go
zip -m geojson2poly-mac.zip geojson2poly

# for linux
echo building binary for linux ...
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o geojson2poly main.go
zip -m geojson2poly-linux-x64.zip geojson2poly

# for windows
echo building binary for windows ...
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o geojson2poly.exe main.go
zip -m geojson2poly-windows-x64.zip geojson2poly.exe

echo done!
