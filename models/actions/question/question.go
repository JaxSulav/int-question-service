package _question

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

func Insert(db *sql.DB, q *question.Question) (int64, error) {
	query := "INSERT INTO question(title, content, created_by_id, created_date, updated_date, active, type_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, q.Title, q.Content, q.CreatedById, q.CreatedDate, q.UpdatedDate, q.Active, q.Type)
	if err != nil {
		log.Printf("Error %q when inserting row into table", err)
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

func Update(db *sql.DB, q *question.Question, oid uint32) error {
	query := "UPDATE question SET title=?, content=?, created_by_id=?, created_date=?, updated_date=?, active=?, type_id=? WHERE id=?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return status.Errorf(codes.Internal, "Error updating: %v", err)
	}
	res, err := stmt.Exec(q.Title, q.Content, q.CreatedById, q.CreatedDate, q.UpdatedDate, q.Active, q.Type, oid)
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
	query := "SELECT title, content, created_by_id, created_date, updated_date, active, type_id FROM question"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Retrieve(db *sql.DB, oid uint32) (*question.Question, error) {
	query := "SELECT title, content, created_by_id, created_date, updated_date, active, type_id FROM question WHERE id=?"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row := models.Dbclient.QueryRow(query, oid)

	var questionItem = question.Question{}
	er := row.Scan(&questionItem.Title, &questionItem.Content, &questionItem.CreatedById, &questionItem.CreatedDate, &questionItem.UpdatedDate, &questionItem.Active, &questionItem.Type)
	if er != nil {
		log.Printf("Error when getting the object %v", er.Error())
		return nil, er
	}
	return &questionItem, nil
}

func Delete(db *sql.DB, oid uint32) error {
	query := "DELETE FROM question WHERE id=?"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Exec(query, oid)
	if err != nil {
		return err
	}
	return err
}
