package accountsclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

type CreateAccountInput struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DepartmentID int    `json:"department_id"`
}

func (c *Client) CreateAccount(
	ctx context.Context,
	input CreateAccountInput,
) error {
	// marshal-kan input, simpan di  body
	body, err := json.Marshal(input)
	// error handling jika gagal
	if err != nil {
		return err
	}

	// bikin request baru
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/internal/accounts",
		bytes.NewBuffer(body),
	)
	// error handling
	if err != nil {
		return err
	}

	// set header request
	req.Header.Set("Content-Type", "application/json")

	// do request dengan client, simpan di resp
	resp, err := c.httpClient.Do(req)
	//error handling
	if err != nil {
		return fmt.Errorf("calling accounts service: %w", err)
	}
	// defer setelah selesai
	defer resp.Body.Close()

	// cek kode status. jika tidak status created, maka kasih tahu
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("calling accounts service returned: %d", resp.StatusCode)
	}

	return nil
}
