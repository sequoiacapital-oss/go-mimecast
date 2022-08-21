package mimecast

import (
	"context"
	"time"

	"github.com/mitchellh/mapstructure"
)

/*
func FindMessage(name string) string {
	url := "https://us-api.mimecast.com/api/message-finder/search"
	data := M{"query" : name, "source": "cloud"}

	respData, err := DoRequest(url, data)
	if err != nil {
		msg := err.Error()
		fmt.Println("Error doing FindGroups: " + msg)
		return msg
	} else {
		return fmt.Sprintf("FindGroups ResponseData: %v", respData)
	}
}
*/

type MessageFinderResponse struct {
	TrackedEmails []struct {
		Status   string `json:"status"`
		Received string `json:"received"`
		FromEnv  struct {
			DisplayableName string `json:"displayableName"`
			EmailAddress    string `json:"emailAddress"`
		} `json:"fromEnv"`
		FromHdr struct {
			DisplayableName string `json:"displayableName"`
			EmailAddress    string `json:"emailAddress"`
		} `json:"fromHdr"`
		Attachments bool `json:"attachments"`
		To          []struct {
			DisplayableName string `json:"displayableName"`
			EmailAddress    string `json:"emailAddress"`
		} `json:"to"`
		SenderIP string `json:"senderIP"`
		Route    string `json:"route"`
		ID       string `json:"id"`
		Sent     string `json:"sent"`
		Subject  string `json:"subject"`
	} `json:"trackedEmails"`
}

type MessageSearchApi struct {
	from      string
	to        string
	startTime string
	endTime   string
	reason    string
}

func (r *MessageSearchApi) Url() string {
	return "https://us-api.mimecast.com/api/message-finder/search"
}

func (r *MessageSearchApi) RequestData() M {
	advancedOptions := map[string]string{
		"from": r.from,
	}

	if len(r.to) > 0 {
		advancedOptions["to"] = r.to
	}

	data := map[string]interface{}{
		"start":                        r.startTime,
		"end":                          r.endTime,
		"searchReason":                 r.reason,
		"advancedTrackAndTraceOptions": advancedOptions,
	}

	return data
}

func (r *MessageSearchApi) RequestMeta() M {
	return nil
}

func FindMessages(ctx context.Context, from string, to string, timeStart time.Time, reason string) MessageFinderResponse {
	timeEnd := time.Now()
	startTime := timeStart.Format("2006-01-02T15:04:05-0700")
	endTime := timeEnd.Format("2006-01-02T15:04:05-0700")

	c := &MessageSearchApi{from: from, to: to, startTime: startTime, endTime: endTime, reason: reason}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, _ := DoRequestWithContext(ctx, req)

	var respStruct MessageFinderResponse
	mapstructure.Decode(respData.Data, &respStruct)
	return respStruct
}
