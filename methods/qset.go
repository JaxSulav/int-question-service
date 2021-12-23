package methods

import (
	"context"
	"log"
	question "questionService/libs"
	"questionService/models"
	_qset "questionService/models/actions/qset"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"strconv"
)

func (*Server) CreateSet(ctx context.Context, req *question.CreateSetRequest) (*question.CreateSetResponse, error) {
	reqData := req.GetSet()

	data := question.Set{
		Time:        reqData.GetTime(),
		Type:        reqData.GetType(),
		CreatedById: reqData.GetCreatedById(),
		CreatedDate: reqData.GetCreatedDate(),
		UpdatedDate: reqData.GetUpdatedDate(),
		Active:      reqData.GetActive(),
		QsName:      reqData.GetQsName(),
	}

	_, err := _qset.Insert(models.Dbclient, &data)
	if err != nil {
		log.Printf("Insert item failed with error %s", err)
		return nil, err
	}

	response := &question.CreateSetResponse{Set: &data}
	return response, nil
}

func (*Server) UpdateSet(ctx context.Context, req *question.UpdateSetRequest) (*question.UpdateSetResponse, error) {
	reqData := req.GetSet()

	data := question.Set{
		Time:        reqData.GetTime(),
		Type:        reqData.GetType(),
		CreatedById: reqData.GetCreatedById(),
		CreatedDate: reqData.GetCreatedDate(),
		UpdatedDate: reqData.GetUpdatedDate(),
		Active:      reqData.GetActive(),
		QsName:      reqData.GetQsName(),
	}
	oid := req.GetId()

	err := _qset.Update(models.Dbclient, &data, oid)
	if err != nil {
		return nil, err
	}
	response := &question.UpdateSetResponse{Set: &data}
	return response, nil
}

func (*Server) ListSet(ctx context.Context, req *question.ListSetRequest) (*question.ListSetResponse, error) {
	rows, err := _qset.List(models.Dbclient)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected Error Occurred: %v", err)
	}
	var responses []*question.Set
	defer func() {
		responses = nil
	}()

	for rows.Next() {
		data := question.Set{}
		err = rows.Scan(&data.Time, &data.Type, &data.CreatedById, &data.CreatedDate, &data.UpdatedDate, &data.Active, &data.QsName)
		if err != nil {
			log.Printf("Error when getting the object %v", err.Error())
			return nil, err
		}
		responses = append(responses, &data)
	}

	response := &question.ListSetResponse{
		Set: responses,
	}
	return response, nil
}

func (*Server) RetrieveSet(ctx context.Context, req *question.RetrieveSetRequest) (*question.RetrieveSetResponse, error) {
	oid := req.GetId()
	res, err := _qset.Retrieve(models.Dbclient, oid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve: %v", err.Error())
	}
	response := &question.RetrieveSetResponse{Set: res}
	return response, nil
}

func (*Server) DeleteSet(ctx context.Context, req *question.DeleteSetRequest) (*question.DeleteSetResponse, error) {
	oid := req.GetId()

	err := _qset.Delete(models.Dbclient, oid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete: %v", err)
	}
	res := "Successfully deleted item with id: " + strconv.Itoa(int(oid))
	response := &question.DeleteSetResponse{
		Success:  true,
		Response: res,
	}
	return response, nil
}
