// Copyright 2017, James Michael Du Pont. All rights reserved.
// Use of this source code is governed by a AGPL-style
// license that can be found in the LICENSE file.

// gogccintro.go [created: Wed,  1 Mar 2017]

package main


import (
	"fmt"
	"os"
	"log"
//	"io/ioutil"
	"flag"
	_ "database/sql"
	_ "github.com/mattn/go-sqlite3"	
	"github.com/knq/dburl"
	_ "github.com/mattes/migrate/migrate"
	_ "github.com/mattes/migrate/driver/sqlite3"
	//"github.com/BurntSushi/toml"
//	"gopkg.in/yaml.v2"
	"github.com/jinzhu/configor"
	"github.com/h4ck3rm1k3/gogccintro/filters"
	//"filter/load_recurse"
	)

//filters := map[string]int{
//	"load_recurse": filter.load_recurse,
//}

var Config = struct {
	APPName string `default:"go gcc introspector"`

	InputDB struct {
		Path     string
	}

	Transform filter.Transform
	
	OutputDB struct {
		Path     string
	}
}{}

func main() {
	config := flag.String("file", "config.yaml", "configuration file")
	flag.StringVar(&Config.InputDB.Path, "inputdb", "", "input db name")
	flag.StringVar(&Config.OutputDB.Path, "outputdb", "", "output db name")
	flag.Parse()

	os.Setenv("CONFIGOR_ENV_PREFIX", "-")
	err := configor.Load(&Config, *config);
	if (err != nil){
		fmt.Printf("err loading %s\n", err)		
	}
	
	fmt.Println("Hello, GCC!")


	fmt.Printf("input db %s\n", Config.InputDB.Path)
	fmt.Printf("output db %s\n", Config.OutputDB.Path)
	fmt.Printf("filter %s\n", Config.Transform.Filter)

	indb, err := dburl.Open(Config.InputDB.Path)
	if err != nil {
		log.Fatal(err)
	}
	outdb, err := dburl.Open(Config.OutputDB.Path)
	if err != nil {
		log.Fatal(err)
	}
	
	filter.DoTransform(indb, outdb, Config.Transform)
	
	//configBytes, err := yaml.Marshal(&Config)
	//fmt.Printf("output %#v\nerr:%#v\n", configBytes, err)
	//ioutil.WriteFile("test.yaml", configBytes, 0644)
}
