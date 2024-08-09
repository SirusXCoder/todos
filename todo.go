package todo

import (
	"encoding/json"
	"errors" // Add missing import
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		// Corrected: errors.New(text:"invalid index") -> errors.New("invalid index")
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		// Corrected: errors.New(text:"invalid index") -> errors.New("invalid index")
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

func (t *Todos) Load(filename string) error {
	// Corrected typo: filenae -> filename
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		// Corrected: if len(file) == 0 [ -> if len(file) == 0 {
		return nil
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	// Corrected typo: filenae -> filename
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	// Corrected: ioutil.WriteFile(filenae, data, perm:0644 ) -> ioutil.WriteFile(filename, data, 0644)
	return os.WriteFile(filename, data, 0644)
}
