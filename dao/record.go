package dao

import "easypay/model"

func InsertRecord(record model.Record) error {
	sqlStr := "INSERT INTO record(pid, txt, recordtime) values( ?, ?, ?);"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.Pid, record.Txt, record.Time)
	return err
}
func GetRecordByPid(Pid int) ([]model.Record, error) {

	var Record []model.Record

	sqlStr := "SELECT   txt,recordtime FROM record where pid=?"
	Stmt, err := DB.Prepare(sqlStr)
	rows, err := Stmt.Query(Pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var record model.Record

		err = rows.Scan(&record.Txt, &record.Time)
		if err != nil {
			return nil, err
		}

		Record = append(Record, record)
	}

	return Record, nil
}
