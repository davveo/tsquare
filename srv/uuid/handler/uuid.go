package handler

import (
	"context"

	uuid_proto "github.com/zbrechave/tsquare/srv/uuid/proto/uuid"
)

type Uuid struct{}

func (e *Uuid) Call(ctx context.Context, req *uuid_proto.Request, rsp *uuid_proto.Response) error {

	return nil
}
