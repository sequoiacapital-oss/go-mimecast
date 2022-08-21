package mimecast

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type GetHoldQueueApi struct {
	recipient string
}

func (r *GetHoldQueueApi) Url() string {
	return "https://us-api.mimecast.com/api/gateway/get-hold-message-list"
}

func (r *GetHoldQueueApi) RequestData() M {
	req := M{"admin": true, "searchBy": map[string]string{"fieldName": "recipient", "value": r.recipient}}
	return req
}

func (r *GetHoldQueueApi) RequestMeta() M {
	return nil
}

type HoldQueueEntryFromHeader struct {
	EmailAddress string
}

type HoldQueueEntry struct {
	Id           string
	ReasonId     string
	Subject      string
	DateReceived string
	FromHeader   HoldQueueEntryFromHeader `json:fromHeader`
}

func GetHoldQueueForRecipient(ctx context.Context, recipient string) ([]*HoldQueueEntry, error) {
	c := &GetHoldQueueApi{recipient: recipient}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return nil, err
	}

	entries := []*HoldQueueEntry{}

	for _, rawEntry := range respData.Data {
		entry := &HoldQueueEntry{}
		mapstructure.Decode(rawEntry, entry)
		fmt.Printf("ENTRY %v\n", entry)
		entries = append(entries, entry)
	}

	return entries, nil
}

type ReleaseFromHoldQueueApi struct {
	id string
}

func (r *ReleaseFromHoldQueueApi) Url() string {
	return "https://us-api.mimecast.com/api/gateway/hold-release"
}

func (r *ReleaseFromHoldQueueApi) RequestData() M {
	req := M{"id": r.id}
	return req
}

func (r *ReleaseFromHoldQueueApi) RequestMeta() M {
	return nil
}

func ReleaseFromHoldQueue(ctx context.Context, id string) (bool, error) {
	c := &ReleaseFromHoldQueueApi{id: id}

	r := &Request{api: c}
	req := r.buildRequest()

	respData, err := DoRequestWithContext(ctx, req)

	if err != nil {
		return false, err
	}

	if len(respData.Data) == 1 {
		data := respData.Data[0]

		released := data["release"].(bool)
		return released, nil
	}

	return false, fmt.Errorf("Error releasing from queue")
}
