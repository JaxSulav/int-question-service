package methods

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (*Server) UpdateType(ctx context.Context, req *question.UpdateTypeRequest) (*question.UpdateTypeResponse, error) {
	reqData := req.GetType()

	data := question.Type{
		Name:        reqData.Name,
		CreatedById: reqData.CreatedById,
		CreatedDate: reqData.CreatedDate,
		UpdatedDate: reqData.UpdatedDate,
		Active:      reqData.Active,
	}
	oid := req.GetId()

	err := _type.Update(models.Dbclient, &data, oid)
	if err != nil {
		return nil, err
	}
	response := &question.UpdateTypeResponse{Type: &data}
	return response, nil
}

func (*Server) ListType(ctx context.Context, req *question.ListTypeRequest) (*question.ListTypeResponse, error) {
	rows, err := _type.List(models.Dbclient)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected Error Occurred: %v", err)
	}
	var i *uint32
	var responses []*question.Type
	defer func() {
		responses = nil
		i = nil
	}()

	for rows.Next() {
		data := question.Type{}
		err = rows.Scan(&i, &data.Name, &data.CreatedById, &data.CreatedDate, &data.UpdatedDate, &data.Active)
		if err != nil {
			log.Printf("Error when getting the object %v", err.Error())
			return nil, err
		}
		responses = append(responses, &data)
	}

	response := &question.ListTypeResponse{
		Type: responses,
	}
	return response, nil
}
