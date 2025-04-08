package security

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/nfnt/resize"
)

func ResizeImage(input []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}

	newImg := resize.Resize(config.Config.Photo.Small, config.Config.Photo.Small, img, resize.Lanczos3)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, newImg, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
