package minio1

import (
	//"context"
	"demo/tiktok/pkg/constants"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func Init() {
	var err error
	MinioClient, err = minio.New(constants.MinioAddress, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioUserName, constants.MinioPassword, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}

	/*

		err = MinioClient.MakeBucket(context.Background(), "videos", minio.MakeBucketOptions{})
		if err != nil {
			panic(err)
		}
		err = MinioClient.MakeBucket(context.Background(), "images", minio.MakeBucketOptions{})
		if err != nil {
			panic(err)
		}

		policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::my-bucketname/*"],"Sid": ""}]}`
		err = MinioClient.SetBucketPolicy(context.Background(), "videos", policy)
		if err != nil {
			panic(err)
		}
		err = MinioClient.SetBucketPolicy(context.Background(), "images", policy)
		if err != nil {
			panic(err)
		}

	*/
}
