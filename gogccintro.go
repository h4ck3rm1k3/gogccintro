// Copyright 2017, James Michael Du Pont. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

// gogccintro.go [created: Wed,  1 Mar 2017]

package main

import "fmt"
import (
	_ "database/sql"
	_ "github.com/mattn/go-sqlite3"	
	_ "github.com/knq/dburl"
	_ "github.com/mattes/migrate/migrate"
	_ "github.com/mattes/migrate/driver/sqlite3"
	)

func main() {
    fmt.Println("Hello, GCC!")
}
