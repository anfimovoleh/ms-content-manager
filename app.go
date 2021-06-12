package app

import "fmt"

type API struct{}

func New() *API {
	return &API{}
}

func (a API) Start() {
	fmt.Println("Hello world")
}
