package model

import (
	"fmt"
	"gorm.io/gorm/clause"
)

type Where map[string]interface{}
type Pagination struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

func crudOne(m interface{}) error {
	return DB.Preload(clause.Associations).First(m).Error
}

func crudAll(m interface{}, list interface{}, p *Pagination) error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Size <= 0 {
		p.Size = 10
	}
	offset := (p.Page - 1) * p.Size
	limit := p.Size
	return DB.Model(m).Preload(clause.Associations).Count(&p.Total).Offset(offset).Limit(limit).Find(list).Error
}

func crudAllBy(list interface{}, where Where) error {
	db := DB.Preload(clause.Associations)
	for k, v := range where {
		db = db.Where(fmt.Sprintf("%s = ?", k), v)
	}
	return db.Find(list).Error
}

func crudAdd(m interface{}) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func crudUpdate(m interface{}) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Save(m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func crudDelete(m interface{}) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Delete(m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func crudDeleteBy(m interface{}, where Where) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	for k, v := range where {
		tx = tx.Where(fmt.Sprintf("%s = ?", k), v)
	}
	if err := tx.Delete(m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
