package main

import (
	"fmt"
	//	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	// Basic information for the Amazon OpenSearch Service domain
	domain := "" // e.g. https://my-domain.region.es.amazonaws.com
	index := "my-index"
	id := "1"
	endpoint := domain + "/" + index + "/" + "_doc" + "/" + id
	region := "" // e.g. us-east-1
	service := "es"

	// Sample JSON document to be included as the request body
	json := `{ "title": "Thor: Ragnarok", "director": "Taika Waititi", "year": "2017" }`
	body := strings.NewReader(json)

	// Get credentials and create the Signature Version 4 signer
	aws_profile := "default" //your AKSK profile name
	os.Setenv("AWS_PROFILE", aws_profile)
	//	credentials := credentials.NewEnvCredentials()
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	signer := v4.NewSigner(sess.Config.Credentials)

	// An HTTP client for sending the request
	client := &http.Client{}

	// Form the HTTP request
	req, err := http.NewRequest(http.MethodPut, endpoint, body)
	if err != nil {
		fmt.Print(err)
	}

	// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
	req.Header.Add("Content-Type", "application/json")

	// Sign the request, send it, and print the response
	signer.Sign(req, body, service, region, time.Now())
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp.Status + "\n")
}
