package repository

import (
	"calendarEvent/internal/database"
	"context"
)

type EventRepo struct {
	dbRepo database.DBRepository
}

func (e EventRepo) CreateOrder(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (e EventRepo) Ping() error {
	//TODO implement me
	panic("implement me")
}

func NewEventRepo(dbRepo database.DBRepository) *EventRepo {
	return &EventRepo{dbRepo: dbRepo}
}
