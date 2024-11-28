package response

import (
	"encoding/json"
	"net/http"
)

type ResponseWrapper struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func HandleSuccess(resp http.ResponseWriter, data interface{}) {
	returnData := ResponseWrapper{
		Data:    data,
		Message: "Success",
		Success: true,
	}

	jsonData, err := json.Marshal(returnData)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Something when wrong"))
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

func HandleError(resp http.ResponseWriter, status int, msg string) {
	errs := ResponseWrapper{
		Success: false,
		Message: msg,
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(errs)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, something error"))
	}
}