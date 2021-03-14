package main

import (
	"fmt"
	"log"
	"time"

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

	rule := s3.Rule{
		Expiration: &s3.LifecycleExpiration{
			Date: aws.Time(time.Now().Add(1 * time.Hour)),
		},
		Prefix: aws.String("*"),
	}

	params := s3.PutBucketLifecycleInput{
		Bucket: aws.String("go-aws-s3-course"),
		LifecycleConfiguration: &s3.LifecycleConfiguration{
			Rules: []*s3.Rule{&rule},
		},
	}

	result, err := svc.PutBucketLifecycle(&params)
	if err != nil {
		log.Fatal("Unable to update bucket CORS information")
	}

	log.Printf("Result: %+v\n", result)
	log.Println("Bucket CORS Information updated!")

}
