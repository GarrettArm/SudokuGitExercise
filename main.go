package main

import (
	"bufio"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed *.gohtml
var tpls embed.FS
var templates = template.Must(template.ParseFS(tpls, "*"))

type Board struct {
	Values     []string
	Chunks     [9][]string
	SourceFile string
}

func (b *Board) Chunk() error {
	for i := 0; i < 9; i++ {
		rowStart, rowEnd := i*9, i*9+9
		b.Chunks[i] = b.Values[rowStart:rowEnd]
	}
	return nil
}

func (b *Board) Load() error {
	f, err := os.Open(b.SourceFile)
	if err != nil {
		return err
	}
	defer f.Close()
	var output []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	b.Values = output
	return nil
}

func (b *Board) ParseForm(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	var tempValues []string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			className := fmt.Sprintf("cell-%v%v", i, j)
			tempValues = append(tempValues, r.FormValue(className))

		}
	}
	b.Values = tempValues
	err = b.Chunk()
	if err != nil {
		return err
	}
	return nil
}

func (b *Board) Save() error {
	text := strings.Join(b.Values, "\n")
	err := os.WriteFile(b.SourceFile, []byte(text), 0666)
	if err != nil {
		return err
	}
	return nil
}

func sudokuHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var b Board
		b.SourceFile = filepath.Join("codebase", "board.txt")
		err := b.Load()
		if err != nil {
			log.Fatal(err)
		}
		err = b.Chunk()
		if err != nil {
			log.Fatal(err)
		}
		templates.ExecuteTemplate(w, "board.gohtml", b.Chunks)
	case "POST":
		var b Board
		b.SourceFile = filepath.Join("codebase", "board.txt")
		err := b.ParseForm(r)
		if err != nil {
			log.Fatal(err)
		}
		err = b.Save()
		if err != nil {
			log.Fatal(err)
		}
		err = b.Chunk()
		if err != nil {
			log.Fatal(err)
		}
		templates.ExecuteTemplate(w, "board.gohtml", b.Chunks)
	}
}

func main() {
	fmt.Printf("Serving app at http://localhost:%v\n\n", 8100)
	http.HandleFunc("/", sudokuHandler)
	log.Fatal(http.ListenAndServe(":8100", nil))
}
