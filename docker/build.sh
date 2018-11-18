#!/bin/bash

cd /go/src/advent
go build -i

cd days
go test
cd ..

go install advent
