package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wailman24/Go-chi-starter.git/internal/models"
	"github.com/wailman24/Go-chi-starter.git/internal/services"
	"github.com/wailman24/Go-chi-starter.git/tokens"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	serv *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		serv: services.NewUserService(),
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var validate = validator.New()
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedpwd, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = hashedpwd

	err = uh.serv.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": user,
	})

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin

	var validate = validator.New()
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := uh.serv.GetUserByEmail(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !CheckPasswordHash(user.Password, res.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := tokens.CreateToken(int(res.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": user,
		"token":   token,
	})

}
