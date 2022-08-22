package mimecast

import (
	"encoding/base64"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	SetMimeCastConfig(MimeCastConfig{
		ApplicationId:  "ApplicationId",
		ApplicationKey: "ApplicationKey",
		AccessKey:      "AccessKey",
		SecretKey:      base64.StdEncoding.EncodeToString([]byte("SecretKey")),
	})
}

type TestApi struct {
	requestData *M
	requestMeta *M
}

func (r *TestApi) Url() string {
	return "https://us-api.mimecast.com/api/test/some-url"
}

func (r *TestApi) RequestData() M {
	if r.requestData == nil {
		return nil
	}
	return *r.requestData
}

func (r *TestApi) RequestMeta() M {
	if r.requestMeta == nil {
		return nil
	}
	return *r.requestMeta
}

func TestRequestId(t *testing.T) {
	r := &Request{}
	assert.Empty(t, r.id)

	initialId := r.Id()
	assert.NotEmpty(t, initialId)

	subsequentId := r.Id()
	assert.Equal(t, initialId, subsequentId)
}

func TestRequestDate(t *testing.T) {
	r := &Request{}
	assert.Empty(t, r.date)

	initialDate := r.Date()
	assert.NotEmpty(t, initialDate)

	subsequentDate := r.Date()
	assert.Equal(t, initialDate, subsequentDate)
}

func TestAuthorizationHeader(t *testing.T) {
	r := &Request{id: "3d796b76-4180-40e5-8f7c-6e0fd643b1a9", date: "Sun, 21 Aug 2022 13:49:40 UTC", api: &TestApi{}}

	assert.Equal(t, r.AuthorizationHeader(), "MC AccessKey:WOywLsnZlsE+D6bp/ug6uch5kxc=")
}

func TestRequest(t *testing.T) {
	api := &TestApi{}
	r := &Request{id: "3d796b76-4180-40e5-8f7c-6e0fd643b1a9", date: "Sun, 21 Aug 2022 13:49:40 UTC", api: api}
	req := r.buildRequest()

	assert.Equal(t, req.URL.String(), api.Url())

	assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
	assert.Equal(t, req.Header.Get("x-mc-app-id"), MimeCastGlobalConfig.ApplicationId)
	assert.Equal(t, req.Header.Get("x-mc-req-id"), r.Id())
	assert.Equal(t, req.Header.Get("x-mc-date"), r.Date())
	assert.Equal(t, req.Header.Get("Authorization"), r.AuthorizationHeader())
}

func TestRequestBodies(t *testing.T) {
	var meta M
	var data M

	api := &TestApi{requestMeta: &meta, requestData: &data}
	r := &Request{id: "3d796b76-4180-40e5-8f7c-6e0fd643b1a9", date: "Sun, 21 Aug 2022 13:49:40 EDT", api: api}
	req := r.buildRequest()

	bodyData, _ := io.ReadAll(req.Body)

	assert.Equal(t, `{"data":[]}`, string(bodyData))

	meta = make(M)

	req = r.buildRequest()
	bodyData, _ = io.ReadAll(req.Body)

	assert.Equal(t, `{"data":[],"meta":{}}`, string(bodyData))

	meta["key"] = "value"
	req = r.buildRequest()
	bodyData, _ = io.ReadAll(req.Body)

	assert.Equal(t, `{"data":[],"meta":{"key":"value"}}`, string(bodyData))

	data = make(M)
	data["url"] = "blah"

	req = r.buildRequest()
	bodyData, _ = io.ReadAll(req.Body)

	assert.Equal(t, `{"data":[{"url":"blah"}],"meta":{"key":"value"}}`, string(bodyData))

}
