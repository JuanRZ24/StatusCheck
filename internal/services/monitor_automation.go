package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"status-dashboard/internal/models"
	"status-dashboard/internal/repository"
)

type AutomationMonitor struct {
	Repo    *repository.AutomationRepository
	BaseURL string 
	APIKey  string
}

func (m *AutomationMonitor) CheckAllAutomations() {
	automations, err := m.Repo.GetAll()
	if err != nil {
		fmt.Println("❌ Error al obtener automatizaciones:", err)
		return
	}

	for _, a := range automations {
		status, msg, runAt := m.checkWorkflowStatus(a.WorkflowID)
		a.LastRunStatus = status
		a.LastRunAt = runAt
		a.ErrorMessage = msg
		m.Repo.Update(&a)

		fmt.Printf("⚙️ [%s] → %s (%s)\n", a.Name, status, runAt.Format(time.RFC822))
	}
}


func (m *AutomationMonitor) checkWorkflowStatus(id string) (string, string, time.Time) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/rest/executions?workflowId=%s&limit=1", m.BaseURL, id), nil)
	req.Header.Set("Authorization", "Bearer "+m.APIKey)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "DOWN", err.Error(), time.Time{}
	}
	defer resp.Body.Close()

	var data struct {
		Data []struct {
			Status    string    `json:"status"`
			StartedAt time.Time `json:"startedAt"`
			StoppedAt time.Time `json:"stoppedAt"`
			Error     string    `json:"error"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "DOWN", "Error parseando JSON", time.Time{}
	}

	if len(data.Data) == 0 {
		return "UNKNOWN", "Sin ejecuciones registradas", time.Time{}
	}

	exec := data.Data[0]
	if exec.Status == "error" {
		return "DOWN", exec.Error, exec.StoppedAt
	}

	return "UP", "", exec.StoppedAt
}

