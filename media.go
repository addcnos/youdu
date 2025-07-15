package youdu

import (
	"context"
	"io"
	"net/http"
)

type FileType string

const (
	FileTypeImage FileType = "image"
	FileTypeFile  FileType = "file"
	FileTypeVoice FileType = "voice"
	FileTypeVideo FileType = "video"
)

type UploadMediaRequest struct {
	File     io.Reader `json:"-"`
	FileName string    `json:"name"`
	FileType FileType  `json:"fileType"`
}

type UploadMediaResponse struct {
	MediaID string `json:"mediaId"`
}

type GetMediaRequest struct {
	MediaID string `json:"mediaId"`
}

type GetMediaResponse struct {
	Name string `json:"name"`
	Size int32  `json:"size"`
	File []byte `json:"file"`
}

type SearchMediaRequest struct {
	MediaID string `json:"mediaId"`
}

type SearchMediaResponse struct {
	Name string `json:"name"`
	Size int32  `json:"size"`
}

func (c *Client) GetMedia(
	ctx context.Context,
	request GetMediaRequest,
) (response GetMediaResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/media/get",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseBodyDecrypt())
	return
}

func (c *Client) UploadMedia(
	ctx context.Context,
	req UploadMediaRequest,
) (response UploadMediaResponse, err error) {
	request, err := c.newRequest(ctx, http.MethodPost, "/cgi/media/upload",
		withRequestBody(req), withRequestAccessToken(), withRequestEncrypt(), withRequestType(uploadRequestType))
	if err != nil {
		return
	}
	err = c.sendRequest(request, &response, withResponseDecrypt())
	return
}

func (c *Client) SearchMedia(
	ctx context.Context,
	request SearchMediaRequest,
) (response SearchMediaResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cgi/media/search",
		withRequestBody(request), withRequestAccessToken(), withRequestEncrypt())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response, withResponseDecrypt())
	return
}
