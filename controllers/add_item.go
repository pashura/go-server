package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../models"
)

type reqItem struct {
	Message string
}

type resItem struct {
	id      int
	message string
}

// AddItem add new item
func AddItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var ri reqItem
		err := decoder.Decode(&ri)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		newItem := models.ItemsStoreMapped.AddItem(ri.Message)
		fmt.Println(newItem)
		idStr := strconv.Itoa(newItem.ID)
		w.Write([]byte("{\"id\":" + idStr + ", \"id\":" + newItem.Message + "}"))
	}
	if r.Method == "GET" {
		fmt.Fprintf(w, "Use POST for addidng items") // send data to client side
	}
}
