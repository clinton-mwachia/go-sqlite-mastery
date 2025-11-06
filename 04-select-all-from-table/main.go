package main

import (
	"database/sql"
	"fmt"

	"github.com/clinton-mwachia/go-sqlite-mastery/04-select-all-from-table/models"
	"github.com/clinton-mwachia/go-sqlite-mastery/04-select-all-from-table/utils"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	// connect to the SQLite database
	db, err := sql.Open("sqlite", "./test.db?_pragma=foreign_keys(1)")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// create the countries table
	_, err = utils.CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Table countries was created successfully.")

	// create a new country
	country := &models.Country{
		Name:       "Kenya",
		Population: 329064917,
		Area:       9826675,
	}

	// insert the country
	countryId, err := utils.Insert(db, country)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print the inserted country
	fmt.Printf(
		"The country %s was inserted with ID:%d\n",
		country.Name,
		countryId,
	)

	// find all countries
	countries, err := utils.FindAll(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range countries {
		fmt.Printf("%s\n", c.Name)
	}
}
