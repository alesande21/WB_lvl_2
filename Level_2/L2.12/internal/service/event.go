package service

import (
	"calendarEvent/internal/entity"
	"context"
)

type EventRepo interface {
	CreateEvent(ctx context.Context, event *entity.Event) error
	Ping() error
}

type EventService struct {
	Repo EventRepo
}

func NewOrderService(repo EventRepo) *EventService {
	return &EventService{Repo: repo}
}

func (s *EventService) CreateEvent(title, date string) error {
	return nil
}

func (s *EventService) UpdateEvent(id, title, date string) error {
	return nil
}

func (s *EventService) DeleteEvent(id string) error {
	return nil
}

func (s *EventService) GetEventsForDay(date string) ([]string, error) {
	return []string{"Event1"}, nil
}

func (s *EventService) GetEventsForWeek(date string) ([]string, error) {
	return []string{"Event1"}, nil
}

func (s *EventService) GetEventsForMonth(date string) ([]string, error) {
	return []string{"Event1"}, nil
}
