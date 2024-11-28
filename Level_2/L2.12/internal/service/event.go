package service

import (
	"calendarEvent/internal/entity"
	"context"
	"log"
	"time"
)

type EventRepo interface {
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) (*entity.Event, error)
	FindById(ctx context.Context, id string) (*entity.Event, error)
	Delete(ctx context.Context, id string) error
	GetEventsByWeek(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetEventsByDay(ctx context.Context, date time.Time) ([]entity.Event, error)
	GetEventsByMonth(ctx context.Context, date time.Time) ([]entity.Event, error)
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
		log.Printf("Ошибка выполнения запроса в FindById: %v\n", err)
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

func (s *EventService) DeleteEvent(ctx context.Context, id string, idUser string) error {
	foundedEvent, err := s.Repo.FindById(ctx, id)
	if err != nil {
		log.Printf("Ошибка выполнения запроса в DeleteEvent: %v\n", err)
		return err
	}

	if foundedEvent.UserID != idUser {
		log.Printf("Нет прав доступа: %v\n", err)
		return err
	}

	err = s.Repo.Delete(ctx, id)
	if err != nil {
		log.Printf("Не удалось удалить: %v\n", err)
		return err
	}

	return nil
}
