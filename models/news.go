package models

import (
	"time"
)

type News struct {
	Title         string    `json: "title"`
	Link          string    `json: "link"`
	Description   string    `josn: "description"`
	UpdatedTime   time.Time `json: "updated_time"`
	PassedHour    int       `json: "passed_hour`
	PassedMinutes int       `json: "passed_minutes`
	PassedDay     int       `json: "passed_day`
}
