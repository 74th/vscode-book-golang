package tasks

import "fmt"

// Repository タスクリポジトリ
type Repository struct {
	tasks []Task
}

// NewRepository 新しいタスクリポジトリを作成する
func NewRepository() *Repository {
	r := new(Repository)
	r.tasks = make([]Task, 2, 20)
	r.tasks[0] = Task{
		ID:   1,
		Text: "task1",
		Done: false,
	}
	r.tasks[1] = Task{
		ID:   2,
		Text: "task2",
		Done: false,
	}
	return r
}

// Add タスクを追加する
func (r *Repository) Add(task Task) int {
	task.ID = len(r.tasks) + 1
	r.tasks = append(r.tasks, task)
	return task.ID
}

// List タスクの一覧を取得する
func (r *Repository) List() []*Task {
	result := []*Task{}
	for i, task := range r.tasks {
		if !task.Done {
			// taskは一時変数のインスタンスのため、
			// リポジトリのtasksのインスタンスを返す
			result = append(result, &r.tasks[i])
		}
	}
	return result
}

// Done タスクを完了させる
func (r *Repository) Done(id int) error {
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("Not found id:%d", id)
}