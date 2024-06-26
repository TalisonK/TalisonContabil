package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/TalisonK/TalisonContabil/internal/config"
	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/internal/database"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/internal/model"
)

type TestUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type TestUserId struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func TestMain(m *testing.M) {

	err := config.Load()

	if err != nil {
		logging.GenericError("Erro ao carregar configurações", err)
		return
	}

	err = database.OpenConnectionLocal()

	if err != nil {
		logging.FailedToOpenConnection(constants.LOCAL, err)
		return
	}

	err = database.OpenConnectionCloud()

	if err != nil {
		logging.FailedToOpenConnection(constants.CLOUD, err)
		return
	}

	exitVal := m.Run()

	defer database.CloseConnections()

	os.Exit(exitVal)
}

func TestGetUsers(t *testing.T) {

	// rr := httptest.NewRecorder()
	// req := httptest.NewRequest(http.MethodGet, "/user", nil)

}

func TestCreateUser(t *testing.T) {

	body := TestUser{
		Name:     "Teste",
		Password: "123",
	}

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(body)

	if err != nil {
		logging.GenericError("Erro ao criar requisição", err)
		t.Fatalf("Erro ao criar requisição")
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/user", &b)

	CreateUser(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	response, err := io.ReadAll(rr.Body)

	if err != nil {
		logging.GenericError("Erro ao ler resposta", err)
		t.Fatalf("Erro ao ler resposta")
	}

	expected := domain.UserDTO{
		Name: "Teste",
		Role: "ROLE_USER",
	}

	var result domain.UserDTO
	err = json.Unmarshal(response, &result)

	if err != nil {
		logging.GenericError("Erro ao deserializar resposta", err)
		t.Fatalf("Erro ao deserializar resposta")
	}

	if result.Name != expected.Name {
		t.Errorf("handler returned wrong name: got %v want %v",
			result.Name, expected.Name)
	}

	if result.Role != expected.Role {
		t.Errorf("handler returned wrong role: got %v want %v",
			result.Role, expected.Role)
	}

	tagErr := model.DeleteUser(result.ID)

	if tagErr != nil {
		t.Errorf("Error while cleaning databases")
	}

}

func TestUpdateUser(t *testing.T) {

	example := domain.User{
		Name:     "testUpdate",
		Password: "123",
	}

	userInBase, tagErr := model.CreateUser(&example)

	if tagErr != nil {
		logging.GenericError("Fail to create user for update test", tagErr.Inner)
		t.Errorf("Fail to create user for update test")
	}

	defer model.DeleteUser(userInBase.ID)

	body := TestUserId{
		Id:   userInBase.ID,
		Name: "newUserName",
	}

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(body)

	if err != nil {
		logging.GenericError("Erro ao criar requisição", err)
		t.Fatalf("Erro ao criar requisição")
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/user", &b)

	UpdateUser(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	response, err := io.ReadAll(rr.Body)

	if err != nil {
		logging.GenericError("Erro ao ler resposta", err)
		t.Fatalf("Erro ao ler resposta")
	}

	expected := domain.UserDTO{
		Name: "newUserName",
	}

	var result domain.UserDTO
	err = json.Unmarshal(response, &result)

	if err != nil {
		logging.GenericError("Erro ao deserializar resposta", err)
		t.Fatalf("Erro ao deserializar resposta")
	}

	if result.Name != expected.Name {
		t.Errorf("handler returned wrong name: got %v want %v",
			result.Name, expected.Name)
	}

}

func TestDeleteUser(t *testing.T) {

	//TOFIX
	// example := domain.User{
	// 	Name:     "testUpdate",
	// 	Password: "123",
	// }

	// userInBase, err := model.CreateUser(&example)

	// if err != nil {
	// 	util.LogHandler("Fail to create user for update test",err)
	// 	t.Errorf("Fail to create user for update test")
	// }

	// u, err := url.Parse("/user")

	// if err != nil {
	// 	util.LogHandler("Fail to parse url",err)
	// 	t.Errorf("Fail to parse url")
	// }

	// q := u.Query()
	// q.Set("id", userInBase.ID)
	// u.RawQuery = q.Encode()

	// rr := httptest.NewRecorder()
	// req := httptest.NewRequest(http.MethodDelete, u.String(), nil)

	// DeleteUser(rr, req)

	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

}
