package youdu

import (
	"context"
	"net/http"
)

type UploadMediaRequest struct {
	File     []byte `json:"file"`
	FileName string `json:"name"`
	FileType string `json:"fileType"`
}

type UploadMediaResponse struct {
	MediaID string `json:"mediaId"`
}

type DownloadMediaRequest struct {
	MediaID string `json:"mediaId"`
}

type SearchMediaRequest struct {
	MediaID string `json:"mediaId"`
}

type SearchMediaResponse struct {
	Name string `json:"name"`
	Size int32  `json:"size"`
}

func (c *Client) DownloadMedia(
	ctx context.Context,
	request DownloadMediaRequest,
) (response []byte, err error) {
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
		withRequestBody(req), withRequestAccessToken(), withRequestEncrypt(), withRequestType(UploadRequestType))
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
