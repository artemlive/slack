package slack

import (
	"encoding/json"
	"errors"
	"github.com/artemlive/slack/slackevents"
)

type ExternalRequest struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	CallbackID string `json:"callback_id"`
	Type       string `json:"type"`
	Team       struct {
		ID     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"team"`
	Channel struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	ActionTs     string `json:"action_ts"`
	MessageTs    string `json:"message_ts"`
	AttachmentID string `json:"attachment_id"`
	Token        string `json:"token"`
}

//select structure
type ExternalSelectResponse struct {
	Options []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"options"`
}

func ParseExternalContent(body string, opts ...slackevents.Option) (ExternalRequest, error) {
	byteString := []byte(body)
	action := ExternalRequest{}
	err := json.Unmarshal(byteString, &action)
	if err != nil {
		return ExternalRequest{}, errors.New("ExternalRequest unmarshalling failed")
	}
	cfg := &slackevents.Config{}
	cfg.VerificationToken = action.Token
	for _, opt := range opts {
		opt(cfg)
	}
	if !cfg.TokenVerified {
		return ExternalRequest{}, errors.New("invalid verification token")
	} else {
		return action, nil
	}
}