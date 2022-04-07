package res

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// vars errors
var invalidJson = errors.New("не валидный json")
var encoderJson = errors.New("не удалось закодировать json")

// Bind - Декодирует тело запросы в структуру
func Bind(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return invalidJson
	}
	return nil
}

// JSON - Кодирует и отправляет данные в json-формате
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", encoderJson)
	}
}

// ERROR - Кодирует и отправляет тело ошибки
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
