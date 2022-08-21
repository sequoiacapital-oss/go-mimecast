package mimecast

import (
	"context"
	"fmt"
)

type AddEmailGroupMemberApi struct {
	groupId      string
	emailAddress string
}

type RemoveEmailGroupMemberApi struct {
	groupId      string
	emailAddress string
}

type AddDomainGroupMemberApi struct {
	groupId string
	domain  string
}

func (r *AddEmailGroupMemberApi) Url() string {
	return "https://us-api.mimecast.com/api/directory/add-group-member"
}

func (r *AddEmailGroupMemberApi) RequestData() M {
	return M{"id": r.groupId, "emailAddress": r.emailAddress}
}

func (r *AddEmailGroupMemberApi) RequestMeta() M {
	return nil
}

func (r *RemoveEmailGroupMemberApi) Url() string {
	return "https://us-api.mimecast.com/api/directory/remove-group-member"
}

func (r *RemoveEmailGroupMemberApi) RequestData() M {
	return M{"id": r.groupId, "emailAddress": r.emailAddress}
}

func (r *RemoveEmailGroupMemberApi) RequestMeta() M {
	return nil
}

func (r *AddDomainGroupMemberApi) Url() string {
	return "https://us-api.mimecast.com/api/directory/add-group-member"
}

func (r *AddDomainGroupMemberApi) RequestData() M {
	return M{"id": r.groupId, "domain": r.domain}
}

func (r *AddDomainGroupMemberApi) RequestMeta() M {
	return nil
}

type AddEmailToGroupError struct {
	Message string
}

func (e AddEmailToGroupError) Error() string {
	return fmt.Sprintf("Error doing AddEmailToGroup: %s", e.Message)
}

func AddEmailToGroup(ctx context.Context, groupId string, email string) error {
	c := &AddEmailGroupMemberApi{groupId: groupId, emailAddress: email}

	r := &Request{api: c}
	req := r.buildRequest()

	_, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return AddEmailToGroupError{Message: err.Error()}
	}

	return nil
}

type RemoveEmailFromGroupError struct {
	Message string
}

func (e RemoveEmailFromGroupError) Error() string {
	return fmt.Sprintf("Error doing RemoveEmailFromGroup: %s", e.Message)
}

func RemoveEmailFromGroup(ctx context.Context, groupId string, email string) error {
	c := &RemoveEmailGroupMemberApi{groupId: groupId, emailAddress: email}

	r := &Request{api: c}
	req := r.buildRequest()

	_, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return RemoveEmailFromGroupError{Message: err.Error()}
	}

	return nil
}

type AddDomainToGroupError struct {
	Message string
}

func (e AddDomainToGroupError) Error() string {
	return fmt.Sprintf("Error doing AddDomainToGroup: %s", e.Message)
}

func AddDomainToGroup(ctx context.Context, groupId string, domain string) error {
	c := &AddDomainGroupMemberApi{groupId: groupId, domain: domain}

	r := &Request{api: c}
	req := r.buildRequest()

	_, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return AddDomainToGroupError{Message: err.Error()}
	}

	return nil
}
