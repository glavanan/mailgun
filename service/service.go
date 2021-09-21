package service

import (
	"errors"

	"github.com/glavanan/mailgun/model"
	"gorm.io/gorm"
)

//Delivered update domain in DB to add an event in deliver
func Delivered(domainName string, db *gorm.DB) error {
	dom := model.Domain{Name: domainName}
	err := db.Model(&dom).UpdateColumn("deliver", gorm.Expr("deliver + ?", 1)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		dom.Deliver = 1
		err := db.Create(&dom).Error
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

//Bounced update bounce to add an event
func Bounced(domainName string, db *gorm.DB) error {
	dom := model.Domain{Name: domainName}
	err := db.Model(&dom).UpdateColumn("bounce", gorm.Expr("bounce + ?", 1)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		dom.Bounce = 1
		err := db.Create(&dom).Error
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

//GetCathAllStatus check status catch-all of a domain
func GetCathAllStatus(domainName string, db *gorm.DB) (string, error) {
	dom := model.Domain{Name: domainName}
	err := db.First(&dom).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "unknow", nil
	}
	if err != nil {
		return "", err
	}

	if dom.Bounce >= 1 || dom.Deliver < 1000 {
		return "not catch-all", nil
	}
	return "catch-all", nil
}
