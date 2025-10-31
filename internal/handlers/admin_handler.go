package handlers

import (
	"encoding/json"
	"net/http"
	"status/internal/models"
	"status/internal/repository"
)

type AdminHandler struct {
	Repo *repository.ServiceRepository

}

func (h *AdminHandler) ServicesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetServices(w, r)
	case http.MethodPost:
		h.CreateService(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}


func (h *AdminHandler) CreateService (w http.ResponseWriter, r *http.Request) {
	var input models.Service

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}


	if input.Name == "" || input.URL == "" {
		http.Error(w, "Faltan campos obligatorios: name o url", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(&input); err != nil {
		http.Error(w,"Error interno del servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&input)
	
}

func (h *AdminHandler) GetServices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	services, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Error al obtener los servicios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}
