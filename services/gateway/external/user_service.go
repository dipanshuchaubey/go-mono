package external

import (
	v1 "carthage/protos/user_service"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceInterface interface {
	GetUsers(ctx context.Context, req *v1.GetUsersRequest) (*v1.GetUsersResponse, error)
	GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error)
}

type UserServiceClient struct {
	us v1.UserServiceClient
}

func UserService() UserServiceInterface {
	conn, conErr := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if conErr != nil {
		fmt.Printf("Error connecting on user service: %v\n", conErr.Error())
	}

	c := v1.NewUserServiceClient(conn)

	return &UserServiceClient{c}
}

func (u *UserServiceClient) GetUsers(ctx context.Context, req *v1.GetUsersRequest) (*v1.GetUsersResponse, error) {
	fmt.Println("Calling User Service with: ", req)

	res, err := u.us.GetUsers(ctx, req)
	if err != nil {
		errMsg := fmt.Errorf("error calling GetUsers for UserIDs %s: %v", req.GetUserIds(), err)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	return res, nil
}

func (u *UserServiceClient) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	fmt.Println("Calling User Service with: ", req)

	res, err := u.us.GetUser(ctx, req)
	if err != nil {
		errMsg := fmt.Errorf("error calling GetUser for UserID %s: %v", req.GetUserId(), err)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	return res, nil
}
