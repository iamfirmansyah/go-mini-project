package controllers

import (
	"go-pos/app"
	"go-pos/entity"
	"go-pos/entity/model"
	"go-pos/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func GetAll(w http.ResponseWriter, r *http.Request) {

	var data []model.Product
	app.Instance.Model(&model.Product{}).Scan(&data)

	response := entity.WebResponse{
		Code:   200,
		Status: "success",
		Data:   data,
	}
	helper.Response(w, response)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	var data model.Product

	id := r.URL.Query()["id"]

	app.Instance.First(&data, id)

	if data.Name == "" {
		response := entity.WebResponse{
			Code:   204,
			Status: "NOT_FOUND",
		}
		helper.Response(w, response)
		return
	}

	response := entity.WebResponse{
		Code:   200,
		Status: "success",
		Data:   data,
	}
	helper.Response(w, response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	var product model.Product

	helper.ReadFromRequestBody(r, &product)

	// ToDo: make validator
	err := validate.Struct(product)

	errors := helper.FormatValidationError(err)

	if errors != nil {
		response := entity.WebResponse{
			Code:   401,
			Status: "ERROR_VALIDATION",
			Data:   errors,
		}
		helper.Response(w, response)
		return
	}

	app.Instance.Create(&product)

	response := entity.WebResponse{
		Code:   200,
		Status: "success",
		Data:   product,
	}
	helper.Response(w, response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var data model.Product

	id := r.URL.Query()["id"]

	app.Instance.First(&data, id)

	if data.Name == "" {
		response := entity.WebResponse{
			Code:   204,
			Status: "NOT_FOUND",
		}
		helper.Response(w, response)
		return
	}

	helper.ReadFromRequestBody(r, &data)

	// ToDo: make validator

	app.Instance.Save(&data)

	response := entity.WebResponse{
		Code:   200,
		Status: "success",
		Data:   data,
	}

	helper.Response(w, response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var data model.Product

	id := r.URL.Query()["id"]

	app.Instance.First(&data, id)

	app.Instance.Delete(&data)

	response := entity.WebResponse{
		Code:   204,
		Status: "success",
	}

	helper.Response(w, response)
}
