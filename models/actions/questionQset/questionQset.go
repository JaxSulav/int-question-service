package _questionQset

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

func Insert(db *sql.DB, s *question.QuestionSet) (int64, error) {
	query := "INSERT INTO question_qset(question_id, qset_id, created_by_id, created_date, updated_date) VALUES (?, ?, ?, ?, ?)"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.QuestionId, s.SetId, s.CreatedById, s.CreatedDate, s.UpdatedDate)
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

func Update(db *sql.DB, s *question.QuestionSet, oid uint32) error {
	query := "UPDATE question_qset SET question_id=?, qset_id=?, created_by_id=? , created_date=? , updated_date=? WHERE id=?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return status.Errorf(codes.Internal, "Error updating: %v", err)
	}
	res, err := stmt.Exec(s.QuestionId, s.SetId, s.CreatedById, s.CreatedDate, s.UpdatedDate, oid)
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
	query := "SELECT question_id, qset_id, created_by_id, created_date, updated_date FROM question_qset"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Retrieve(db *sql.DB, oid uint32) (*question.QuestionSet, error) {
	query := "SELECT question_id, qset_id, created_by_id, created_date, updated_date FROM question_qset WHERE id=?"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row := models.Dbclient.QueryRow(query, oid)

	var questionSetItem = question.QuestionSet{}
	er := row.Scan(&questionSetItem.QuestionId, &questionSetItem.SetId, &questionSetItem.CreatedById, &questionSetItem.CreatedDate, &questionSetItem.UpdatedDate)
	if er != nil {
		log.Printf("Error when getting the object %v", er.Error())
		return nil, er
	}
	return &questionSetItem, nil
}

func Delete(db *sql.DB, oid uint32) error {
	query := "DELETE FROM question_qset WHERE id=?"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Exec(query, oid)
	if err != nil {
		return err
	}
	return err
}
