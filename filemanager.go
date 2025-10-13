package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func LoadTasks(filePath string) ([]*Task, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("failed to open the file")
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, errors.New("failed to read the file")
	}
	if len(data) == 0 {
		return []*Task{}, nil
	}
	var tasks []*Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(filePath string, tasks []*Task) error {
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return errors.New("failed to convert tasks to JSON")
	}

	tmp := filePath + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return errors.New("failed to save data to file")
	}
	return os.Rename(tmp, filePath)
}
