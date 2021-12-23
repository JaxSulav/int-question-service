package methods

import (
	"context"
	"log"
	question "questionService/libs"
	"questionService/models"
	_questionQset "questionService/models/actions/questionQset"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"strconv"
)

func (*Server) CreateQuestionSet(ctx context.Context, req *question.CreateQuestionSetRequest) (*question.CreateQuestionSetResponse, error) {
	reqData := req.GetQuestionSet()

	data := question.QuestionSet{
		QuestionId:  reqData.GetQuestionId(),
		SetId:       reqData.GetSetId(),
		CreatedById: reqData.GetCreatedById(),
		CreatedDate: reqData.GetCreatedDate(),
		UpdatedDate: reqData.GetUpdatedDate(),
	}

	_, err := _questionQset.Insert(models.Dbclient, &data)
	if err != nil {
		log.Printf("Insert item failed with error %s", err)
		return nil, err
	}

	response := &question.CreateQuestionSetResponse{QuestionSet: &data}
	return response, nil
}

func (*Server) UpdateQuestionSet(ctx context.Context, req *question.UpdateQuestionSetRequest) (*question.UpdateQuestionSetResponse, error) {
	reqData := req.GetQuestionSet()

	data := question.QuestionSet{
		QuestionId:  reqData.GetQuestionId(),
		SetId:       reqData.GetSetId(),
		CreatedById: reqData.GetCreatedById(),
		CreatedDate: reqData.GetCreatedDate(),
		UpdatedDate: reqData.GetUpdatedDate(),
	}
	oid := req.GetId()

	err := _questionQset.Update(models.Dbclient, &data, oid)
	if err != nil {
		return nil, err
	}
	response := &question.UpdateQuestionSetResponse{QuestionSet: &data}
	return response, nil
}

func (*Server) ListQuestionSet(ctx context.Context, req *question.ListQuestionSetRequest) (*question.ListQuestionSetResponse, error) {
	rows, err := _questionQset.List(models.Dbclient)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected Error Occurred: %v", err)
	}
	var responses []*question.QuestionSet
	defer func() {
		responses = nil
	}()

	for rows.Next() {
		data := question.QuestionSet{}
		err = rows.Scan(&data.QuestionId, &data.SetId, &data.CreatedById, &data.CreatedDate, &data.UpdatedDate)
		if err != nil {
			log.Printf("Error when getting the object %v", err.Error())
			return nil, err
		}
		responses = append(responses, &data)
	}

	response := &question.ListQuestionSetResponse{
		QuestionSet: responses,
	}
	return response, nil
}

func (*Server) RetrieveQuestionSet(ctx context.Context, req *question.RetrieveQuestionSetRequest) (*question.RetrieveQuestionSetResponse, error) {
	oid := req.GetId()
	res, err := _questionQset.Retrieve(models.Dbclient, oid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve: %v", err.Error())
	}
	response := &question.RetrieveQuestionSetResponse{QuestionSet: res}
	return response, nil
}

func (*Server) DeleteQuestionSet(ctx context.Context, req *question.DeleteQuestionSetRequest) (*question.DeleteQuestionSetResponse, error) {
	oid := req.GetId()

	err := _questionQset.Delete(models.Dbclient, oid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete: %v", err)
	}
	res := "Successfully deleted item with id: " + strconv.Itoa(int(oid))
	response := &question.DeleteQuestionSetResponse{
		Success:  true,
		Response: res,
	}
	return response, nil
}
