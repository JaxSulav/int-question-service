package main

import "questionService/models/migrations/tables"

func main() {
	err := tables.CreateTypeTable()
	if err != nil {
		panic(err.Error())
	}

	err = tables.CreateQuestionTable()
	if err != nil {
		panic(err.Error())
	}

	err = tables.CreateQsetTable()
	if err != nil {
		panic(err.Error())
	}

	err = tables.CreateQuestionQsetTable()
	if err != nil {
		panic(err.Error())
	}

}
