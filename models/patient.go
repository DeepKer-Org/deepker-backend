package models

import "github.com/gocql/gocql"

type Patient struct {
	ID           gocql.UUID `json:"id"`
	Name         string     `json:"name"`
	Age          int        `json:"age"`
	CurrentState string     `json:"current_state"`
	Medications  []string   `json:"medications"`
	Auditable    Auditable  `json:"auditable"`
}
