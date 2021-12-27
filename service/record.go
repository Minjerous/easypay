package service

import (
	"easypay/dao"
	"easypay/model"
)

func AddRecord(record model.Record) error {
	err := dao.InsertRecord(record)
	return err
}
func GetRecord(pid int) ([]model.Record, error) {
	record, err := dao.GetRecordByPid(pid)
	return record, err
}
