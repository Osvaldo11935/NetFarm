package storages

import (
	"NetFarm/application/interfaces/iinfrastructure/istorages"
	"NetFarm/shared/errors"
	"NetFarm/shared/extensions"
	"NetFarm/shared/models/common"
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"mime/multipart"
	"net/url"
	"os"
)

var bucket = extensions.GetEnv("BUCKET_NAME")

var fileKeys = extensions.GetEnv("FILE_CONFIG_GOOGLE_STORAGE")

var baseUrlGoogleStorage = extensions.GetEnv("BASE_URL_GOOGLE_STORAGE")

type GoogleStorage struct {
}

func NewGoogleStorage() istorages.IFileManagerStorage {
	return &GoogleStorage{}
}

func (s *GoogleStorage) Upload(c *gin.Context) (string, *common.ErrorResponse) {

	if _, err := os.Stat(fileKeys); os.IsNotExist(err) {
		return "", errors.NewGoogleStorageFileKeysNotFound()
	}

	ctx := appengine.NewContext(c.Request)

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(fileKeys))

	if err != nil {
		return "", errors.NewGoogleStorageUnknownError(err.Error(), "falha ao fazer upload do ficheiro")
	}

	f, uploadedFile, err := c.Request.FormFile("file")

	if err != nil {
		return "", errors.NewGoogleStorageUnknownError(err.Error(), "falha ao fazer upload do ficheiro")
	}

	defer func(f multipart.File) {
		if err := f.Close(); err != nil {
			errors.NewGoogleStorageUnknownError(err.Error(), "falha ao fazer upload do ficheiro")
		}
	}(f)

	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return "", errors.NewGoogleStorageUnknownError(err.Error(), "falha ao fazer upload do ficheiro")
	}

	if err := sw.Close(); err != nil {
		return "", errors.NewGoogleStorageUnknownError(err.Error(), "falha ao fazer upload do ficheiro")
	}

	u, err := url.Parse(fmt.Sprintf("%s/%s/%s", baseUrlGoogleStorage, bucket, uploadedFile.Filename))

	if err != nil {
		return "", errors.NewGoogleStorageUnknownError(err.Error(), "falha ao fazer upload do ficheiro")
	}
	
	return u.String(), nil
}
