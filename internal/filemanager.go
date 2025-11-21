package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadTasks(filePath string) ([]*Task, error) {
	file, err := os.Open(filePath)

	if err != nil {
		if os.IsNotExist(err) {

			empty := []*Task{}
			data, _ := json.Marshal(empty)

			if err := os.WriteFile(filePath, data, 0644); err != nil {
				return nil, fmt.Errorf("failed to create file: %w", err)
			}
			return empty, nil
		}
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("failed to read the file")
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
		return fmt.Errorf("failed to convert tasks to JSON")
	}

	tmp := filePath + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return fmt.Errorf("failed to save data to file")
	}
	return os.Rename(tmp, filePath)
}
