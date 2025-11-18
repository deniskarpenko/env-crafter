package project

import "time"

type Project struct {
	id          string
	name        string
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
	services    []Service
	events      []DomainEvent
}
