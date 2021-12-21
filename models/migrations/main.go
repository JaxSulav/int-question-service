package main

func main() {
	err := CreateQuestionTable()
	if err != nil {
		panic(err.Error())
	}

	err = CreateQuestionSetTable()
	if err != nil {
		panic(err.Error())
	}

	err = CreateQuestionMtmQuestionSetTable()
	if err != nil {
		panic(err.Error())
	}

}
