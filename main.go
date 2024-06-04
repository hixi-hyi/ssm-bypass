package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Parameter struct {
	Name         string
	Value        string
	EscapedValue string
}

func GetParameters(path string) []Parameter {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	svc := ssm.New(sess)

	var parameters []Parameter
	nextToken := ""
	for {
		req := &ssm.GetParametersByPathInput{
			Path:           aws.String(path),
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
	return parameters
}

func escape(value string) string {
	escapedValue := value
	escapedValue = strings.ReplaceAll(escapedValue, "\"", "\\\"")
	escapedValue = strings.ReplaceAll(escapedValue, "\n", "\\n")
	return escapedValue
}

func escapeParameters(parameters []Parameter) []Parameter {
	for i, v := range parameters {
		parameters[i].EscapedValue = escape(v.Value)
	}
	return parameters
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "panic: %v\n", r)
			os.Exit(1)
		}
	}()
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s /your/path/\n", os.Args[0])
		os.Exit(1)
	}
	path := args[0]
	parameters := GetParameters(path)
	parameters = escapeParameters(parameters)
	for _, v := range parameters {
		fmt.Printf("export %s=%s\n", v.Name, v.EscapedValue)
	}
}
