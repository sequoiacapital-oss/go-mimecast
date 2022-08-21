package mimecast

import (
	"context"
)

type DirectorySyncApi struct {
}

func (r *DirectorySyncApi) Url() string {
	return "https://us-api.mimecast.com/api/directory/execute-sync"
}

func (r *DirectorySyncApi) RequestData() M {
	return nil
}

func (r *DirectorySyncApi) RequestMeta() M {
	return nil
}

func DirectorySync(ctx context.Context) string {
	c := &DirectorySyncApi{}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return err.Error()
	}

	if len(respData.Data) == 1 {
		data := respData.Data[0]

		return data["syncStatus"].(string)
	}

	return "status: unknown"
}
