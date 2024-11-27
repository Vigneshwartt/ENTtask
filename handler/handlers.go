package handler

import (
	"encoding/json"
	"net/http"
	"routersmux/drivers"
	"routersmux/helpers"
	"routersmux/models"
	"routersmux/repository"
	"strconv"

	"github.com/gorilla/mux"
)

type CommonResponse struct {
	Message    string `json:"message"`
	Statuscode int    `json:"code"`
	Errors     string `json:"errors,omitempty"`
	Data       any    `json:"data,omitempty"`
}

var (
	MainConnection *repository.DataConnection
)

func init() {
	MainConnection = &repository.DataConnection{DB: drivers.ConnectingDatabase()}
}

func CreateUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user models.Employee

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := MainConnection.NewCreateUserDetails(&user)
	if value.Error != nil {
		json.NewEncoder(w).Encode(value.Error)
		return
	}

	err = json.NewEncoder(w).Encode(CommonResponse{
		Data:       user,
		Message:    "Data Created",
		Statuscode: http.StatusCreated,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func GetMultipleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user []models.Employee

	value := MainConnection.NewGetUser(&user)
	if value.Error != nil {
		json.NewEncoder(w).Encode(value.Error.Error())
		helpers.ErrorData.Println("Error occured while getting values")
		return
	}

	for _, datas := range user {
		err := json.NewEncoder(w).Encode(&datas)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			helpers.WarnData.Println("Error occured getting values")
			return
		}
	}
	err := json.NewEncoder(w).Encode(CommonResponse{
		Message:    "Data Created",
		Statuscode: http.StatusCreated,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	var user models.Employee
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helpers.ErrorData.Println("Error occured while getting values")
		return
	}

	data := MainConnection.NewSingleuser(id, &user)
	if data.Error != nil {
		json.NewEncoder(w).Encode(data.Error.Error())
		return
	}
	err = json.NewEncoder(w).Encode(CommonResponse{
		Data:       user,
		Message:    "Data Retrevied Sucessfully",
		Statuscode: http.StatusAccepted,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	var user models.Employee
	params := mux.Vars(r)
	val := params["employee_id"]
	ids, err := strconv.Atoi(val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := MainConnection.NewUpdateUsers(ids, &user)

	if data.Error != nil {
		json.NewEncoder(w).Encode(data.Error.Error())
		return
	} else if data.RowsAffected == 0 {
		json.NewEncoder(w).Encode(CommonResponse{
			Message:    "No rows affected",
			Statuscode: http.StatusAccepted,
		})
	} else {
		json.NewEncoder(w).Encode(CommonResponse{
			Message:    "updated rows successfully",
			Statuscode: http.StatusAccepted,
		})
	}

}

func Deleteuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user models.Employee

	id := mux.Vars(r)["id"]

	ids, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.ErrorData.Println("Error occured while getting values")
		return
	}

	data := MainConnection.NewDeleteUser(ids, &user)

	if data.Error != nil {
		json.NewEncoder(w).Encode(data.Error.Error())
		return
	} else if data.RowsAffected == 0 {
		json.NewEncoder(w).Encode(CommonResponse{
			Message:    "No Rows Affected",
			Statuscode: http.StatusBadRequest,
		})
	} else {
		json.NewEncoder(w).Encode(CommonResponse{
			Message:    "Rows Affected",
			Statuscode: http.StatusOK,
		})
	}
}
