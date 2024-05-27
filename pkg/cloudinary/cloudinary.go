package cloudx

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/pkg/errors"
)

var cloudinaryURL = "cloudinary://875353593587434:CljfYcEUdaFat0UUAVJ6nlrqNQQ@dhg9m6lhd"

func UploadImage(file multipart.File) (string, error) {

	uploadParams := uploader.UploadParams{
		Folder: "uploads",
	}

	ctx := context.Background()
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", errors.New("failed to create cloudinary connection: " + err.Error())
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", errors.New("failed to upload image to Cloudinary: " + err.Error())
	}
	return uploadResult.SecureURL, nil
}
