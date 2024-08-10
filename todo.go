package todo

import (
	"encoding/json"
	"errors" // Add missing import
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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

func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task#"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++

		task := blue(item.Task)
		done := blue("no")

		createdAtStr := item.CreatedAt.Format(time.RFC822)
		completedAtStr := item.CompletedAt.Format(time.RFC822)

		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", task))
			done = green("yes")

			createdAtStr = gray(createdAtStr)

			completedAtStr = gray(completedAtStr)
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: createdAtStr},
			{Text: completedAtStr},
		})
	}
	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (t *Todos) CountPending() int {

	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}
	return total
}
