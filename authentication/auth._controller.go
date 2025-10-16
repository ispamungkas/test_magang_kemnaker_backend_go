package authentication

import (
	"KemnakerMagang/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type userBody struct {
	Name      string `json:"name" form:"name" query:"name" validate:"required"`
	Domicile string `json:"domicilie" form:"domicilie" query:"domicilie" validate:"required"`
}

type AuthenticaionService interface {
	GetAllUser(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
}

type authenticaionService struct {
	db *sql.DB
}

func NewAuthenticationService(
	db *sql.DB,
) AuthenticaionService {
	return &authenticaionService{
		db: db,
	}
}

func (as *authenticaionService) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := as.db.Query("SELECT * FROM users")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse, err := json.Marshal(models.ErrorResponse{Status_Code: 401, Message: "Bad Request"})

		if err != nil {
			log.Println("Error parsing marshal")
		}

		w.Write(errResponse)

	}

	defer res.Close()

	var response []*models.User

	for res.Next() {
		temp := new(models.User)
		err := res.Scan(&temp.ID, &temp.Name, &temp.Domicile)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errResponse, err := json.Marshal(models.ErrorResponse{Status_Code: 401, Message: "Bad Request"})

			if err != nil {
				log.Println("Error parsing marshal")
			}

			w.Write(errResponse)
		}
		response = append(response, temp)
	}

	successResponse := models.SuccessResponse{
		Status_Code: 200,
		Message:     "Data sucessfuly fetch",
		Data:        response,
	}

	jresponse, err := json.Marshal(successResponse)

	if err != nil {
		log.Println("Error parsing marshal")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jresponse)

}

func (as *authenticaionService) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body userBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	err = validate.Struct(body)

	if err != nil {
		var errMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errMessages = append(errMessages, fmt.Sprintf("Field %s: %s", e.Field(), e.Tag()))
		}
		errResponse, err := json.Marshal(models.ErrorResponse{Status_Code: 401, Message: strings.Join(errMessages, ",")})

		if err != nil {
			log.Println("Error parsing")
		}

		w.Write(errResponse)
	}

	query := fmt.Sprintf("INSERT INTO USERS (name, domicile) VALUES ('%s', '%s')", body.Name, body.Domicile)
	
	res, err := as.db.Query(query)
	if err != nil || res == nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse, err := json.Marshal(models.ErrorResponse{Status_Code: 401, Message: "Bad Request"})

		if err != nil {
			log.Println("Error parsing marshal")
		}

		w.Write(errResponse)

	}
	
	defer res.Close()

	successResponse := models.SuccessResponse{
		Status_Code: 200,
		Message:     "Data sucessfuly added",
		Data:        body,
	}

	jresponse, err := json.Marshal(successResponse)

	if err != nil {
		log.Println("Error parsing marshal")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jresponse)

}

func (as *authenticaionService) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	query := fmt.Sprintf("SELECT * FROM users WHERE id = %v", id)
	res, err := as.db.Query(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse, err := json.Marshal(models.ErrorResponse{Status_Code: 401, Message: "Bad Request"})

		if err != nil {
			log.Println("Error parsing marshal")
		}

		w.Write(errResponse)

	}

	defer res.Close()

	var response []*models.User

	for res.Next() {
		temp := new(models.User)
		err := res.Scan(&temp.ID, &temp.Name, &temp.Domicile)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errResponse, err := json.Marshal(models.ErrorResponse{Status_Code: 401, Message: "Bad Request"})

			if err != nil {
				log.Println("Error parsing marshal")
			}

			w.Write(errResponse)
		}
		response = append(response, temp)
	}

	successResponse := models.SuccessResponse{
		Status_Code: 200,
		Message:     "Data sucessfuly fetch",
		Data:        response,
	}

	jresponse, err := json.Marshal(successResponse)

	if err != nil {
		log.Println("Error parsing marshal")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jresponse)

}
