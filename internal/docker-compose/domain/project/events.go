package project

import "time"

type DomainEvent interface {
	EventName() string
	OccurredAt() time.Time
	AggregateID() string
}

type baseEvent struct {
	eventName   string
	occurredAt  time.Time
	aggregateID string
}

func (e baseEvent) EventName() string     { return e.eventName }
func (e baseEvent) OccurredAt() time.Time { return e.occurredAt }
func (e baseEvent) AggregateID() string   { return e.aggregateID }

type Created struct {
	baseEvent
	ProjectName string
	Description string
}
