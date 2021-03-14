package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func listBucketItems(sess *session.Session) {
	svc := s3.New(sess)
	resp, err := svc.ListObjectsV2(
		&s3.ListObjectsV2Input{
			Bucket: aws.String("go-aws-s3-course"),
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}

func uploadItem(sess *session.Session) {
	f, err := os.Open("03-downloads-and-uploads/my-file.txt")
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
}

func downloadItem(sess *session.Session) {
	file, err := os.Create("03-downloads-and-uploads/downloaded.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	downloader := s3manager.NewDownloader(sess)

	// number of bytes downloaded or error
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String("go-aws-s3-course"),
			Key:    aws.String("my-file.txt"),
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Successfully downloaded")
}

func main() {
	fmt.Println("Downloads and Uploads")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatal("Could not get session")
	}

	uploadItem(sess)
	listBucketItems(sess)
	downloadItem(sess)
}
