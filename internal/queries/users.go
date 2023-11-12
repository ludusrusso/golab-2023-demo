package queries

import (
	v1 "golab-2023/gen/proto/golab2023/users/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *User) Pb() *v1.User {
	return &v1.User{
		Id:        u.ID,
		Name:      u.Name,
		CreatedAt: timestamppb.New(u.CreatedAt.Time),
		UpdatedAt: timestamppb.New(u.UpdatedAt.Time),
	}
}
