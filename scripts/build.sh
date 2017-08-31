#!/usr/bin/env bash

SRC="intel-nvdimm/type"
DST="intel-nvdimm"
glide install
go generate intel-nvdimm/type/constants.go
cp $SRC/* $DST
go build
