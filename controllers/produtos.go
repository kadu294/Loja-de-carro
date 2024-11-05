package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/kadu294/Loja-de-carro/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsCarros := models.BuscaTodosOsCarros()
	temp.ExecuteTemplate(w, "Index", todosOsCarros)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		modelo := r.FormValue("modelo")
		ano := r.FormValue("ano")
		marca := r.FormValue("marca")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		anoConvertidaParaInt, err := strconv.Atoi(ano)
		if err != nil {
			log.Println("Erro na conversão do ano:", err)
		}

		models.CriaNovoCarro(modelo, anoConvertidaParaInt, marca, descricao, precoConvertidoParaFloat)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoCarro := r.URL.Query().Get("id")
	models.DeletaCarro(idDoCarro)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoCarro := r.URL.Query().Get("id")
	carro := models.EditaCarro(idDoCarro)
	temp.ExecuteTemplate(w, "Edit", carro)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		modelo := r.FormValue("modelo")
		ano := r.FormValue("ano")
		marca := r.FormValue("marca")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do id:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preco:", err)
		}

		anoConvertidaParaInt, err := strconv.Atoi(ano)
		if err != nil {
			log.Println("Erro na conversão do ano:", err)
		}

		models.AtualizaCarro(idConvertidaParaInt, modelo, anoConvertidaParaInt, marca, descricao, precoConvertidoParaFloat)
	}
	http.Redirect(w, r, "/", 301)
}
