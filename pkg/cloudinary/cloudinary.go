package cloudx

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

var cloudinaryURL = "cloudinary://875353593587434:CljfYcEUdaFat0UUAVJ6nlrqNQQ@dhg9m6lhd"

func UploadImage(file multipart.File) (string, error) {

	uploadParams := uploader.UploadParams{
		Folder: "uploads",
	}

	ctx := context.Background()
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", fmt.Errorf("failed to create cloudinary connection: %v", err)
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to Cloudinary: %v", err)
	}
	fmt.Println(uploadResult)
	return uploadResult.SecureURL, nil
}
