package storage

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"path"
	"strings"
	"time"

	minio "github.com/minio/minio-go"
	uuid "github.com/satori/go.uuid"
)

// Clienter client interface
type Clienter interface {
	Fetch(objectName string) (*minio.Object, error)
	Save(f io.Reader, fSize int64, filename string, contentType string) (Filer, error)
	SaveMultipart(fh *multipart.FileHeader) (Filer, error)
	Delete(objectName string) error
	PublicURL(objectName string) string
}

// Client storage client struct
type client struct {
	minio  *minio.Client
	config Configer
}

// Fetch file
func (c *client) Fetch(objectName string) (*minio.Object, error) {
	object, err := c.minio.GetObject(c.config.Bucket(), objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return nil, ErrorFileFetching
	}
	return object, nil
}

// PresignedURL returns temporary public file URL
func (c *client) PresignedURL(objectName string) (string, error) {
	u, err := c.minio.PresignedGetObject(c.config.Bucket(), objectName, time.Duration(7*24*time.Hour), url.Values{})
	if err != nil {
		log.Println(err)
		return "", ErrorURLFetching
	}
	return u.RequestURI(), nil
}

// Save file and returns object name
func (c *client) SaveMultipart(fh *multipart.FileHeader) (Filer, error) {
	f, err := fh.Open()
	if err != nil {
		return nil, err
	}
	f.Close()
	return c.Save(f, fh.Size, fh.Filename, fh.Header.Get("Content-Type"))
}

// Save file and returns object name
func (c *client) Save(f io.Reader, fSize int64, filename string, contentType string) (Filer, error) {
	obj := &File{}
	obj.originName = filename
	obj.name = fmt.Sprintf("%s%s", uuid.Must(uuid.NewV1()).String(), path.Ext(filename))
	obj.path = c.objectPath(obj.name)
	obj.size = fSize
	obj.mime = contentType
	opts := minio.PutObjectOptions{
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
		ContentType:  obj.mime,
	}
	n, err := c.minio.PutObject(c.config.Bucket(), obj.path, f, obj.size, opts)
	if err != nil {
		log.Println(err)
		return nil, ErrorFileSaving
	}
	if n != obj.size {
		return nil, ErrorCouldNotUpload
	}
	return obj, nil
}

// Delete file
func (c *client) Delete(objectName string) error {
	if err := c.minio.RemoveObject(c.config.Bucket(), objectName); err != nil {
		log.Println(err)
		return ErrorFileDeleting
	}
	return nil
}

// PublicURL returns public URL, compatible to DigitalOcean
func (c *client) PublicURL(objectName string) string {
	return fmt.Sprintf("https://%s.%s/%s", c.config.Bucket(), c.config.Endpoint(), objectName)
}

func (c *client) objectPath(filename string) string {
	folder := strings.Trim(time.Now().Format("2006/01/02"), "/")
	return fmt.Sprintf("%s/%s/%s", c.config.BaseDir(), folder, strings.Trim(filename, "/"))
}
