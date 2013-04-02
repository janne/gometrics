package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Metric struct {
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
		fmt.Fprintln(res, m.Key, "=", m.Value)
	}
}
