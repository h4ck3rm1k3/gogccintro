package models
import (
	/* "flag"	 */
	"fmt"
	"log"
	"testing"
	/* 			"time" */
	_ "github.com/mattn/go-sqlite3"	
	"github.com/knq/dburl"
	"github.com/mattes/migrate/migrate"
	//"github.com/knq/xo/examples/booktest/sqlite3/models"
	_ "github.com/mattes/migrate/driver/sqlite3"
	)


func TestMain(t *testing.T){

	db, err := dburl.Open("file:test.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	
	// use synchronous versions of migration functions ...
	err2, ok := migrate.UpSync("sqlite3://test.sqlite3", "../migrations")
	if !ok {
		fmt.Println("Oh no ...")
		log.Fatal(err2)
		// do sth with allErrors slice
	}
	
	
	f := GccTuParserSourcefile{
		Filename: "test.c.tu",
	}
	f.Save(db)
	if err != nil {
		log.Fatal(err)
	}
	
}
