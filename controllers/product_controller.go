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
		Data: data,
	}
	helper.Response(w, response, 200)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	var data model.Product

	id := r.URL.Query()["id"]

	app.Instance.First(&data, id)

	if data.Name == "" {
		response := entity.WebResponse{
			Data: nil,
		}
		helper.Response(w, response, 204)
		return
	}

	response := entity.WebResponse{
		Data: data,
	}
	helper.Response(w, response, 200)
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
			Data: errors,
		}
		helper.Response(w, response, 400)
		return
	}

	app.Instance.Create(&product)

	response := entity.WebResponse{
		Data: product,
	}
	helper.Response(w, response, 200)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var data model.Product

	id := r.URL.Query()["id"]

	app.Instance.First(&data, id)

	if data.Name == "" {
		response := entity.WebResponse{
			Data: nil,
		}
		helper.Response(w, response, 204)
		return
	}

	helper.ReadFromRequestBody(r, &data)

	// ToDo: make validator

	app.Instance.Save(&data)

	response := entity.WebResponse{
		Data: data,
	}

	helper.Response(w, response, 200)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var data model.Product

	id := r.URL.Query()["id"]

	app.Instance.First(&data, id)

	app.Instance.Delete(&data)

	helper.Response(w, nil, 204)
}
