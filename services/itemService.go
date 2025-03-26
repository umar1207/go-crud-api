package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/umar1207/go-rest-api/models"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Items)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemId, _ := strconv.Atoi(r.PathValue("id"))
	for _, item := range models.Items {
		if item.ID == itemId {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = len(models.Items) + 1
	models.Items = append(models.Items, item)
	w.WriteHeader(http.StatusCreated)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemId, _ := strconv.Atoi(r.PathValue("id"))
	for index, item := range models.Items {
		if item.ID == itemId {
			models.Items = append(models.Items[:index], models.Items[index+1:]...)
			var updatedItem models.Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = itemId
			models.Items = append(models.Items, updatedItem)
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemId, _ := strconv.Atoi(r.PathValue("id"))
	for index, item := range models.Items {
		if item.ID == itemId {
			models.Items = append(models.Items[:index], models.Items[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
