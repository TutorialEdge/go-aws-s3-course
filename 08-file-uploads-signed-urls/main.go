package main

import "fmt"

func main() {
	fmt.Println("File Uploads using Signed URLs")
}
package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	fmt.Println("File Uploads with Presigned URLs")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal("Could not get session")
	}
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
        Bucket: aws.String("myBucket"),
        Key:    aws.String("myKey"),
    })
    urlStr, err := req.Presign(15 * time.Minute)

	log.Printf("Result: %+v\n", urlStr)
}
