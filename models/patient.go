package models

import "github.com/gocql/gocql"

type Patient struct {
	ID           gocql.UUID `json:"id"`
	Name         string     `json:"name"`
	Age          int        `json:"age"`
	CurrentState string     `json:"current_state"`
	Medications  []string   `json:"medications"`
	CreatedAt    int64      `json:"created_at"`
}
