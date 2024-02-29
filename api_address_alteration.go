package mimecast

import (
	"context"
)

type AddressAlterationGetDefinition struct {
	folderId string
}

func (r *AddressAlterationGetDefinition) Url() string {
	return "https://us-api.mimecast.com/api/policy/address-alteration/get-definition"
}

func (r *AddressAlterationGetDefinition) RequestData() M {
	return M{"folderId": r.folderId}
}

func (r *AddressAlterationGetDefinition) RequestMeta() M {
	return nil
}

func GetAddressAlterationDefinitions(ctx context.Context, folderId string) (*Response, error) {
	c := &AddressAlterationGetDefinition{folderId: folderId}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
