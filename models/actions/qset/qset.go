package qset

import (
	"context"
	"database/sql"
	"log"
	question "questionService/libs"
	"questionService/models"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Insert(db *sql.DB, s *question.Set) (int64, error) {
	query := "INSERT INTO qset(time, type_id, created_by_id, created_date, updated_date, active, qs_name) VALUES (?, ?, ?, ?, ?, ?, ?)"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.Time, s.Type, s.CreatedById, s.CreatedDate, s.UpdatedDate, s.Active, s.QsName)
	if err != nil {
		log.Printf("Error %s when inserting row into table", err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error %s when getting last inserted product", err)
		return 0, err
	}
	log.Printf("Object with ID %d created", id)
	return id, nil
}

func Update(db *sql.DB, s *question.Set, oid uint32) error {
	query := "UPDATE qset SET time=?, type_id=?, created_by_id=?, created_date=?, updated_date=?, active=?, qs_name=? WHERE id=?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return status.Errorf(codes.Internal, "Error updating: %v", err)
	}
	res, err := stmt.Exec(s.Time, s.Type, s.CreatedById, s.CreatedDate, s.UpdatedDate, s.Active, s.QsName, oid)
	if err != nil {
		return status.Errorf(codes.Internal, "Error saving: %v", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return status.Errorf(codes.Internal, "Error getting rows affected: %v", err)
	}
	log.Printf("Rows affected: %v", rows)
	return nil
}

func List(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT time, type_id, created_by_id, created_date, updated_date, active, qs_name FROM qset"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Retrieve(db *sql.DB, oid uint32) (*question.Set, error) {
	query := "SELECT time, type_id, created_by_id, created_date, updated_date, active, qs_name FROM qset WHERE id=?"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row := models.Dbclient.QueryRow(query, oid)

	var qsetItem = question.Set{}
	er := row.Scan(&qsetItem.Time, &qsetItem.Type, &qsetItem.CreatedById, &qsetItem.CreatedDate, &qsetItem.UpdatedDate, &qsetItem.Active, &qsetItem.QsName)
	if er != nil {
		log.Printf("Error when getting the object %v", er.Error())
		return nil, er
	}
	return &qsetItem, nil
}

func Delete(db *sql.DB, oid uint32) error {
	query := "DELETE FROM qset WHERE id=?"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Exec(query, oid)
	if err != nil {
		return err
	}
	return err
}
