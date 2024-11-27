package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Event структура для события
type Event struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// Инициализация базы данных и других сервисов
// Используется временная заглушка
var events []Event

// Сериализация ошибок в JSON
func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Сериализация успешного ответа в JSON
func respondWithResult(w http.ResponseWriter, result string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

// Middleware для логирования запросов
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Обработчик для создания события
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	// Чтение данных из запроса
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверка входных данных
	if event.UserID == 0 || event.Title == "" || event.StartTime.IsZero() || event.EndTime.IsZero() {
		respondWithError(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Добавление события (в реальном проекте это должно быть записано в базу данных)
	events = append(events, event)

	respondWithResult(w, "Event created successfully")
}

// Обработчик для обновления события
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	var updatedEvent Event
	// Чтение данных из запроса
	err := json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Поиск и обновление события
	for i, event := range events {
		if event.ID == updatedEvent.ID {
			events[i] = updatedEvent
			respondWithResult(w, "Event updated successfully")
			return
		}
	}

	respondWithError(w, "Event not found", http.StatusNotFound)
}

// Обработчик для удаления события
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var eventToDelete Event
	// Чтение данных из запроса
	err := json.NewDecoder(r.Body).Decode(&eventToDelete)
	if err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Удаление события
	for i, event := range events {
		if event.ID == eventToDelete.ID {
			events = append(events[:i], events[i+1:]...)
			respondWithResult(w, "Event deleted successfully")
			return
		}
	}

	respondWithError(w, "Event not found", http.StatusNotFound)
}

// Обработчик для получения событий по дням
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь необходимо добавить логику получения событий для дня
	// Для примера, возвращаем все события
	respondWithResult(w, fmt.Sprintf("Events for the day: %v", events))
}

// Обработчик для получения событий по неделям
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь необходимо добавить логику получения событий для недели
	// Для примера, возвращаем все события
	respondWithResult(w, fmt.Sprintf("Events for the week: %v", events))
}

// Обработчик для получения событий по месяцам
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь необходимо добавить логику получения событий для месяца
	// Для примера, возвращаем все события
	respondWithResult(w, fmt.Sprintf("Events for the month: %v", events))
}
