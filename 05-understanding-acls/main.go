package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	fmt.Println("Understanding ACLs")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal("Could not get session")
	}

	f, err := os.Open("05-understanding-acls/my-file.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	defer f.Close()

	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String("go-aws-s3-course"),
		Key:    aws.String("my-file.txt"),
		Body:   f,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Upload Result: %+v\n", result)

	f2, err := os.Open("05-understanding-acls/private.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	defer f.Close()

	result, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("go-aws-s3-course"),
		Key:    aws.String("private.txt"),
		Body:   f2,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Upload Result: %+v\n", result)
}
