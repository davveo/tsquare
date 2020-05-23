package handler

import (
	"encoding/json"
	"github.com/micro/go-micro/v2/client"
	auth "github.com/zbrechave/tsquare/srv/auth-srv/proto/auth"
	question "github.com/zbrechave/tsquare/srv/question-srv/proto/question"
	"net/http"
)

var (
	authClient     auth.AuthService
	questionClient question.QuestionService
)

func Init() {
	authClient = auth.NewAuthService("go.micro.service.auth", client.DefaultClient)
	questionClient = question.NewQuestionService("go.micro.service.question", client.DefaultClient)
}

func QuestionCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
