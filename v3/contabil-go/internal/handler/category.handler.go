package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/internal/model"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Get categories from database

	categories, err := model.GetCategories()

	if err != nil {
		logging.FailedToFindOnDB("All Categories", "", err.Inner)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to get categories")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
		return
	}

}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	var cat domain.Category

	json.NewDecoder(r.Body).Decode(&cat)

	err := model.CreateCategory(cat)

	if err != nil {

		logging.GenericError("Fail to create a category", err.Inner)
		w.WriteHeader(err.HtmlStatus)
		fmt.Fprintln(w, err.Inner.Error())
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		return
	}

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	var cat domain.Category

	json.NewDecoder(r.Body).Decode(&cat)

	err := model.UpdateCategory(cat)

	if err != nil {

		logging.GenericError("Fail to update a category", err.Inner)
		w.WriteHeader(err.HtmlStatus)
		fmt.Fprintln(w, err.Inner.Error())
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		return
	}

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	err := model.DeleteCategory(id)

	if err != nil {

		logging.GenericError("Fail to delete a category", err.Inner)
		w.WriteHeader(err.HtmlStatus)
		fmt.Fprintln(w, err.Inner.Error())
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}

}
