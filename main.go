package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
    flag.Parse()
    args := flag.Args()
    sess, err := session.NewSession()
    if err != nil {
        panic(err)
    }
    svc := ssm.New(sess)
    res, err := svc.GetParametersByPath(&ssm.GetParametersByPathInput{
        Path: aws.String(args[0]),
        WithDecryption: aws.Bool(true),
    })
    if err != nil {
        panic(err)
    }
    for _, v := range res.Parameters {
        arr := strings.Split(*v.Name, "/")
        name := arr[len(arr)-1]
        fmt.Printf("export %s=%s\n", name, *v.Value)
    }
}
