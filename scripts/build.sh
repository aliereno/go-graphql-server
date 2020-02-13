#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
outputPath="build"
app="gql-server"
output="$outputPath/$app"
src="$srcPath/$app/$pkgFile"

printf "\nBuilding: $app\n"
time go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: $app\n\n"

# WINDOWS 64Bit BUILD
#env GOOS=windows GOARCH=amd64 go build -o build/gql-server cmd/gql-server/main.go