package main

import (
	"flag"
	"os"
	"strings"
	"path/filepath"
    "fmt"
    "net/http"
	"io/ioutil"
)
type Page struct {
    Title string
    Body  []byte
}
func handler(w http.ResponseWriter, r *http.Request) {
	path := dir + r.URL.Path[0:]
	page,err := loadPage(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Write(page.Body)
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1)
}
var dir string
func main() {
	port := flag.String("port","8080","默认8080端口")
	flag.Parse()
	dir = getCurrentDirectory()
    http.HandleFunc("/", handler)
	fmt.Println("侦听 http://localhost:"+*port+" 中...")
    http.ListenAndServe(":"+*port, nil)
}

func loadPage(path string) (*Page, error) {
    body, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    return &Page{Title: path, Body: body}, nil
}