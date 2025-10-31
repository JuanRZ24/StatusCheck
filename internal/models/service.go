package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name              string `json:"name"`
	URL               string `json:"url"`
	Method            string `json:"method" gorm:"default:GET"`
	ExpectedStatus    int    `json:"expected_status" gorm:"default:200"`
	ExpectedSubstring string `json:"expected_substring"`
	TimeoutMS         int    `json:"timeout_ms" gorm:"default:5000"`
	IntervalSec       int    `json:"interval_sec" gorm:"default:120"`
	Retries           int    `json:"retries" gorm:"default:2"`
	Enabled           bool   `json:"enabled" gorm:"default:true"`
	Status            string `json:"status" gorm:"default:UNKNOWN"`
	LatencyMS         int    `json:"latency_ms"`
	Error             string `json:"error"`
}
