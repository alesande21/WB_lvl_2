package repository

import (
	"calendarEvent/internal/database"
	"calendarEvent/internal/entity"
	"calendarEvent/internal/utils"
	"context"
	"fmt"
	"log"
)

type EventRepo struct {
	dbRepo database.DBRepository
}

func (e *EventRepo) CreateEvent(ctx context.Context, event *entity.Event) (*entity.Event, error) {
	query := `
		INSERT INTO bid (id, user_id, title, date_time)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, title, date_time
	`

	newUuid, _ := utils.GenerateUUIDV7()
	if newUuid == "" {
		log.Println("Ошибка: не удалось сгенерировать UUID")
		return nil, fmt.Errorf("не удалось сгенерировать UUID")
	}

	row := e.dbRepo.QueryRow(ctx, query, newUuid, event.UserID, event.Title, event.Date)
	err := row.Scan(&event.ID, &event.UserID, &event.Title, &event.Date)
	if err != nil {
		log.Printf("Ошибка выполнения запроса в CreateEvent: %v\n", err)
		return nil, err
	}

	return event, nil

}

func (e *EventRepo) Ping() error {
	//TODO implement me
	panic("implement me")
}

func NewEventRepo(dbRepo database.DBRepository) *EventRepo {
	return &EventRepo{dbRepo: dbRepo}
}
