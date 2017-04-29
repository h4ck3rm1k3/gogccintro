# gogccintro
golang interfaces for the gcc introspector

[![Build Status](https://travis-ci.org/h4ck3rm1k3/gogccintro.svg?branch=master)](https://travis-ci.org/h4ck3rm1k3/gogccintro)
[![GoDoc](https://godoc.org/github.com/h4ck3rm1k3/gogccintro?status.png)](http://godoc.org/github.com/h4ck3rm1k3/gogccintro)

## Install

    go get github.com/h4ck3rm1k3/gogccintro

## Usage

   Compile programs with these flags to get tu files

    CFLAGS=-O0 -fdump-translation-unit -save-temps
    CPPFLAGS=-O0 -fdump-translation-unit -save-temps
    CXXFLAGS=-O0 -fdump-translation-unit -save-temps


## Author

James Michael Du Pont [jamesmikedupont@gmail.com]

## Copyright & License

Copyright (c) 2017, James Michael Du Pont.
All rights reserved.
Use of this source code is governed by a AGPL-style license that can be
found in the LICENSE file.
