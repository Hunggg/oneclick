package main

import (
	// "oneclick/cmd"

	"fmt"
	"log"
	"oneclick/config"
	cockroachdb "oneclick/services/query/categories/cockroachDB"
)

// "oneclick/cmd"


func main() {
	// cmd.Execute()
	db, err := config.NewCockroachDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	var c *cockroachdb.CockroachDB
	c, err = cockroachdb.NewCockroachDB(db)
	if err != nil {
		log.Fatal(err)
	}
	cate, _ := c.GetCategoryById(1)
	fmt.Println(cate)
}