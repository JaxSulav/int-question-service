package methods

import (
	"context"
	"log"
	question "questionService/libs"
	"questionService/models"
	_question "questionService/models/actions/question"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"strconv"
)

func (*Server) CreateQuestion(ctx context.Context, req *question.CreateQuestionRequest) (*question.CreateQuestionResponse, error) {
	reqData := req.GetQuestion()

	data := question.Question{
		Title:       reqData.GetTitle(),
		Content:     reqData.GetContent(),
		CreatedById: reqData.GetCreatedById(),
		CreatedDate: reqData.GetCreatedDate(),
		UpdatedDate: reqData.GetUpdatedDate(),
		Active:      reqData.GetActive(),
		Type:        reqData.GetType(),
	}

	_, err := _question.Insert(models.Dbclient, &data)
	if err != nil {
		log.Printf("Insert item failed with error %s", err)
		return nil, err
	}

	response := &question.CreateQuestionResponse{Question: &data}
	return response, nil
}

func (*Server) UpdateQuestion(ctx context.Context, req *question.UpdateQuestionRequest) (*question.UpdateQuestionResponse, error) {
	reqData := req.GetQuestion()

	data := question.Question{
		Title:       reqData.GetTitle(),
		Content:     reqData.GetContent(),
		CreatedById: reqData.GetCreatedById(),
		CreatedDate: reqData.GetCreatedDate(),
		UpdatedDate: reqData.GetUpdatedDate(),
		Active:      reqData.GetActive(),
		Type:        reqData.GetType(),
	}
	oid := req.GetId()

	err := _question.Update(models.Dbclient, &data, oid)
	if err != nil {
		return nil, err
	}
	response := &question.UpdateQuestionResponse{Question: &data}
	return response, nil
}

func (*Server) ListQuestion(ctx context.Context, req *question.ListQuestionRequest) (*question.ListQuestionResponse, error) {
	rows, err := _question.List(models.Dbclient)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected Error Occurred: %v", err)
	}
	var responses []*question.Question
	defer func() {
		responses = nil
	}()

	for rows.Next() {
		data := question.Question{}
		err = rows.Scan(&data.Title, &data.Content, &data.CreatedById, &data.CreatedDate, &data.UpdatedDate, &data.Active, &data.Type)
		if err != nil {
			log.Printf("Error when getting the object %v", err.Error())
			return nil, err
		}
		responses = append(responses, &data)
	}

	response := &question.ListQuestionResponse{
		Question: responses,
	}
	return response, nil
}

func (*Server) RetrieveQuestion(ctx context.Context, req *question.RetrieveQuestionRequest) (*question.RetrieveQuestionResponse, error) {
	oid := req.GetId()
	res, err := _question.Retrieve(models.Dbclient, oid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve: %v", err.Error())
	}
	response := &question.RetrieveQuestionResponse{Question: res}
	return response, nil
}

func (*Server) DeleteQuestion(ctx context.Context, req *question.DeleteQuestionRequest) (*question.DeleteQuestionResponse, error) {
	oid := req.GetId()

	err := _question.Delete(models.Dbclient, oid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete: %v", err)
	}
	res := "Successfully deleted item with id: " + strconv.Itoa(int(oid))
	response := &question.DeleteQuestionResponse{
		Success:  true,
		Response: res,
	}
	return response, nil
}
