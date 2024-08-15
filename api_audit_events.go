package mimecast

import (
	"context"
)

type GetAuditEventsApi struct {
	pageToken string
	from      string
	to        string
}

func (r *GetAuditEventsApi) Url() string {
	return "https://us-api.mimecast.com/api/audit/get-audit-events"
}

func (r *GetAuditEventsApi) RequestData() M {
	req := M{"startDateTime": r.from, "endDateTime": r.to}
	return req
}

func (r *GetAuditEventsApi) RequestMeta() M {
	if r.pageToken != "" {
		return M{"pagination": M{"pageSize": 10, "pageToken": r.pageToken}}
	}
	return nil
}

func GetAuditEvents(ctx context.Context, from string, to string, pageToken string) (*Response, error) {
	c := &GetAuditEventsApi{to: to, from: from, pageToken: pageToken}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
