package main

import (
	"net/http"

	"github.com/Go-Blog-Proje/admin/helpers"
	admin_models "github.com/Go-Blog-Proje/admin/models"
	"github.com/Go-Blog-Proje/config"
)

func init() {
	helpers.LoadEnvVariable()
}
func main() {

	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()

	http.ListenAndServe(":8080", config.Routes())
}
