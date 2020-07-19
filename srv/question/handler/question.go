package handler

import (
	"context"

	proto "github.com/davveo/tsquare/proto/question"
	log "github.com/micro/go-micro/v2/logger"
)

type QuestionService struct{}

func Init() {

}

func (qs *QuestionService) CreateQuestion(ctx context.Context, req *proto.Request, resp *proto.Response) (err error) {
	log.Info("[收到创建请求 in CreateQuestion]")
	return
}
