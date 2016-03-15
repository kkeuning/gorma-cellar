#! /bin/bash

set -e

rm -f ./gorma-cellar
go build .
./gorma-cellar
