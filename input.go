package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		fmt.Fprintf(res, "%v,%v,%v\n", t.Format(time.RFC3339), m.Key, m.Value)
	}
}
