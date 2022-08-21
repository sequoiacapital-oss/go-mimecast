package mimecast

import "context"

type GetArchiveViewLogsApi struct {
	start     string
	end       string
	pageToken string
}

type GetArchiveSearchLogsApi struct {
	start     string
	end       string
	pageToken string
}

func (r *GetArchiveViewLogsApi) Url() string {
	return "https://us-api.mimecast.com/api/archive/get-view-logs"
}

func (r *GetArchiveViewLogsApi) RequestData() M {
	req := M{"start": r.start, "end": r.end}
	return req
}

func (r *GetArchiveViewLogsApi) RequestMeta() M {
	if r.pageToken != "" {
		return M{"pagination": M{"pageSize": 10, "pageToken": r.pageToken}}
	}
	return nil
}

func (r *GetArchiveSearchLogsApi) Url() string {
	return "https://us-api.mimecast.com/api/archive/get-search-logs"
}

func (r *GetArchiveSearchLogsApi) RequestData() M {
	req := M{"start": r.start, "end": r.end}
	return req
}

func (r *GetArchiveSearchLogsApi) RequestMeta() M {
	if r.pageToken != "" {
		return M{"pagination": M{"pageSize": 10, "pageToken": r.pageToken}}
	}
	return nil
}

func GetArchiveViewLogs(ctx context.Context, fromTime string, endTime string, pageToken string) (*Response, error) {
	c := &GetArchiveViewLogsApi{start: fromTime, end: endTime, pageToken: pageToken}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	return respData, nil
}

func GetArchiveSearchLogs(ctx context.Context, fromTime string, endTime string, pageToken string) (*Response, error) {
	c := &GetArchiveSearchLogsApi{start: fromTime, end: endTime, pageToken: pageToken}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
