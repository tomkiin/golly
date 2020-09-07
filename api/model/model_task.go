package model

import "golly/errors"

type Task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Seed    string `json:"seed"`
	BaseURL string `json:"base_url"`
	Status  string `json:"status" gorm:"default:就绪"` // 就绪、运行、结束
	Rules   []Rule `json:"rules" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Task) TableName() string {
	return "task"
}

func (t *Task) One() error {
	return crudOne(t)
}

func (t *Task) All(p *Pagination) (*[]Task, error) {
	list := &[]Task{}
	if err := crudAll(t, list, p); err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Task) Add() error {
	return crudAdd(t)
}

func (t *Task) Update() error {
	taskID := t.ID
	var rule Rule
	if err := rule.DeleteBy(taskID); err != nil {
		return err
	}
	if err := rule.AddBy(t); err != nil {
		return err
	}
	return crudUpdate(t)
}

func (t *Task) Delete() error {
	if err := crudOne(t); err != nil {
		return err
	}
	if t.Status == "运行" {
		return errors.TaskRunError
	}
	return crudDelete(t)
}
