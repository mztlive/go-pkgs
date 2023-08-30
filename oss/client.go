package oss

import alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"

func createAliOssClient(cfg AliOssConfig) (*alioss.Client, error) {
	client, err := alioss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	return client, nil
}
