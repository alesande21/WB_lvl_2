package repository

import (
	"calendarEvent/internal/database"
	"calendarEvent/internal/entity"
	"calendarEvent/internal/utils"
	"context"
	"fmt"
	"log"
	"time"
)

type EventRepo struct {
	dbRepo database.DBRepository
}

func (e *EventRepo) Create(ctx context.Context, event *entity.Event) (*entity.Event, error) {
	query := `
		INSERT INTO events (id, user_id, title, date_time)
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

func (e *EventRepo) FindById(ctx context.Context, id string) (*entity.Event, error) {
	query := `
		SELECT id, user_id, title, date_time
		FROM events
		WHERE id = $1
	`

	var event entity.Event
	row := e.dbRepo.QueryRow(ctx, query, id)
	err := row.Scan(&event.ID, &event.UserID, &event.Title, &event.Date)
	if err != nil {
		log.Printf("Ошибка выполнения запроса в FindById: %v\n", err)
		return nil, err
	}

	return &event, nil
}

func (e *EventRepo) Update(ctx context.Context, event *entity.Event) (*entity.Event, error) {
	query := `
		UPDATE events
		SET title = $2, date_time = $3
		WHERE id = $1
		RETURNING id, user_id, title, date_time
	`
	row := e.dbRepo.QueryRow(ctx, query, event.ID, event.Title, event.Date)
	err := row.Scan(&event.ID, &event.UserID, &event.Title, &event.Date)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (e *EventRepo) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM events
		WHERE id = $1
	`

	_, err := e.dbRepo.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventRepo) GetEventsByWeek(ctx context.Context, date time.Time) ([]entity.Event, error) {
	query := `
		SELECT id, user_id, title, date_time
		FROM events
		WHERE DATE(event_date) >= DATE_TRUNC('week', $1::DATE)
  		AND DATE(event_date) < DATE_TRUNC('week', $2::DATE) + INTERVAL '7 days';
	`

	rows, err := e.dbRepo.Query(ctx, query, date, date)

	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса выборки заказов: %v", err)
	}
	defer rows.Close()

	var events []entity.Event

	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.ID, &event.UserID, &event.Title, &event.Date)
		if err != nil {
			return nil, fmt.Errorf("ошибка в сканировании данных заказов: %v", err)
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка после итерации по строкам: %v", err)
	}

	return events, nil
}

func (e *EventRepo) GetEventsByDay(ctx context.Context, date time.Time) ([]entity.Event, error) {
	query := `
		SELECT id, user_id, title, date_time
		FROM events
		WHERE DATE(event_date) = $1
	`

	rows, err := e.dbRepo.Query(ctx, query, date)

	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса выборки заказов: %v", err)
	}
	defer rows.Close()

	var events []entity.Event

	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.ID, &event.UserID, &event.Title, &event.Date)
		if err != nil {
			return nil, fmt.Errorf("ошибка в сканировании данных заказов: %v", err)
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка после итерации по строкам: %v", err)
	}

	return events, nil
}

func (e *EventRepo) GetEventsByMonth(ctx context.Context, date time.Time) ([]entity.Event, error) {
	query := `
		SELECT id, user_id, title, date_time
		FROM events
		WHERE EXTRACT(YEAR FROM event_date) = $1
  		AND EXTRACT(MONTH FROM event_date) = $2;
	`

	rows, err := e.dbRepo.Query(ctx, query, date.Year(), date.Month())

	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса выборки заказов: %v", err)
	}
	defer rows.Close()

	var events []entity.Event

	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.ID, &event.UserID, &event.Title, &event.Date)
		if err != nil {
			return nil, fmt.Errorf("ошибка в сканировании данных заказов: %v", err)
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка после итерации по строкам: %v", err)
	}

	return events, nil
}

func (e *EventRepo) Ping() error {
	return nil
}

func NewEventRepo(dbRepo database.DBRepository) *EventRepo {
	return &EventRepo{dbRepo: dbRepo}
}
