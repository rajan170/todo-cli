package todo

import (
	"encoding/json"
	"errors"
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

/*
ls := *t: This line creates a local variable ls
and assigns the dereferenced value of the receiver t
to it. Similar to the previous example, it's done to
make the code more concise and readable by using ls
instead of *t throughout the method.
*/
func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

/*
*t = append(ls[:index-1], ls[index:]...): If the index is valid,
this line updates the todo list by removing the todo item at the
specified index. It uses the append function to concatenate two
slices: the elements before the specified index (ls[:index-1]) and
the elements after the specified index (ls[index:]...).
*/
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644) //0644 is the unix-style permission code.
}
