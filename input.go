package main

import (
	"database/sql"
	"encoding/json"
	"github.com/bmizerany/pq"
	"io/ioutil"
	"net/http"
	"os"
	//"time"
)

type Metric struct {
	Time  string
	Key   string
	Value float32
}

func dbConnect() (db *sql.DB, err error) {
	dbUrl := os.Getenv("DATABASE_URL")
	source, err := pq.ParseURL(dbUrl)
	if err != nil {
		db, err = sql.Open("postgres", source)
	}
	return
}

func dbInsert(db *sql.DB, query string, args ...interface{}) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	_, err = tx.Exec(query, args)
	if err != nil {
		return
	}
	err = tx.Commit()
	return
}

func InputHandler(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	var metrics []Metric
	err = json.Unmarshal(body, &metrics)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	db, err := dbConnect()
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	for _, m := range metrics {
		//t, err := time.Parse(time.RFC3339, m.Time)
		//if err != nil {
		//	http.Error(res, err.Error(), 400)
		//	return
		//}
		// err = dbInsert(db, "INSERT INTO metrics (time, key, value) VALUES (t.Format(time.RFC3339), m.Key, m.Value)")
		err = dbInsert(db, "INSERT INTO metrics (key, value) VALUES (?, ?)", m.Key, m.Value)
		if err != nil {
			http.Error(res, err.Error(), 500)
		}
	}
}
