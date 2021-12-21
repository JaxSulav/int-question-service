package main

import (
	"context"
	"log"
	"questionService/models"
	"time"
)

var Dbclient = models.DbConn()

func CreateQuestionTable() error {
	query := `CREATE TABLE IF NOT EXISTS question(
				id INT PRIMARY KEY AUTO_INCREMENT, 
				title TEXT,
       			content TEXT, 
				created_by_id INT,
				created_date DATE,
				updated_date DATE,
				active BOOLEAN
			)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Dbclient.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	log.Printf("Table %v created", "question")
	return nil
}

func CreateQuestionSetTable() error {
	query := `CREATE TABLE IF NOT EXISTS questionset(
				id INT PRIMARY KEY AUTO_INCREMENT, 
				set_id INT,
				created_by_id INT,
				created_date DATE,
				updated_date DATE
			)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Dbclient.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	log.Printf("Table %v created", "questionset")
	return nil
}

func CreateQuestionMtmQuestionSetTable() error {
	query := `CREATE TABLE IF NOT EXISTS question_questionset(
				id INT PRIMARY KEY AUTO_INCREMENT, 
				question_id INT,
				questionset_id INT,
				FOREIGN KEY (question_id) REFERENCES questionset(id) ON DELETE CASCADE,
				FOREIGN KEY (questionset_id) REFERENCES question(id) ON DELETE CASCADE
			)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Dbclient.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	log.Printf("Table %v created", "question_questionset")
	return nil
}
