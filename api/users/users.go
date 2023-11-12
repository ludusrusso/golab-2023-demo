package users

import (
	"context"
	"errors"
	v1 "golab-2023/gen/proto/golab2023/users/v1"
	c "golab-2023/gen/proto/golab2023/users/v1/usersv1connect"
	"golab-2023/internal/queries"

	connect "github.com/bufbuild/connect-go"
	"github.com/jackc/pgx/v5"
)

type srv struct {
	q *queries.Queries
}

func NewUsersServiceServer(q *queries.Queries) c.UsersServiceHandler {
	return &srv{
		q: q,
	}
}

func (s *srv) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	user, err := s.q.CreateUser(ctx, req.Msg.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		} else {
			return nil, err
		}
	}

	return connect.NewResponse(&v1.CreateUserResponse{
		User: user.Pb(),
	}), nil
}

func (s *srv) DeleteUser(ctx context.Context, req *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	user, err := s.q.DeleteUser(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		} else {
			return nil, err
		}
	}

	return connect.NewResponse(&v1.DeleteUserResponse{
		User: user.Pb(),
	}), nil
}

func (s *srv) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	user, err := s.q.GetUser(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		} else {
			return nil, err
		}
	}

	return connect.NewResponse(&v1.GetUserResponse{
		User: user.Pb(),
	}), nil
}

func (s *srv) ListUsers(ctx context.Context, req *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	users, err := s.q.ListUsers(ctx, queries.ListUsersParams{
		Offset: req.Msg.Offset,
		Limit:  req.Msg.Limit,
	})
	if err != nil {
		return nil, err
	}

	total, err := s.q.CountUsers(ctx)
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*v1.User, 0, len(users))
	for _, user := range users {
		pbUsers = append(pbUsers, user.Pb())
	}

	return connect.NewResponse(&v1.ListUsersResponse{
		Totat: int32(total),
		Users: pbUsers,
	}), nil
}
