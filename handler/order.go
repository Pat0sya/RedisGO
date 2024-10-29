package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create called")
}
func (o *Order) List(w http.ResponseWriter, r *http.Request) {

	fmt.Println("List called")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GetByID called")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("UpdateByID called")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("DeleteByID called")
}
