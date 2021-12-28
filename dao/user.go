package dao

import "easypay/model"

func InsertUser(user model.User) error {
	sqlStr := "INSERT INTO user(username, password)  values(?, ?);"
	Stmt, err := DB.Prepare(sqlStr)
	Stmt.Exec(user.Username, user.Password)
	return err
}

func UpdateUser(user model.User) error {
	sqlStr := "update user set password=? where username =?"
	Stmt, err := DB.Prepare(sqlStr)
	Stmt.Exec(user.Password, user.Username)
	return err
}

func SelectIdByUsername(username string) (int, error) {
	var userId int
	sqlStr := "SELECT id  FROM user where username = ? "
	Stmt, err := DB.Prepare(sqlStr)
	row := Stmt.QueryRow(username)
	if row.Err() != nil {
		return userId, row.Err()
	}
	err = row.Scan(&userId)
	if err != nil {
		return userId, err
	}

	return userId, nil
}
func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	sqlStr := "SELECT id, password FROM user where username = ? "
	Stmt, err := DB.Prepare(sqlStr)
	row := Stmt.QueryRow(username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err = row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func IncreaseMoney(user model.User) error {
	sqlStr := "update  user  set  money=money +? where username = ?;"
	Stmt, err := DB.Prepare(sqlStr)
	_, err = Stmt.Exec(user.Money, user.Username)
	return err
}
func IsMoneyEnough(name string) (float64, error) {
	var Mymoney float64
	sqlStr := "SELECT money  FROM user where username = ? "
	Stmt, err := DB.Prepare(sqlStr)
	row := Stmt.QueryRow(name)
	if row.Err() != nil {
		return Mymoney, row.Err()
	}
	err = row.Scan(&Mymoney)
	if err != nil {
		return Mymoney, err
	}
	return Mymoney, nil
}
func SelectMoneyByName(username string) (float64, error) {
	var Money float64
	sqlStr := "SELECT money FROM user where username = ? "
	Stmt, err := DB.Prepare(sqlStr)
	row := Stmt.QueryRow(username)
	if row.Err() != nil {
		return Money, row.Err()
	}
	err = row.Scan(&Money)
	if err != nil {
		return Money, err
	}

	return Money, nil
}
