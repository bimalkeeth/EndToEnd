package main

import "fmt"

type Cleaner interface {
	Clean() bool
}

type Eraser interface {
	Erase() bool
}

type Destoyer interface {
	Cleaner
	Eraser
}

type PPS = **string
type WebController struct{}

func (wc *WebController) GetName() string {
	return "Web Controller"
}

type Indexer interface {
	Index()
}
type AppController struct {
	*WebController
	Indexer
}

type IndexString string

func (hs IndexString) Index() {
	fmt.Println("Index page")
}

func main() {

	ac := new(AppController)
	fmt.Println(ac.WebController.GetName())
	fmt.Println(ac.GetName())
	ac = &AppController{new(WebController), IndexString("")}
	ac.Index()

}
