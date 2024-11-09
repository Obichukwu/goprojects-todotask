package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type TaskManager []Task

func (taskMgr *TaskManager) Add(title string, description string) {
	lastId := len(*taskMgr)
	task := Task{
		Id:          lastId + 1,
		Title:       title,
		Description: description,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
	*taskMgr = append(*taskMgr, task)
}

func (taskMgr *TaskManager) validateIndex(index int) error {
	if index < 0 || index >= len(*taskMgr) {
		return errors.New("invalid index")
	}
	return nil
}

func (taskMgr *TaskManager) Get(index int) (*Task, error) {
	// Validate index to avoid out-of-bounds errors
	if err := taskMgr.validateIndex(index); err != nil {
		return nil, err
	}

	// Return the task at the specified index
	return &(*taskMgr)[index], nil
}

func (taskMgr *TaskManager) Complete(index int) error {
	completedTask, err := taskMgr.Get(index)
	if err != nil {
		return err
	}

	// Set the CompletedAt field to the current time
	completedAt := time.Now()
	completedTask.CompletedAt = &completedAt

	// Optionally, you can also set the UpdatedAt field to the current time
	updatedAt := time.Now()
	completedTask.UpdatedAt = &updatedAt

	return nil
}

func (taskMgr *TaskManager) Delete(index int) error {
	// Validate index directly on taskMgr to ensure it's within bounds
	if err := taskMgr.validateIndex(index); err != nil {
		return err
	}

	// Perform deletion directly on the original slice
	*taskMgr = append((*taskMgr)[:index], (*taskMgr)[index+1:]...)
	return nil
}

// Load taskMgr from a JSON file
func (taskMgr *TaskManager) loadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v", filePath, err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v", filePath, err)
	}

	err = json.Unmarshal(bytes, taskMgr)
	if err != nil {
		return fmt.Errorf("could not unmarshal JSON: %v", err)
	}
	fmt.Printf("Loaded %d tasks\n", len(*taskMgr))
	return nil
}

func (taskMgr *TaskManager) saveTaskManagerToFile(filePath string) error {
	// Marshal the TaskManager into JSON format
	data, err := json.MarshalIndent(taskMgr, "", "    ")
	if err != nil {
		return fmt.Errorf("could not marshal task manager: %v", err)
	}

	// Write the marshaled JSON data to the specified file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %v", filePath, err)
	}

	return nil
}
