package main

import (
	"context"

	"github.com/bufbuild/connect-go"

	todov1 "github.com/dillendev/training-connect-go/gen/todo/v1"
)

type TodoService struct{}

func (s *TodoService) ListTasks(ctx context.Context, req *connect.Request[todov1.ListTasksRequest]) (*connect.Response[todov1.ListTasksResponse], error) {
	return connect.NewResponse(&todov1.ListTasksResponse{Tasks: []*todov1.ListTasksResponse_Task{
		{
			Id:   1,
			Name: "Buy milk",
		},
	}}), nil
}
