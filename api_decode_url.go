package mimecast

import (
	"context"
	"fmt"
)

type MimeCastResponseI interface {
	TextResponse() string
}

type DecodeUrlApi struct {
	urlToDecode string
}

type DecodeUrlResponse struct {
	url     string `json:"url"`
	success bool   `json:"success"`
}

func (r *DecodeUrlApi) Url() string {
	return "https://us-api.mimecast.com/api/ttp/url/decode-url"
}

func (r *DecodeUrlApi) RequestData() M {
	return M{"url": r.urlToDecode}
}

func (r *DecodeUrlApi) RequestMeta() M {
	return nil
}

func (r *DecodeUrlApi) DecodeResponse(resp []M) MimeCastResponseI {
	return &DecodeUrlResponse{url: resp[0]["url"].(string), success: resp[0]["success"].(bool)}
}

func (r *DecodeUrlResponse) TextResponse() string {
	if r.success != true {
		return "Unsucessful in decoding url"
	}

	return fmt.Sprintf("Decoded url: %v\n", r.url)
}

func DecodeUrl(ctx context.Context, urlToDecode string) string {
	c := &DecodeUrlApi{urlToDecode: urlToDecode}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		msg := err.Error()
		fmt.Println("Error1 doing DecodeUrl: " + msg)
		return msg
	}

	r2 := c.DecodeResponse(respData.Data)
	return r2.TextResponse()
}
