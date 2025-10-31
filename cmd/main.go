package main


import(
	"log"
	"net/http"
	"status/internal/db"
	"status/internal/models"
	"status/internal/handlers"
	"status/internal/repository"
)


func main(){
	conn := db.Connect()

	conn.AutoMigrate(&models.Service{})
	log.Println("Migracion completada")

	repo := repository.ServiceRepository{DB:conn}

	AdminHandler := handlers.AdminHandler{
		Repo : &repo,
	}

	http.HandleFunc("/admin/services", AdminHandler.ServicesHandler)


	
	http.ListenAndServe(":8080", nil)
}