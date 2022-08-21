package mimecast

import "context"

type GetInternalUsersApi struct {
	domain    string
	pageToken string
}

func (r *GetInternalUsersApi) Url() string {
	return "https://us-api.mimecast.com/api/user/get-internal-users"
}

func (r *GetInternalUsersApi) RequestData() M {
	req := M{"domain": r.domain}

	return req
}

func (r *GetInternalUsersApi) RequestMeta() M {
	if r.pageToken != "" {
		return M{"pagination": M{"pageSize": 100, "pageToken": r.pageToken}}
	}
	return nil
}

func GetInternalUsers(ctx context.Context, domain string, pageToken string) (*Response, error) {
	c := &GetInternalUsersApi{domain: domain, pageToken: pageToken}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	return respData, nil
}

type UserInfo struct {
	Name         string
	EmailAddress string
	Alias        bool
	AddressType  string
	Source       string
}

type InternalUsers struct {
	Users []UserInfo
}
