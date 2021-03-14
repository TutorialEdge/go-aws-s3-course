package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	fmt.Println("Understanding ACLs")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal("Could not get session")
	}
	svc := s3.New(sess)

	rule := s3.CORSRule{
		AllowedHeaders: aws.StringSlice([]string{"Authorization"}),
		AllowedOrigins: aws.StringSlice([]string{"*"}),
		MaxAgeSeconds:  aws.Int64(3000),
		AllowedMethods: aws.StringSlice([]string{"PUT", "POST"}),
	}

	params := s3.PutBucketCorsInput{
		Bucket: aws.String("go-aws-s3-course"),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{&rule},
		},
	}

	result, err := svc.PutBucketCors(&params)
	if err != nil {
		log.Fatal("Unable to update bucket CORS information")
	}

	log.Printf("Result: %+v\n", result)
	log.Println("Bucket CORS Information updated!")

}
