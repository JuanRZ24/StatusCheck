package services



import (
	"log"
	"net/http"
	"time"
	"status/internal/repository"
)


type MonitorService struct {
	Repo *repository.ServiceRepository
}



func (m *MonitorService) CheckAllServices(){
	services, err := m.Repo.GetAll()

	if err != nil {
		log.Println("Error al obtener servicios:", err)
		return
	}

	for _, s := range services {
		if !s.Enabled {
			continue
		}

		

		start := time.Now()
		resp, err  := http.Get(s.URL)
		latency := time.Since(start).Milliseconds()

		if err != nil {
			s.Status = "DOWN"
			s.Error = err.Error()
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == s.ExpectedStatus {
				s.Status = "UP"
				s.Error	 = ""
			} else {
				s.Status = "DOWN"
				s.Error = "Codigo inesperado"
			}
		}
		s.LatencyMS = int(latency)
		m.Repo.Update(&s)
		log.Printf("🔍 [%s] %s → %s (%d ms)\n", s.Name, s.URL, s.Status, latency)
	}

	
}