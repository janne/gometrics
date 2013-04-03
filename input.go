package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pq"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Metric struct {
	Time  string
	Key   string
	Value float32
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
	for _, m := range metrics {
		t, err := time.Parse(time.RFC3339, m.Time)
		if err != nil {
			http.Error(res, err.Error(), 400)
			return
		}
		dbUrl := os.Getenv("DATABASE_URL")
		dataSourceName, err := pq.ParseURL(dbUrl)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		db, err := sql.Open("postgres", dataSourceName)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		rows, err := db.Query("select * from metrics")
		if err != nil {
		  http.Error(res, err.Error(), 500)
		  return
		}
		var fields []interface{}
		for rows.Next() {
		  rows.Scan(fields...)
		  err := rows.Scan()
		  if err != nil {
			  http.Error(res, err.Error(), 500)
			  return
		  }
		}
		//fmt.Fprintf(rows, err)
		fmt.Fprintf(res, "%v,%v,%v\n", t.Format(time.RFC3339), m.Key, m.Value)
	}
}
