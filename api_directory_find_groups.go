package mimecast

import (
	"context"
	"fmt"
)

type FindGroupsApi struct {
	name string
}

func (r *FindGroupsApi) Url() string {
	return "https://us-api.mimecast.com/api/directory/find-groups"
}

func (r *FindGroupsApi) RequestData() M {
	return M{"query": r.name, "source": "cloud"} // can also be "ldap"
}

func (r *FindGroupsApi) RequestMeta() M {
	return nil
}

func FindGroups(ctx context.Context, name string) string {
	c := &FindGroupsApi{name: name}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		msg := err.Error()
		fmt.Println("Error doing FindGroups: " + msg)
		return msg
	} else {
		return fmt.Sprintf("FindGroups ResponseData: %v", respData.Data)
	}
}
