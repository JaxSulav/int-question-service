package tables

import (
	"context"
	"log"
	"questionService/models"
	"time"
)

func CreateTypeTable() error {
	query := `CREATE TABLE IF NOT EXISTS type(
				id INT PRIMARY KEY AUTO_INCREMENT, 
				name VARCHAR(120),
				created_by_id INT,
				created_date DATE,
				updated_date DATE,
				active BOOLEAN
			)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := models.Dbclient.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	log.Printf("Table %v created", "methods")
	return nil
}

func CreateQuestionTable() error {
	query := `CREATE TABLE IF NOT EXISTS question(
				id INT PRIMARY KEY AUTO_INCREMENT, 
				title VARCHAR(200),
       			content TEXT, 
				created_by_id INT,
				created_date DATE,
				updated_date DATE,
				active BOOLEAN
			)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := models.Dbclient.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	log.Printf("Table %v created", "question")
	return nil
}

func CreateQsetTable() error {
	query := `CREATE TABLE IF NOT EXISTS qset(
				id INT PRIMARY KEY AUTO_INCREMENT, 
				qs_name VARCHAR(120),
				time TIME,
				type_id INT,
				FOREIGN KEY (type_id) REFERENCES type(id) ON DELETE CASCADE,
				created_by_id INT,
				created_date DATE,
				updated_date DATE,
				active BOOLEAN
			)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := models.Dbclient.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	log.Printf("Table %v created", "qset")
	return nil
}

func CreateQuestionQsetTable() error {
	query := `CREATE TABLE IF NOT EXISTS question_qset(
				id INT PRIMARY KEY AUTO_INCREMENT, 
				question_id INT,
				qset_id INT,
				created_by_id INT,
				created_date DATE,
				updated_date DATE,
				FOREIGN KEY (question_id) REFERENCES question(id) ON DELETE CASCADE,
				FOREIGN KEY (qset_id) REFERENCES qset(id) ON DELETE CASCADE
			)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := models.Dbclient.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	log.Printf("Table %v created", "question_qset")
	return nil
}
