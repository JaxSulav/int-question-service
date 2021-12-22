package _type

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	question "questionService/libs"
	"time"
)

func Insert(db *sql.DB, t *question.Type) (int64, error) {
	query := "INSERT INTO type(name, created_by_id, created_date, updated_date, active) VALUES (?, ?, ?, ?, ?)"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, t.Name, t.CreatedById, t.CreatedDate, t.UpdatedDate, t.Active)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
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

func Update(db *sql.DB, t *question.Type, oid uint32) error {
	query := "UPDATE type SET name=?, created_by_id=?, created_date=?, updated_date=?, active=? WHERE id=?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return status.Errorf(codes.Internal, "Error updating: %v", err)
	}
	res, err := stmt.Exec(t.Name, t.CreatedById, t.CreatedDate, t.UpdatedDate, t.Active, oid)
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
	query := "SELECT * FROM type"
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
