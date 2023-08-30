package oss

import (
	"context"
	"fmt"
	"io"

	alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// UploadParams 上传文件的配置
type UploadParams struct {
	FileName string
	Stream   io.Reader
	Bucket   string
}

// Upload 返回文件在oss的url地址
// 上传文件的模式为公共读
func Upload(ctx context.Context, cfg AliOssConfig, uploadParams UploadParams) error {

	var (
		client *alioss.Client
		bucket *alioss.Bucket
		err    error
	)

	client, err = createAliOssClient(cfg)
	if err != nil {
		return err
	}

	bucket, err = client.Bucket(uploadParams.Bucket)
	if err != nil {
		return err
	}

	err = bucket.PutObject(uploadParams.FileName, uploadParams.Stream, alioss.ObjectACL(alioss.ACLPublicRead))
	if err != nil {
		return fmt.Errorf("upload file to oss failed, err: %w", err)
	}

	return nil
}
