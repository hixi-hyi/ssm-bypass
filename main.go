package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Parameter struct {
	Name  string
	Value string
}

func main() {
	flag.Parse()
	args := flag.Args()
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	svc := ssm.New(sess)
	var parameters []Parameter
	nextToken := ""
	for {
		req := &ssm.GetParametersByPathInput{
			Path:           aws.String(args[0]),
			WithDecryption: aws.Bool(true),
		}
		if nextToken != "" {
			req.NextToken = aws.String(nextToken)
		}
		res, err := svc.GetParametersByPath(req)
		if err != nil {
			panic(err)
		}
		for _, v := range res.Parameters {
			arr := strings.Split(*v.Name, "/")
			name := arr[len(arr)-1]
			parameters = append(parameters, Parameter{Name: name, Value: *v.Value})
		}
		if res.NextToken == nil {
			break
		}
		nextToken = *res.NextToken
	}
	for _, v := range parameters {
		fmt.Printf("export %s=%s\n", v.Name, v.Value)
	}
}
