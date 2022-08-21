package mimecast

import "context"

type GetTTPUrlLogsApi struct {
	from      string
	to        string
	pageToken string
}

type GetTTPAttachmentLogsApi struct {
	from      string
	to        string
	pageToken string
}

func (r *GetTTPUrlLogsApi) Url() string {
	return "https://us-api.mimecast.com/api/ttp/url/get-logs"
}

func (r *GetTTPUrlLogsApi) RequestData() M {
	req := M{"from": r.from, "to": r.to, "oldestFirst": true}
	return req
}

func (r *GetTTPUrlLogsApi) RequestMeta() M {
	if r.pageToken != "" {
		return M{"pagination": M{"pageSize": 10, "pageToken": r.pageToken}}
	}
	return nil
}

func (r *GetTTPAttachmentLogsApi) Url() string {
	return "https://us-api.mimecast.com/api/ttp/attachment/get-logs"
}

func (r *GetTTPAttachmentLogsApi) RequestData() M {
	req := M{"from": r.from, "to": r.to, "oldestFirst": true}
	return req
}

func (r *GetTTPAttachmentLogsApi) RequestMeta() M {
	if r.pageToken != "" {
		return M{"pagination": M{"pageSize": 10, "pageToken": r.pageToken}}
	}
	return nil
}

func GetTTPUrlLogs(ctx context.Context, fromTime string, endTime string, pageToken string) (*Response, error) {
	c := &GetTTPUrlLogsApi{from: fromTime, to: endTime, pageToken: pageToken}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	return respData, nil
}

func GetTTPAttachmentLogs(ctx context.Context, fromTime string, endTime string, pageToken string) (*Response, error) {
	c := &GetTTPAttachmentLogsApi{from: fromTime, to: endTime, pageToken: pageToken}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
