package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Cell struct {
	X           int
	Y           int
	Class       string
	DisplayChar string
}

type PageData struct {
	Grid []Cell
}

const (
	GridWidth  = 15
	GridHeight = 15
)

var grid []Cell

func main() {
	initializeGrid()
	http.HandleFunc("/", handler)
	http.HandleFunc("/left", handleLeft)
	http.HandleFunc("/down", handleDown)
	http.HandleFunc("/up", handleUp)
	http.HandleFunc("/right", handleRight)

	println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Grid: grid,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func handleLeft(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "<span class='cell'>move left</span>")
	}
}

func handleDown(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "<span class='cell'>move down</span>")
	}
}

func handleUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "<span class='cell'>move up</span>")
	}
}

func handleRight(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "<span class='cell'>move right</span>")
	}
}

func initializeGrid() {
	grid = make([]Cell, GridWidth*GridHeight)
	for y := 0; y < GridHeight; y++ {
		for x := 0; x < GridWidth; x++ {
			index := y*GridWidth + x
			grid[index] = Cell{
				X:           x,
				Y:           y,
				Class:       "",
				DisplayChar: "•",
			}
		}
	}
}
