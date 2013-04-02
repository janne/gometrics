package main

import (
	"bufio"
	"fmt"
	"github.com/knieriem/markdown"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", readme)
	fmt.Println("Listening at port", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func readme(res http.ResponseWriter, req *http.Request) {
	parser := markdown.NewParser(&markdown.Extensions{Smart: true})
	file, err := os.Open("README.md")
	if err != nil {
		fmt.Fprintln(res, err)
	} else {
		w := bufio.NewWriter(res)
		parser.Markdown(file, markdown.ToHTML(w))
		w.Flush()
	}
}
