package internal

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
)

type Reminder struct {
	ID          int       `json:"id"`
	TaskID      int       `json:"task_id"`
	Description string    `json:"description"`
	DueDateTime time.Time `json:"due_date_time"`
}

var DB *pg.DB

func CreateNewReminder(w http.ResponseWriter, r *http.Request) {
	var reminder Reminder
	err := json.NewDecoder(r.Body).Decode(&reminder)
	if err != nil {
		http.Error(w, "Invalid reminder data", http.StatusBadRequest)
		return
	}

	err = PgCreateReminder(&reminder)
	if err != nil {
		http.Error(w, "Failed to create reminder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Reminder created successfully"))
}

func GetAllReminders(w http.ResponseWriter, r *http.Request) {
	reminders, err := PgGetAllReminders()
	if err != nil {
		http.Error(w, "Failed to retrieve reminders", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(reminders)
	if err != nil {
		http.Error(w, "Error while Encoding all the reminders", http.StatusInternalServerError)
		return
	}
}

func GetSingleReminderByID(w http.ResponseWriter, r *http.Request) {
	reminderID := chi.URLParam(r, "reminderID")

	reminder, err := PgGetReminderByID(reminderID)
	if err != nil {
		http.Error(w, "Failed to retrieve reminder", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(reminder)
	if err != nil {
		http.Error(w, "Error while Encoding the single reminder", http.StatusInternalServerError)
		return
	}
}

func UpdateReminderByID(w http.ResponseWriter, r *http.Request) {
	reminderID := chi.URLParam(r, "reminderID")

	var updatedReminder Reminder
	err := json.NewDecoder(r.Body).Decode(&updatedReminder)
	if err != nil {
		http.Error(w, "Invalid reminder data", http.StatusBadRequest)
		return
	}

	err = PgUpdateReminderByID(reminderID, &updatedReminder)
	if err != nil {
		http.Error(w, "Failed to update reminder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reminder updated successfully."))
}

func DeleteReminderByID(w http.ResponseWriter, r *http.Request) {
	reminderID := chi.URLParam(r, "reminderID")

	err := PgDeleteReminderByID(reminderID)
	if err != nil {
		http.Error(w, "Failed to delete reminder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reminder deleted successfully."))
}

func PgGetAllReminders() ([]Reminder, error) {
	var reminders []Reminder
	err := DB.Model(&reminders).Select()
	return reminders, err
}

func PgGetReminderByID(reminderID string) (Reminder, error) {
	var reminder Reminder
	err := DB.Model(&reminder).Where("id = ?", reminderID).Select()
	return reminder, err
}

func PgCreateReminder(reminder *Reminder) error {
	_, err := DB.Model(reminder).Insert()
	return err
}

func PgUpdateReminderByID(reminderID string, updatedReminder *Reminder) error {
	_, err := DB.Model(updatedReminder).Where("id = ?", reminderID).Update()
	return err
}

func PgDeleteReminderByID(reminderID string) error {
	_, err := DB.Model(&Reminder{}).Where("id = ?", reminderID).Delete()
	return err
}