package mimecast

import (
	"context"
	"fmt"
	"net/url"
)

type CreateManagedUrlApi struct {
	urlToManage string
	action      string
	comment     string
}

func (r *CreateManagedUrlApi) Url() string {
	return "https://us-api.mimecast.com/api/ttp/url/create-managed-url"
}

func (r *CreateManagedUrlApi) RequestData() M {
	data := M{"url": r.urlToManage, "action": r.action, "comment": r.comment}

	myurl, _ := url.Parse(r.urlToManage) // TODO: HANDLE ERROR

	if myurl.Path == "" || myurl.Path == "/" {
		data["matchType"] = "domain"
	} else {
		data["matchType"] = "explicit"
	}

	return data
}

type ManagedUrlError struct {
	Message string
	Action  string
}

func (e ManagedUrlError) Error() string {
	return fmt.Sprintf("Error doing %s: %s", e.Action, e.Message)
}

func (r *CreateManagedUrlApi) RequestMeta() M {
	return nil
}

func PermitUrl(ctx context.Context, urlToPermit string, comment string) (string, error) {
	return createManagedUrl(ctx, urlToPermit, comment, "permit")
}

func BlockUrl(ctx context.Context, urlToBlock string, comment string) (string, error) {
	return createManagedUrl(ctx, urlToBlock, comment, "block")
}

func createManagedUrl(ctx context.Context, urlToManage string, comment string, action string) (string, error) {
	c := &CreateManagedUrlApi{urlToManage: urlToManage, comment: comment, action: action}

	r := &Request{api: c}
	req := r.buildRequest()

	_, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return "", ManagedUrlError{Message: err.Error(), Action: action}
	}

	var blockType = "domain"

	if c.RequestData()["matchType"] == "explicit" {
		blockType = "url"
	}

	return "ok, " + action + "ed that " + blockType, nil
}
