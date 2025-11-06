package utils

import (
	"database/sql"

	"github.com/clinton-mwachia/go-sqlite-mastery/04-select-all-from-table/models"
	_ "github.com/glebarez/go-sqlite"
)

func CreateTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS countries (
        id INTEGER PRIMARY KEY,
        name     TEXT NOT NULL,
        population INTEGER NOT NULL,
        area INTEGER NOT NULL
    );`

	return db.Exec(sql)
}

func Insert(db *sql.DB, c *models.Country) (int64, error) {
	sql := `INSERT INTO countries (name, population, area) 
            VALUES (?, ?, ?);`
	result, err := db.Exec(sql, c.Name, c.Population, c.Area)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func FindAll(db *sql.DB) ([]models.Country, error) {
	sql := `SELECT * FROM countries`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		c := &models.Country{}
		err := rows.Scan(&c.Id, &c.Name, &c.Population, &c.Area)
		if err != nil {
			return nil, err
		}
		countries = append(countries, *c)
	}
	return countries, nil
}
