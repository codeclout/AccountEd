package ports

import (
	"github.com/google/uuid"
)

type OrganizationEvent interface {
	CommandStartEvent | CommandSucceededEvent | CommandFailedEvent
}

type OrganizationEventResult interface {
	Failure(msg string) string
	Reply(msg string) string
}

type CommandStartEvent struct {
	Command   string
	RequestID uuid.UUID
	ServiceID uuid.Domain
}

type CommandSucceededEvent struct {
	Command   string
	RequestID uuid.UUID
	ServiceID uuid.Domain
}

type CommandFailedEvent struct {
	Command   string
	RequestID uuid.UUID
	ServiceID uuid.Domain
}
