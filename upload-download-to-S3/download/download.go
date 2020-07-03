package download

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	connect "upload-download-to-S3/connect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	awsS3Bucket = "ram-test-0001"
)

//Downloading fucn is used to download file from AWS S3
func Downloading(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Downloading ...")
	filename := strings.Replace(r.URL.Path, "/get/", "B8479642CA_output.pdf", 1)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("error while creating file ", err)
		return
	}
	var AWSSession = connect.ConnectAWS()
	downloader := s3manager.NewDownloader(AWSSession)
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(awsS3Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		fmt.Println("error while retrieving", err)
		return
	}

	http.ServeFile(w, r, filename)
}
