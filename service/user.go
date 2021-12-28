package service

import (
	"database/sql"
	"easypay/dao"
	"easypay/model"
	"easypay/tool"
)

func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if tool.Match(user.Password, password) {
		return true, nil
	} else {
		return false, nil
	}
	return true, nil
}
func SelectIdByUsername(username string) (int, error) {

	userId, err := dao.SelectIdByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return userId, nil
		}
		return userId, err
	}

	return userId, err
}
func IncreaseMoney(user model.User) error {
	err := dao.IncreaseMoney(user)
	return err
}
func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, err
}

func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}

func IsMoneyEnough(money float64, name string) (bool, error) {
	Mymoey, err := dao.IsMoneyEnough(name)
	if money > Mymoey {
		return false, err
	}
	return true, err
}
func Password(user model.User) error {
	err := dao.UpdateUser(user)
	return err
}
func SelectMoneyByName(name string) (float64, error) {
	Money, err := dao.SelectMoneyByName(name)
	return Money, err
}
func GetRecordByName(name string, id int) ([]model.Record, error) {
	Record, err := dao.GetRecordByName(name, id)
	return Record, err
}
