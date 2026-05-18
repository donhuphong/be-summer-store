package database

import (
	appconfig "be-summer-store/internal/config"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	R2Client      *s3.Client
	R2Presign     *s3.PresignClient
	BucketName    string
	PublicBaseURL string
)

func InitR2() {
	cfg := appconfig.AppConfig.R2CloudFare

	endpoint := "https://" + cfg.R2AccountID + ".r2.cloudflarestorage.com"

	awsCfg, err := awsconfig.LoadDefaultConfig(
		context.TODO(),

		awsconfig.WithRegion("auto"),

		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.R2AccessKeyID,
			cfg.R2SecretAccessKey,
			"",
		)),
	)

	if err != nil {
		panic("Không thể khởi tạo R2: " + err.Error())
	}

	R2Client = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})
	R2Presign = s3.NewPresignClient(R2Client)
	BucketName = cfg.R2BucketName
	PublicBaseURL = cfg.R2PublicURL
}
