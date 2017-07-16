package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
)

var (
	hostPort = flag.String("hostport", "localhost:8080", "server host and port")
	logPath  = flag.String("log", "", "Log file path, default is output")
)

const (
	mdDir = "md"
)

func main() {
	flag.Parse()

	if *logPath != "" {
		logOut, err := os.Open(*logPath)

		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(logOut)
	}

	buf := new(bytes.Buffer)
	files, _ := ioutil.ReadDir(mdDir)

	paths := sortFiles(files)

	for _, path := range paths {
		file, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		io.Copy(buf, file)
	}

	content := template.HTML(blackfriday.MarkdownCommon(buf.Bytes()))

	t, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		log.Fatal(err)
	}

	view := struct{ Content template.HTML }{content}
	buf = new(bytes.Buffer)

	if err := t.ExecuteTemplate(buf, "layout.html", view); err != nil {
		log.Fatal(err)
	}

	html := buf.Bytes()
	buf.Reset()

	router := httprouter.New()
	router.GET("/", httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write(html)
	}))

	http.ListenAndServe(*hostPort, router)
}

func sortFiles(files []os.FileInfo) []string {
	s := make([]string, len(files), len(files))

	for _, f := range files {
		if !f.IsDir() {
			index, err := strconv.Atoi(strings.Split(f.Name(), "_")[0])

			if err != nil {
				log.Println(err)
			}

			s[index-1] = fmt.Sprintf("%s/%s", mdDir, f.Name())
		}
	}

	return s
}
