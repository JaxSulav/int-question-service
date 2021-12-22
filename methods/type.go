package methods

import (
	"context"
	"log"
	question "questionService/libs"
	"questionService/models"
	"questionService/models/actions/type"
)

func (*Server) CreateType(ctx context.Context, req *question.CreateTypeRequest) (*question.CreateTypeResponse, error) {
	reqData := req.GetType()

	data := question.Type{
		Name:        reqData.GetName(),
		CreatedById: reqData.GetCreatedById(),
		CreatedDate: reqData.GetCreatedDate(),
		UpdatedDate: reqData.GetUpdatedDate(),
		Active:      reqData.GetActive(),
	}

	_, err := _type.Insert(models.Dbclient, &data)
	if err != nil {
		log.Printf("Insert item failed with error %s", err)
		return nil, err
	}

	response := &question.CreateTypeResponse{Type: &data}
	return response, nil
}
