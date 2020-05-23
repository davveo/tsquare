package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	proto "github.com/zbrechave/tsquare/srv/question-srv/proto/question"
)

type QuestionService struct{}

func Init() {

}

func (qs *QuestionService) CreateQuestion(ctx context.Context, req *proto.Request, resp *proto.Response) (err error) {
	log.Info("[收到创建请求 in CreateQuestion]")
	return
}
