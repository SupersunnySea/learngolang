package main

import "fmt"

type iSayBestLanguage interface {
	BestLanuage()
}

type Golang struct {
}

type Php struct {
}

func (g Golang) BestLanuage() {
	fmt.Print("golang is the best language!")
}

func (p Php) BestLanuage() {
	fmt.Print("php is the best language!")
}

func sayBestLanguage(i iSayBestLanguage) {
	i.BestLanuage()
}

func main() {
	var i iSayBestLanguage
	var golang Golang
	var php Php
	i = golang
	sayBestLanguage(i)
	i = php
	sayBestLanguage(i)
}
