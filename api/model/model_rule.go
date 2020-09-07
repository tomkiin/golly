package model

type Rule struct {
	ID      int    `json:"id"`
	TaskID  int    `json:"task_id"`
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
	Runtime int    `json:"runtime"`
	Find    string `json:"find"` // attr, text
	Attr    string `json:"attr,omitempty"`
}

func (Rule) TableName() string {
	return "rule"
}

func (r *Rule) AddBy(task *Task) error {
	for _, rule := range task.Rules {
		if err := crudAdd(rule); err != nil {
			return err
		}
	}
	return nil
}

func (r *Rule) DeleteBy(taskID int) error {
	where := Where{"task_id": taskID}
	return crudDeleteBy(r, where)
}
