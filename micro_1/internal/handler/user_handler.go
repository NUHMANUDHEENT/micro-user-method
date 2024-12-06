package handler

import (
	"context"
	"errors"
	"fmt"
	"golang-application/internal/model"
	"golang-application/internal/repo"
	userpb "golang-application/proto"
)

type UserServiceClient struct {
	userpb.UnimplementedUserServiceServer
	repo repo.UserRepository
}

func NewUserHandler(repo repo.UserRepository) *UserServiceClient {
	return &UserServiceClient{
		repo: repo,
	}
}

func (h *UserServiceClient) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := req.User
	userData := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}
	fmt.Println("user", userData)

	userid, err := h.repo.CreateUser(userData)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		Message: "User created successfully",
		Id:      int32(userid),
	}, nil
}

func (h *UserServiceClient) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.repo.ListUser()
	if err != nil {
		return nil, err
	}

	var userList []*userpb.User
	for _, user := range users {
		userList = append(userList, &userpb.User{
			Id:    int32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &userpb.ListUsersResponse{
		Users: userList,
	}, nil
}

// GetUser retrieves a user by ID
func (h *UserServiceClient) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId := req.GetUserId()
	fmt.Println("userid", userId)
	// Fetch user from the repository by ID
	user, err := h.repo.GetUser(int(userId))
	if err != nil {
		return nil, err
	}

	// Return the user details in the response
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    int32(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
	}, nil
}

// UpdateUser updates an existing user
func (h *UserServiceClient) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	userData := req.GetUser()
	user := model.User{
		Name:  req.User.Name,
		Email: req.User.Email,
	}
	user.ID = uint(userData.Id)

	// Call repository to update the user details
	err := h.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	// Return success message in response
	return &userpb.UpdateUserResponse{
		Message: "User updated successfully",
	}, nil
}

// DeleteUser deletes a user by ID
func (h *UserServiceClient) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	userId := req.GetUserId()

	// Call repository to delete the user
	err := h.repo.DeleteUser(int(userId))
	if err != nil {
		return nil, err
	}

	// Return success message in response
	return &userpb.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
func (h *UserServiceClient) ListUserNames(ctx context.Context, req *userpb.ListUserNamesRequest) (*userpb.ListUserNamesResponse, error) {

	resp, err := h.repo.ListUserNames()
	if err != nil {
		return nil, errors.New("invalid method")

	}
	return &userpb.ListUserNamesResponse{Names: resp}, nil
}
