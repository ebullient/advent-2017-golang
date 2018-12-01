#!/bin/bash

mkdir -p /go/src/advent/2017
mkdir -p /go/src/advent/2018

cd /go/src/advent
go build -i

cd 2017
for x in $(ls *.go); do
  gofmt -w $x
done
go test
cd ..

cd 2018
for x in $(ls *.go); do
  gofmt -w $x
done
go test
cd ..

go install advent
