package models

import "time"

// Repository represents a GitHub repository
type Repository struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	LastUpdate   time.Time `json:"updated_at"`
	Private      bool      `json:"private"`
	Organization string    `json:"organization,omitempty"`
}

type Orgs struct {
	Login string `json:"login"`
}
