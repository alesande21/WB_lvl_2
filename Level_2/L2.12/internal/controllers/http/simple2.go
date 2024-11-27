package http

import (
	"calendarEvent/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Event struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

var events []Event

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

func createEventHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"только POST методы поддерживаются"}`, http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, `{"error":"не удалось разобрать параметры"}`, http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		date := r.FormValue("date")

		if title == "" || date == "" {
			http.Error(w, `{"error":"не переданы обязательные параметры"}`, http.StatusBadRequest)
			return
		}

		// вызов бизнес-логики
		err = service.CreateEvent(title, date)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":"event created"}`))
	}
}

// updateEventHandler обрабатывает обновление события
func updateEventHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"только POST методы поддерживаются"}`, http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, `{"error":"не удалось разобрать параметры"}`, http.StatusBadRequest)
			return
		}

		id := r.FormValue("id")
		title := r.FormValue("title")
		date := r.FormValue("date")

		if id == "" || title == "" || date == "" {
			http.Error(w, `{"error":"не переданы обязательные параметры"}`, http.StatusBadRequest)
			return
		}

		err = service.UpdateEvent(id, title, date)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":"event updated"}`))
	}
}

// deleteEventHandler обрабатывает удаление события
func deleteEventHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"только POST методы поддерживаются"}`, http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, `{"error":"не удалось разобрать параметры"}`, http.StatusBadRequest)
			return
		}

		id := r.FormValue("id")

		if id == "" {
			http.Error(w, `{"error":"не передан id события"}`, http.StatusBadRequest)
			return
		}

		err = service.DeleteEvent(id)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":"event deleted"}`))
	}
}

// eventsForDayHandler обрабатывает запрос событий на день
func eventsForDayHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		if date == "" {
			http.Error(w, `{"error":"не передана дата"}`, http.StatusBadRequest)
			return
		}

		events, err := service.GetEventsForDay(date)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusServiceUnavailable)
			return
		}

		response, _ := json.Marshal(map[string]interface{}{"result": events})
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

// eventsForWeekHandler обрабатывает запрос событий на неделю
func eventsForWeekHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		if date == "" {
			http.Error(w, `{"error":"не передана дата"}`, http.StatusBadRequest)
			return
		}

		events, err := service.GetEventsForWeek(date)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusServiceUnavailable)
			return
		}

		response, _ := json.Marshal(map[string]interface{}{"result": events})
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

// eventsForMonthHandler обрабатывает запрос событий на месяц
func eventsForMonthHandler(service *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		if date == "" {
			http.Error(w, `{"error":"не передана дата"}`, http.StatusBadRequest)
			return
		}

		events, err := service.GetEventsForMonth(date)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusServiceUnavailable)
			return
		}

		response, _ := json.Marshal(map[string]interface{}{"result": events})
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
