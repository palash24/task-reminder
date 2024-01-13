package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	DueDateTime time.Time `json:"due_date_time"`
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = PgCreateTask(&newTask)
	if err != nil {
		http.Error(w, "Error inserting new task into the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task created successfully"))
}

func GetAllTheTasks(w http.ResponseWriter, r *http.Request) {
	allTasks, err := PgGetAllTasks()
	if err != nil {
		http.Error(w, "Error while retrieving all tasks from the database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allTasks)
	if err != nil {
		http.Error(w, "Error while Encoding the tasks", http.StatusInternalServerError)
		return
	}
}

func GetIndividualTasks(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
	if err != nil {
		http.Error(w, "Task ID isn't valid", http.StatusBadRequest)
		return
	}

	singleTask, err := PgGetTaskByID(taskID)
	if err != nil {
		http.Error(w, "Error retrieving task from the database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(singleTask)
	if err != nil {
		http.Error(w, "Error while Encoding the single task", http.StatusInternalServerError)
		return
	}
}

func UpdateTaskByTaskID(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
	if err != nil {
		http.Error(w, "Task ID isn't valid", http.StatusBadRequest)
		return
	}

	var singleTask Task
	err = json.NewDecoder(r.Body).Decode(&singleTask)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = PgUpdateTaskByID(taskID, &singleTask)
	if err != nil {
		http.Error(w, "Error updating task in the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task updated successfully."))
}

func DeleteTaskByTaskID(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
	if err != nil {
		http.Error(w, "Task ID isn't valid", http.StatusBadRequest)
		return
	}

	err = PgDeleteTaskByID(taskID)
	if err != nil {
		http.Error(w, "Error while deleting a task from the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully."))
}

func PgCreateTask(task *Task) error {
	_, err := DB.Model(task).Insert()
	return err
}

func PgGetAllTasks() ([]Task, error) {
	var tasks []Task
	err := DB.Model(&tasks).Select()
	return tasks, err
}

func PgGetTaskByID(taskID int) (Task, error) {
	var task Task
	err := DB.Model(&task).Where("id = ?", taskID).Select()
	return task, err
}

func PgUpdateTaskByID(taskID int, updatedTask *Task) error {
	_, err := DB.Model(updatedTask).Where("id = ?", taskID).Update()
	return err
}

func PgDeleteTaskByID(taskID int) error {
	_, err := DB.Model(&Task{}).Where("id = ?", taskID).Delete()
	return err
}