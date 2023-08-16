package main

import(
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type question struct {
	Name string `json:"what is your name?"`
	Age int `json:"How old are you?"`
}

type response struct{
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(q question)(response, error){
	return response{Message: fmt.Sprintf("%s is %d years old", q.Name, q.Age )},nil
}

func main(){
	lambda.Start(HandleLambdaEvent)
}