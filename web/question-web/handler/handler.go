package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/client"
	auth "github.com/zbrechave/tsquare/srv/auth-srv/proto/auth"
	question "github.com/zbrechave/tsquare/srv/question-srv/proto/que"
)

var (
	authClient auth.AuthService
)

func Init()  {
	questionClient =
}

func QuestionCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	questionClient := question.NewQuestionService("go.micro.service.question", client.DefaultClient)
	rsp, err := questionClient.Call(context.TODO(), &question.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
