#!/bin/bash

cd /go/src/advent
go build -i

cd days
for x in $(ls *.go); do
  gofmt -w $x
done
go test
cd ..

go install advent
