package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/internal/model"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	var body domain.Total

	json.NewDecoder(r.Body).Decode(&body)

	result, tagerr := model.GetExpensesByDate(body.UserID, body.Month, body.Year, true, true)

	if tagerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, logging.FailedToFindOnDB(fmt.Sprintf("Expenses from user %s", body.UserID), constants.LOCAL, tagerr.Inner), tagerr.Inner.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func CreateExpense(w http.ResponseWriter, r *http.Request) {

	var body domain.ExpenseDTO

	json.NewDecoder(r.Body).Decode(&body)

	result, tagErr := model.CreateExpenseHandler(body)

	if tagErr != nil {
		w.WriteHeader(tagErr.HtmlStatus)
		fmt.Fprintln(w, logging.GenericError("Error received while tring do create expense", tagErr.Inner), tagErr.Inner.Error())
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func CreateExpenseList(w http.ResponseWriter, r *http.Request) {

	var body domain.ExpenseDTO

	json.NewDecoder(r.Body).Decode(&body)

	tagErr := model.CreateExpenseListHandler(body)

	if tagErr != nil {
		w.WriteHeader(tagErr.HtmlStatus)
		fmt.Fprintln(w, logging.GenericError("Error received while tring do create expense list", tagErr.Inner), tagErr.Inner.Error())
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Expense with List created successfully")

}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {

	var body domain.ExpenseDTO

	json.NewDecoder(r.Body).Decode(&body)

	result, tagErr := model.UpdateExpenseHandler(body)

	if tagErr != nil {
		w.WriteHeader(tagErr.HtmlStatus)
		fmt.Fprintln(w, logging.GenericError("Error received while tring do update expense", tagErr.Inner), tagErr.Inner.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	tagErr := model.DeleteExpenseHandler(id)

	if tagErr != nil {
		w.WriteHeader(tagErr.HtmlStatus)
		fmt.Fprintln(w, logging.GenericError("Error received while tring do delete expense", tagErr.Inner), tagErr.Inner.Error())
	}

	w.WriteHeader(http.StatusOK)

}
