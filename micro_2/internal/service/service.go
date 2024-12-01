package service

import (
	"context"
	"errors"
	models "golang-application/internal/model"
	userpb "golang-application/proto"
	"log"
	"sync"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.User) (*userpb.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req models.User) (*userpb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, id int64) (*userpb.DeleteUserResponse, error)
	ListUser(ctx context.Context) (*userpb.ListUsersResponse, error)
	GetUser(ctx context.Context, id int64) (models.User, error)
	ListUserNames(ctx context.Context, method, waitTime int) ([]string, error)
}

type userService struct {
	UserClient userpb.UserServiceClient
	Mutx       sync.Mutex
}

func NewUserService(userClient userpb.UserServiceClient) UserService {
	return &userService{
		UserClient: userClient,
	}
}

// CreateUser implementation
func (s *userService) CreateUser(ctx context.Context, req models.User) (*userpb.CreateUserResponse, error) {

	user := userpb.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	res, err := s.UserClient.CreateUser(ctx, &userpb.CreateUserRequest{User: &user})
	if err != nil {
		return &userpb.CreateUserResponse{}, err
	}
	return res, nil
}

// UpdateUser implementation
func (s *userService) UpdateUser(ctx context.Context, req models.User) (*userpb.UpdateUserResponse, error) {
	user := userpb.User{
		Id:       int32(req.ID),
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	res, err := s.UserClient.UpdateUser(ctx, &userpb.UpdateUserRequest{User: &user})
	if err != nil {
		return &userpb.UpdateUserResponse{}, err
	}
	return res, nil
}

// DeleteUser implementation
func (s *userService) DeleteUser(ctx context.Context, id int64) (*userpb.DeleteUserResponse, error) {
	res, err := s.UserClient.DeleteUser(ctx, &userpb.DeleteUserRequest{UserId: int32(id)})
	if err != nil {
		return &userpb.DeleteUserResponse{}, err
	}
	return res, nil
}

// ListUser implementation
func (s *userService) ListUser(ctx context.Context) (*userpb.ListUsersResponse, error) {
	res, err := s.UserClient.ListUsers(ctx, &userpb.ListUsersRequest{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUser implementation
func (s *userService) GetUser(ctx context.Context, id int64) (models.User, error) {
	res, err := s.UserClient.GetUser(ctx, &userpb.GetUserRequest{UserId: int32(id)})
	if err != nil {
		return models.User{}, err
	}
	userData := models.User{
		Name:  res.User.Name,
		Email: res.User.Email,
		Phone: res.User.Phone,
	}
	userData.ID = uint(res.User.Id)

	return userData, nil
}
func (s *userService) ListUserNames(ctx context.Context, method, waitTime int) ([]string, error) {
	switch method {
	case 1:
		log.Printf("Request from method 1 for %d second waitTime", waitTime)

		
		select {
		case <-time.After(time.Duration(waitTime) * time.Second): 
		case <-ctx.Done():
			return nil, ctx.Err()
		}

		s.Mutx.Lock()
		defer s.Mutx.Unlock()

		res, err := s.UserClient.ListUserNames(ctx, &userpb.ListUserNamesRequest{
			Method:   int32(method),
			WaitTime: int32(waitTime),
		})
		if err != nil {
			return nil, err
		}
		return res.Names, nil

	case 2:
		log.Printf("Request from method 2 for %d second waitTime", waitTime)

		resultCh := make(chan *userpb.ListUserNamesResponse, 1)
		errorCh := make(chan error, 1)

		go func() {
			res, err := s.UserClient.ListUserNames(ctx, &userpb.ListUserNamesRequest{
				Method:   int32(method),
				WaitTime: int32(waitTime),
			})
			if err != nil {
				errorCh <- err
				return
			}
			resultCh <- res
		}()

		// Deliberately wait for the given time
		select {
		case <-time.After(time.Duration(waitTime) * time.Second): // Wait for specified seconds
		case <-ctx.Done(): // Handle cancellations
			return nil, ctx.Err()
		}

		// Process the response
		select {
		case res := <-resultCh:
			return res.Names, nil
		case err := <-errorCh:
			return nil, err
		case <-ctx.Done():
			return nil, ctx.Err()
		}

	default:
		return nil, errors.New("invalid method")
	}
}
