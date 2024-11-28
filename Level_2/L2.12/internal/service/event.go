package service

import (
	"calendarEvent/internal/entity"
	"context"
	"log"
)

type EventRepo interface {
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) (*entity.Event, error)
	FindById(ctx context.Context, id string) (*entity.Event, error)
	Ping() error
}

type EventService struct {
	Repo EventRepo
}

func NewOrderService(repo EventRepo) *EventService {
	return &EventService{Repo: repo}
}

//func (s *EventService) CreateEvent(ctx context.Context, event *entity.Event) error {
//	return nil
//}

func (s *EventService) UpdateEvent(ctx context.Context, newEvent *entity.Event) (*entity.Event, error) {
	foundedEvent, err := s.Repo.FindById(ctx, newEvent.ID)
	if err != nil {
		log.Printf("Ошибка выполнения запроса в UpdateEvent: %v\n", err)
		return nil, err
	}

	if foundedEvent.UserID != newEvent.UserID {
		log.Printf("Нет прав доступа: %v\n", err)
		return nil, err
	}

	newEvent, err = s.Repo.Update(ctx, newEvent)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
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
