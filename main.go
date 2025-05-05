package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id    int16  `json:"id" validate:"required"`
	Name  string `json:"name" validate:"len=20"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	http.HandleFunc("/user", UserHandler)
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		panic(err)
	}
}

func WriteJson(w http.ResponseWriter, staus int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "method not allowed",
		})
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	validate := validator.New()
	if errValidate := validate.Struct(user); errValidate != nil {
		errs := errValidate.(validator.ValidationErrors)
		var field string
		for _, fieldErr := range errs {
			field = fieldErr.Field()
		}

		WriteJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": field,
		})
		return
	}

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("user %v", user)

	WriteJson(w, http.StatusOK, map[string]any{
		"ok": true,
	})
}
