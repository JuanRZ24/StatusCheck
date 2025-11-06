package main


import(
	"log"
	"net/http"
	"status/internal/db"
	"status/internal/models"
	"status/internal/handlers"
	"status/internal/repository"
	"status/internal/services"
	"time"
)


func main(){
	conn := db.Connect()

	conn.AutoMigrate(&models.Service{})
	log.Println("Migracion completada")

	repo := repository.ServiceRepository{DB:conn}

	AdminHandler := handlers.AdminHandler{
		Repo : &repo,
	}
	monitor := services.MonitorService{Repo: &repo}

// 🛰️ Ejecutar el monitor en segundo plano cada 60 segundos
		go func() {
			for {
				monitor.CheckAllServices()
				time.Sleep(1 * time.Minute)
			}
		}()

	http.HandleFunc("/admin/services", AdminHandler.ServicesHandler)


	
	http.ListenAndServe(":8080", nil)
}