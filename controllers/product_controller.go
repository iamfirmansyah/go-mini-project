package controllers

import (
	"encoding/json"
	"go-pos/app"
	"go-pos/entity"
	"go-pos/entity/model"
	"go-pos/helper"
	"net/http"
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
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

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

	json.NewDecoder(r.Body).Decode(&data)
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
