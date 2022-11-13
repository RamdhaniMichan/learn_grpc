package controller

import (
	"context"
	"errors"
	"fmt"
	"grpc_hello_world/entity"
	"grpc_hello_world/infrastructure/dao"
	"grpc_hello_world/model/user"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserService struct {
	user.UnimplementedUserServiceServer
	Db   *gorm.DB
	repo *dao.Repository
}

func NewUserService(repo *dao.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GreetUser(ctx context.Context, req *user.GreetingRequest) (*user.GreetingResponse, error) {

	var users entity.User

	users.Name = req.Name
	users.Salutation = req.Salutation

	validate := users.ValidateSave()

	if validate != nil {
		st := status.New(codes.InvalidArgument, "Invalid Argument")

		v := &errdetails.BadRequest_FieldViolation{
			Field:       "Error Information",
			Description: validate.Error(),
		}

		br := &errdetails.BadRequest{}
		br.FieldViolations = append(br.FieldViolations, v)
		st, _ = st.WithDetails(br)
		return nil, st.Err()
	}

	err := u.repo.UserRepo.Create(ctx, &users)

	// userRes, err := repository.SaveUser(u.Db, &users)

	if err != nil {
		return nil, err
	}

	salutationMessage := fmt.Sprintf("Howdy, %v %v, nice to see you in the future!",
		users.Name, users.Salutation)

	return &user.GreetingResponse{GreetingMessage: salutationMessage}, nil

}

func (u *UserService) GreetUpdate(ctx context.Context, req *user.GreetingRequest) (*user.GreetingResponse, error) {
	var users entity.User

	users.Name = req.Name
	users.Salutation = req.Salutation

	validate := users.ValidateSave()

	if validate != nil {
		st := status.New(codes.InvalidArgument, "Invalid Argument")

		v := &errdetails.BadRequest_FieldViolation{
			Field:       "Error Information",
			Description: validate.Error(),
		}

		br := &errdetails.BadRequest{}
		br.FieldViolations = append(br.FieldViolations, v)
		st, _ = st.WithDetails(br)
		return nil, st.Err()
	}

	_, err := u.repo.UserRepo.Update(ctx, &users)

	// userRes, err := repository.SaveUser(u.Db, &users)

	if err != nil {
		return nil, err
	}

	salutationMessage := fmt.Sprintf("Data %v success update",
		users.UUID)

	return &user.GreetingResponse{GreetingMessage: salutationMessage}, nil
}

func (u *UserService) GreetByID(ctx context.Context, req *user.GreetingRequestID) (*user.GreetingResponseID, error) {

	userID, err := u.repo.UserRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &user.GreetingResponseID{Name: userID.Name, Salutation: userID.Salutation}, nil

}

func (u *UserService) GreetDeleteByID(ctx context.Context, req *user.GreetingRequestID) (*user.GreetingResponse, error) {

	var users entity.User

	users.UUID = req.Id

	validate := users.ValidateFind()

	if validate != nil {
		return nil, errors.New(validate.Error())
	}

	err := u.repo.UserRepo.DeleteByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	delete := fmt.Sprintf("data %v sucessfully delete", users.UUID)

	return &user.GreetingResponse{GreetingMessage: delete}, nil

}

func (u *UserService) GreetAllUser(ctx context.Context, req *user.GreetingAllRequest) (*user.Greetings, error) {

	var allUser []*user.GreetingRequest

	users, err := u.repo.UserRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, u2 := range users {
		allUser = append(allUser, &user.GreetingRequest{
			Name:       u2.Name,
			Salutation: u2.Salutation,
		})
	}

	return &user.Greetings{Greetings: allUser}, nil
}
