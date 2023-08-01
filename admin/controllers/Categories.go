package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Go-Blog-Proje/admin/helpers"
	"github.com/Go-Blog-Proje/admin/models"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Categories struct{}

func (categories Categories) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("categories/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (categories Categories) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	categoryTitle := r.FormValue("category-title")
	categorySlug := slug.Make(categoryTitle)

	models.Category{Title: categoryTitle, Slug: categorySlug}.Add()

	helpers.SetAlert(w, r, "Kayıt Başarıyla Eklendi")
	http.Redirect(w, r, "/admin/kategoriler", http.StatusSeeOther)
}

func (categories Categories) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	category := models.Category{}.Get(params.ByName("id"))
	category.Delete()
	helpers.SetAlert(w, r, "Kayıt Başarıyla Silindi...")
	http.Redirect(w, r, "/admin/kategoriler", http.StatusSeeOther)
}
