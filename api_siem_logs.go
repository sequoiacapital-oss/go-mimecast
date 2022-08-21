package mimecast

import "context"

type GetSiemLogsApi struct {
	token string
}

func (r *GetSiemLogsApi) Url() string {
	return "https://us-api.mimecast.com/api/audit/get-siem-logs"
}

func (r *GetSiemLogsApi) RequestData() M {
	req := M{"type": "MTA", "fileFormat": "json"}
	if r.token != "" {
		req["token"] = r.token
	}

	return req
}

func (r *GetSiemLogsApi) RequestMeta() M {
	return nil
}

func GetSiemLogs(ctx context.Context, token string) (*Response, error) {
	c := &GetSiemLogsApi{token: token}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
