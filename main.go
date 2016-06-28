package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"os/exec"
)

func main() {

	// setup string flag for config file
	configfile := flag.String("configfile", "config.ini", "config file")
	flag.Parse()

	// get string
	cf := *configfile
	fmt.Println("config file: ", cf)

	// read config
	config := ReadConfig(cf)
	fmt.Println(config.Endpoint)

	//take photo
	u1 := uuid.NewV4()
	fname := fmt.Sprintf("%s.jpg", u1)

	// actually take the photo
	args := []string{config.CameraCommand}
	args = append(args, fname)
	fmt.Printf("Camera args: %s", args)

	out, err := exec.Command("fswebcam", args...).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Took Photo: %s\n", string(out))

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	uploader := s3manager.NewUploader(session.New(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
	}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   file,
		Bucket: aws.String(config.bucket),
		Key:    aws.String("mykey"),
	})
	if err != nil {
		log.Fatal("Failed to uplod", err)
	}
	log.Println("sucessfully uploaded to", result.Location)

}
