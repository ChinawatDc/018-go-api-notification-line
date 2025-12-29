package lineoa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	accessToken string
	httpClient  *http.Client
	baseURL     string
}

type ClientOptions struct {
	AccessToken string
}

func NewClient(opt ClientOptions) *Client {
	return &Client{
		accessToken: opt.AccessToken,
		httpClient:  &http.Client{Timeout: 15 * time.Second},
		baseURL:     "https://api.line.me",
	}
}

func (c *Client) ReplyText(ctx context.Context, replyToken string, text string) error {
	reqBody := ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   []TextMessage{{Type: "text", Text: text}},
	}
	return c.postJSON(ctx, "/v2/bot/message/reply", reqBody)
}

func (c *Client) PushText(ctx context.Context, to string, text string) error {
	reqBody := PushMessageRequest{
		To:       to,
		Messages: []TextMessage{{Type: "text", Text: text}},
	}
	return c.postJSON(ctx, "/v2/bot/message/push", reqBody)
}

func (c *Client) postJSON(ctx context.Context, path string, body any) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+path, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rb, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("LINE API error %d: %s", resp.StatusCode, string(rb))
	}
	return nil
}
