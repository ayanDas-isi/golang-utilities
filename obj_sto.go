package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
)

const (
	apiKey_land            = "Cjz7ufQk2e78fmyViRd1JvNkwPvj_LxjOzkQSlzpjzgZ"
	serviceInstanceID_land = "crn:v1:bluemix:public:cloud-object-storage:global:a/54ba5c628eb3b9741efc80bd432510aa:7a4eb081-ebaa-4773-ad79-919eb6d18214::"
	authEndpoint           = "https://iam.cloud.ibm.com/identity/token"
	serviceEndpoint        = "https://s3.us-south.cloud-object-storage.appdomain.cloud"
	apiKey                 = "qpVdYLX8zApUWx7iNqiJ980ZpmwuKpwUVsaXh_s4nW1L"
	serviceInstanceID      = "crn:v1:bluemix:public:cloud-object-storage:global:a/54ba5c628eb3b9741efc80bd432510aa:2750d2b3-fc07-48c5-97d3-c65571fa95d4::"
)

func main() {

	//newBucket := "new-bucketee"
	//newColdBucket := "new-cold-bucketee"

	conf := aws.NewConfig().
		WithEndpoint(serviceEndpoint).
		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(),
			authEndpoint, apiKey, serviceInstanceID)).
		WithS3ForcePathStyle(true)
	start := time.Now()
	//get_full_bucket_content("ups-invoices-logging-bucket-prod", conf)
	get_object(conf, "ups-invoices-logging-bucket-prod", "000a7603-b884-11e9-b0d5-5df8198dbd83")
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	//fmt.Println(c)
}
func get_object(conf *aws.Config, bucketName string, key string) map[string]interface{} {
	// users will need to create bucket, key (flat string name)
	sess := session.Must(session.NewSession())

	client := s3.New(sess, conf)

	Input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	// Call Function
	res, _ := client.GetObject(Input)

	body, _ := ioutil.ReadAll(res.Body)
	s := string(body)
	//fmt.Println(s)
	a := read_json(s)
	fmt.Println(a["date"])
	write_file("output.txt", s)
	return a
}

func list_buckets(conf *aws.Config) {
	sess := session.Must(session.NewSession())

	client := s3.New(sess, conf)

	d, _ := client.ListBuckets(&s3.ListBucketsInput{})
	fmt.Println(*d.Buckets[0].Name)
}

func get_full_bucket_content(bucketName string, conf *aws.Config) *s3.ListObjectsV2Output {
	// Create client
	sess := session.Must(session.NewSession())
	client := s3.New(sess, conf)

	// Bucket Name
	//Bucket := "us-dc-training-prod"
	var filelist []string
	// Call Function
	Input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	//fmt.Println(Input)

	l, e := client.ListObjectsV2(Input)
	//fmt.Println(l)
	fmt.Println(e)
	for file := range l.Contents {
		filelist = append(filelist, *l.Contents[file].Key)
	}
	cnt := 0
	for *l.IsTruncated == true && cnt < 30 {
		cnt += 1
		fmt.Println(cnt)
		Input := &s3.ListObjectsV2Input{
			Bucket:            aws.String(bucketName),
			ContinuationToken: l.ContinuationToken,
		}
		l, _ := client.ListObjectsV2(Input)
		//fmt.Println(e)
		//fmt.Println(l)
		for file := range l.Contents {
			filelist = append(filelist, *l.Contents[file].Key)
		}
	}
	fmt.Println(len(filelist))

	return l
}

func get_bucket_content(bucketName string, conf *aws.Config) *s3.ListObjectsV2Output {
	// Create client
	sess := session.Must(session.NewSession())
	client := s3.New(sess, conf)

	// Bucket Name
	//Bucket := "us-dc-training-prod"

	// Call Function
	Input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	fmt.Println(Input)

	l, e := client.ListObjectsV2(Input)
	//fmt.Println(l)
	fmt.Println(e)

	return l
}
