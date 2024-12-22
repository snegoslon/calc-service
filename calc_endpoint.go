package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalcEndpoint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		status := http.StatusMethodNotAllowed
		http.Error(writer, http.StatusText(status), status)
		return
	}

	var req Request
	decodeError := json.NewDecoder(request.Body).Decode(&req)
	if decodeError != nil {
		status := http.StatusBadRequest
		http.Error(writer, http.StatusText(status), status)
	}

	result, err := Calc(req.Expression)
	if err != nil {
		if err == ErrorInvalidExpression || err == ErrorInvalidCharacter {
			http.Error(writer, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		} else {
			log.Println("Error:", err)
			http.Error(writer, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		}
		return
	}

	resultStr := fmt.Sprintf("%f", result)

	response := Response{Result: resultStr}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}
