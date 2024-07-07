package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/horacioskrp/kitoo/models"
)

func TestUSerServiceHander(t *testing.T) {
	userStore := &mockUSerStore{}
	handler := NewHandler(userStore)

	t.Run("faild if user payload is invalid", func(t *testing.T) {
		payload := models.RegisterUserPayload{
			Firstname: "ubuyah",
			Lastname:  "oroshimaru",
			Username:  "jeskrp",
			Email:     "jegskrp@gmail.com",
			Password:  "98729859@play",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadGateway, rr.Code)
		}
	})
}

type mockUSerStore struct{}

// GetUSerByUSername implements models.UserStore.
func (m *mockUSerStore) GetUSerByUSername(username string) (*models.User, error) {
	return nil, nil
}

func (m *mockUSerStore) GetUserByEmail(email string) (*models.User, error) {
	return nil, nil
}
func (m *mockUSerStore) GetUSerByID(id int) (*models.User, error) {
	return nil, nil
}

func (m *mockUSerStore) CReateUSer(models.User) error {
	return nil
}
