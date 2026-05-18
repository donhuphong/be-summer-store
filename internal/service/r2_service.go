package service

import (
	"be-summer-store/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type PresignResult struct {
	URL       string `json:"url"`
	PublicURL string `json:"public_url"`
}

func GetPresignURL(ctx context.Context, key string) (*PresignResult, error) {
	if key == "" {
		return nil, errors.New("key không được để trống")
	}

	req, err := database.R2Presign.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(database.BucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(10*time.Minute))

	if err != nil {
		return nil, fmt.Errorf("không thể tạo presigned URL: %w", err)
	}

	return &PresignResult{
		URL: req.URL,
		//PublicURL: database.PublicBaseURL + "/" + key,
		PublicURL: "https://pub-b5b91ee94546441e9efcc8f03d17cb48.r2.dev/" + key,
	}, nil
}
