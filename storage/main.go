package storage

import (
	"log"

	minio "github.com/minio/minio-go"
)

// Setup storage client
func Setup(config Configer) (Clienter, error) {
	minioClient, err := minio.New(config.Endpoint(), config.AccessKey(), config.SecretKey(), config.Secure())
	if err != nil {
		log.Println(err)
		return nil, ErrorConnection
	}
	return &client{minioClient, config}, nil
}
