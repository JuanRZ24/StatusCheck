package repository

import(
	"status/internal/models"
	"gorm.io/gorm"
)


type ServiceRepository struct {
	DB *gorm.DB
}

func (r *ServiceRepository) Create(service *models.Service) error {
	return r.DB.Create(service).Error
}


func (r *ServiceRepository) GetAll() ([]models.Service, error){
	var services []models.Service
	err := r.DB.Find(&services).Error
	return services, err
}


func (r *ServiceRepository) Update(s *models.Service) error {
	return r.DB.Save(s).Error
}