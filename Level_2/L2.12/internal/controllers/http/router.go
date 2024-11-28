package http

import (
	"calendarEvent/internal/entity"
	"calendarEvent/internal/service"
	"encoding/json"
	"fmt"
	log2 "github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

type jsonEvent struct {
	ID     string    `json:"id"`
	UserID string    `json:"user_id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
}

type ErrorResponse struct {
	Reason string `json:"reason"`
}

type RequestUpdateJSONEvent struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

func NewRouter(orderService *service.EventService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", logMiddleware(createEventHandler(orderService)))
	mux.HandleFunc("/update_event", logMiddleware(updateEventHandler(orderService)))
	mux.HandleFunc("/delete_event", logMiddleware(deleteEventHandler(orderService)))
	mux.HandleFunc("/events_for_day", logMiddleware(eventsForDayHandler(orderService)))
	mux.HandleFunc("/events_for_week", logMiddleware(eventsForWeekHandler(orderService)))
	mux.HandleFunc("/events_for_month", logMiddleware(eventsForMonthHandler(orderService)))
	return mux
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Запрос: %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Обработан за %s", time.Since(start))
	}
}

func sendErrorResponse(w http.ResponseWriter, code int, resp ErrorResponse) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

// создание ивента
func createEventHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			sendErrorResponse(w, http.StatusMethodNotAllowed, ErrorResponse{Reason: "Только POST методы поддерживаются."})
			return
		}

		event, status, err := ParseAndValidationEvent(r)

		if err != nil {
			log.Println("Неверный формат для ивента!")
			sendErrorResponse(w, status, ErrorResponse{Reason: "Неверный формат для предложения."})
			return
		}

		err = service.Repo.Create(r.Context(), event)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, ErrorResponse{Reason: "Не удалось создать ивент."})
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		msg := "Ивент принят в обработку."
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			log2.Errorf("CreateOrder-> json.NewEncoder: ошибка при кодирования овета: %s", err.Error())
			sendErrorResponse(w, http.StatusInternalServerError, ErrorResponse{Reason: "Ошибка кодирования ответа."})
		}
	}
}

// обновление события
func updateEventHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			sendErrorResponse(w, http.StatusMethodNotAllowed, ErrorResponse{Reason: "Только POST методы поддерживаются."})
			return
		}

		event, status, err := ParseAndValidationEvent(r)
		if err != nil {
			log.Println("Неверный формат для извения ивента!")
			sendErrorResponse(w, status, ErrorResponse{Reason: "Неверный формат для предложения."})
			return
		}

		event, err = service.UpdateEvent(r.Context(), event)
		if err != nil {
			sendErrorResponse(w, http.StatusServiceUnavailable, ErrorResponse{Reason: "Не удалось обновить."})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		msg := "Ивент принят в обработку."
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			log2.Errorf("CreateOrder-> json.NewEncoder: ошибка при кодирования овета: %s", err.Error())
			sendErrorResponse(w, http.StatusInternalServerError, ErrorResponse{Reason: "Ошибка кодирования ответа."})
		}
	}
}

// удаление события
func deleteEventHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			sendErrorResponse(w, http.StatusMethodNotAllowed, ErrorResponse{Reason: "Только POST методы поддерживаются."})
			return
		}

		err := r.ParseForm()
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Неверные параметры запроса."})
			return
		}

		reqId := r.FormValue("id")
		reqIdUser := r.FormValue("id_user")

		if reqId == "" || reqIdUser == "" {
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Не переданы обязательные параметры."})
			return
		}

		err = service.DeleteEvent(r.Context(), reqId, reqIdUser)
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Не удалось удалить."})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		msg := "Ивент удален."
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			log2.Errorf("CreateOrder-> json.NewEncoder: ошибка при кодирования овета: %s", err.Error())
			sendErrorResponse(w, http.StatusInternalServerError, ErrorResponse{Reason: "Ошибка кодирования ответа."})
		}
	}
}

// ивенты за день
func eventsForDayHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date_time")
		if date == "" {
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Не передана дата."})
			return
		}

		layout := "2006-01-02"
		parsedTime, err := time.Parse(layout, date)
		if err != nil {
			log.Println("Ошибка при парсинге времени!")
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Формат даты неверный."})
			return
		}

		events, err := service.Repo.GetEventsByDay(r.Context(), parsedTime)
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Не удалось получить список ивентов."})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(events); err != nil {
			log2.Errorf("CreateOrder-> json.NewEncoder: ошибка при кодирования овета: %s", err.Error())
			sendErrorResponse(w, http.StatusInternalServerError, ErrorResponse{Reason: "Ошибка кодирования ответа."})
		}
	}
}

// за неделю
func eventsForWeekHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date_time")
		if date == "" {
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Не передана дата."})
			return
		}

		layout := "2006-01-02"
		parsedTime, err := time.Parse(layout, date)
		if err != nil {
			log.Println("Ошибка при парсинге времени!")
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Формат даты неверный."})
			return
		}

		events, err := service.Repo.GetEventsByWeek(r.Context(), parsedTime)
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Не удалось получить список ивентов."})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(events); err != nil {
			log2.Errorf("CreateOrder-> json.NewEncoder: ошибка при кодирования овета: %s", err.Error())
			sendErrorResponse(w, http.StatusInternalServerError, ErrorResponse{Reason: "Ошибка кодирования ответа."})
		}
	}
}

// за месяц
func eventsForMonthHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		if date == "" {
			http.Error(w, `{"error":"не передана дата"}`, http.StatusBadRequest)
			return
		}

		layout := "2006-01-02"
		parsedTime, err := time.Parse(layout, date)
		if err != nil {
			log.Println("Ошибка при парсинге времени!")
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Формат даты неверный."})
			return
		}
		events, err := service.Repo.GetEventsByMonth(r.Context(), parsedTime)
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, ErrorResponse{Reason: "Не удалось получить список ивентов."})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(events); err != nil {
			log2.Errorf("CreateOrder-> json.NewEncoder: ошибка при кодирования овета: %s", err.Error())
			sendErrorResponse(w, http.StatusInternalServerError, ErrorResponse{Reason: "Ошибка кодирования ответа."})
		}
	}
}

func ParseAndValidationEvent(r *http.Request) (*entity.Event, int, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	title := r.FormValue("title")
	dateStr := r.FormValue("date")
	userID := r.FormValue("user_id")
	eventID := r.FormValue("id")

	if title == "" || dateStr == "" || userID == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("пропущены поля")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("неверный формат даты")
	}

	event := &entity.Event{
		ID:     eventID,
		UserID: userID,
		Title:  title,
		Date:   date,
	}
	return event, http.StatusOK, nil
}
