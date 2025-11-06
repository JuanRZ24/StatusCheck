package models

type Automation struct {
	gorm.Model
	Name          string `json:"name"`
	WorkflowID    string `json:"workflow_id"`
	LastRunStatus string `json:"last_run_status"`
	LastRunAt     time.Time
	ErrorMessage  string
	Enabled       bool
}
