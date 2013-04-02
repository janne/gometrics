package main

import (
	"bufio"
	"fmt"
	"github.com/knieriem/markdown"
	"net/http"
	"os"
)

func ReadmeHandler(res http.ResponseWriter, req *http.Request) {
	parser := markdown.NewParser(&markdown.Extensions{Smart: true})
	file, err := os.Open("README.md")
	if err != nil {
		fmt.Fprintln(res, err)
	} else {
		w := bufio.NewWriter(res)
		w.WriteString("<html>")
		w.WriteString(`<head><link href="http://kevinburke.bitbucket.org/markdowncss/markdown.css" rel="stylesheet"></link></head>`)
		w.WriteString(`<body><a href="https://github.com/janne/gometrics"><img style="position: absolute; top: 0; right: 30px; border: 0;" src="https://github.com/jamesflorentino/fork-ribbons/raw/master/ribbons/red-white.png" alt="Fork me on GitHub"></a>`)
		parser.Markdown(file, markdown.ToHTML(w))
		w.WriteString("</body></html>")
		w.Flush()
	}
}
