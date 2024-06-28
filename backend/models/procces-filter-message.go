package models

type ProcessFilterMsg struct {
	Name       string `json:"name"`
	Monitoring bool   `json:"monitoring"`
}
