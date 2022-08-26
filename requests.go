package mimecast

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type M map[string]interface{}

type Request struct {
	id   string
	date string
	api  ApiEndpoint
}

// x-mc-req-id header
func (r *Request) Id() string {
	if r.id == "" {
		r.id = uuid.New().String()
	}
	return r.id
}

// x-mc-date header
// format: Tue, 24 Nov 2015 12:50:11 UTC
func (r *Request) Date() string {
	if r.date == "" {
		r.date = time.Now().UTC().Format(time.RFC1123)
	}

	return r.date
}

func (r *Request) AuthorizationHeader() string {
	url, _ := url.Parse(r.api.Url())

	dataToSign := r.Date() + ":" + r.Id() + ":" + url.RequestURI() + ":" + MimeCastGlobalConfig.ApplicationKey

	decodedSecretKey, err := base64.StdEncoding.DecodeString(MimeCastGlobalConfig.SecretKey)

	if err != nil {
		fmt.Println("ERROR ERROR")
	}

	mac := hmac.New(sha1.New, decodedSecretKey)

	mac.Write([]byte(dataToSign))

	header := "MC " + MimeCastGlobalConfig.AccessKey + ":" + base64.StdEncoding.EncodeToString([]byte(mac.Sum(nil)))
	return header
}

func (r *Request) buildRequest() *http.Request {
	var d []M = make([]M, 0)

	if r.api.RequestData() != nil {
		d = append(d, r.api.RequestData())
	}

	params := make(M)
	params["data"] = d

	meta := r.api.RequestMeta()
	if meta != nil {
		params["meta"] = meta
	}

	jsonValue, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", r.api.Url(), bytes.NewBuffer(jsonValue))

	if err != nil {
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-mc-app-id", MimeCastGlobalConfig.ApplicationId)
	req.Header.Add("x-mc-req-id", r.Id())
	req.Header.Add("x-mc-date", r.Date())
	req.Header.Add("Authorization", r.AuthorizationHeader())

	return req
}

func BuildMimecastHttpRequest(a ApiEndpoint) *http.Request {
	r := &Request{api: a}
	return r.buildRequest()
}

func DoRequestWithContext(ctx context.Context, req *http.Request) (*Response, error) {
	_, debug := os.LookupEnv("DEBUG")

	req = req.WithContext(ctx)

	dump2, err := httputil.DumpRequestOut(req, debug)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		fmt.Printf("%q\n\n", dump2)
	}

	resp, err := MimeCastHttpClient.Do(req)

	if err != nil {
		return nil, errors.New("Error performing MimeCast request: " + err.Error())
	}

	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, debug)
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		fmt.Printf("DUMP: %q\n", dump)
	}

	type MimeCastJsonResponse struct {
		Meta struct {
			Status      int  `json:"status"`
			IsLastToken bool `json:"isLastToken,omitempty"`
			Pagination  struct {
				Next string `json:"next,omitempty"`
			} `json:"pagination,omitempty"`
		} `json:"meta"`
		Fail []struct {
			Errors []struct {
				Field     string `json:"field"`
				Code      string `json:"code"`
				Message   string `json:"message"`
				Retryable bool   `json:"retryable"`
			} `json:"errors"`
		} `json:"fail"`
		Data []M `json:"data"`
	}

	respJson := &MimeCastJsonResponse{}
	err = json.NewDecoder(resp.Body).Decode(respJson)

	if err != nil {
		return nil, err
	}

	if len(respJson.Fail) > 0 {
		e := respJson.Fail[0]
		if len(e.Errors) > 0 {
			msg := e.Errors[0].Message

			return nil, errors.New("MimeCast operation error: " + msg)
		}
	}

	nextToken := resp.Header.Get("Mc-Siem-Token")
	if len(nextToken) == 0 {
		nextToken = respJson.Meta.Pagination.Next
	}

	return &Response{
		Status:      resp.StatusCode,
		IsLastToken: respJson.Meta.IsLastToken,
		Data:        respJson.Data,
		MetaData: ResponseMetadata{
			RateLimit:          decodeIntHeader(resp.Header.Get("X-Ratelimit-Limit")),
			RateLimitReset:     decodeIntHeader(resp.Header.Get("X-Ratelimit-Reset")),
			RateLimitRemaining: decodeIntHeader(resp.Header.Get("X-Ratelimit-Remaining")),
			NextToken:          nextToken,
		},
	}, nil
}

func decodeIntHeader(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
