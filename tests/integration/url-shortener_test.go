package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const uri = "http://localhost:8080"

var code *string

func TestMakeShortUrl(t *testing.T) {
	t.Run("success_create_item", func(t *testing.T) {
		data := &struct {
			OriginalUrl string `json:"original_url"`
		}{
			OriginalUrl: "https://www.google.com",
		}

		jsonData, _ := json.Marshal(data)

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/url", uri), bytes.NewBuffer(jsonData))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		fmt.Println(string(body))
		respBody := &struct {
			Result struct {
				Ok bool `json:"ok"`
			} `json:"result"`
		}{}

		err = json.Unmarshal(body, respBody)
		assert.NoError(t, err)

		assert.Equal(t, true, respBody.Result.Ok)
	})
	t.Run("failed_marshal_url", func(t *testing.T) {
		data := &struct {
			OriginalUrl int `json:"original_url"`
		}{
			OriginalUrl: 100000,
		}

		jsonData, _ := json.Marshal(data)

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/url", uri), bytes.NewBuffer(jsonData))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestGetUrlList(t *testing.T) {
	t.Run("success_get_url_list", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/url/list", uri), nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		respBody := &struct {
			Result struct {
				Urls []struct {
					Id          int    `json:"id"`
					OriginalUrl string `json:"original_url"`
					Code        string `json:"code"`
				} `json:"urls"`
			} `json:"result"`
		}{}

		err = json.Unmarshal(body, respBody)
		assert.NoError(t, err)
		code = &respBody.Result.Urls[0].Code

		assert.Equal(t, "https://www.google.com", respBody.Result.Urls[0].OriginalUrl)
	})
	t.Run("success_get_url", func(t *testing.T) {
		data := &struct {
			Code string `json:"code"`
		}{
			Code: *code,
		}

		jsonData, _ := json.Marshal(data)

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/url", uri), bytes.NewBuffer(jsonData))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		respBody := &struct {
			Result struct {
				OriginalUrl string `json:"original_url"`
			} `json:"result"`
		}{}

		err = json.Unmarshal(body, respBody)
		assert.NoError(t, err)

		assert.Equal(t, "https://www.google.com", respBody.Result.OriginalUrl)
	})
	t.Run("failed_marshal_get_url", func(t *testing.T) {
		data := &struct {
			Code int `json:"code"`
		}{
			Code: 100000,
		}

		jsonData, _ := json.Marshal(data)

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/url", uri), bytes.NewBuffer(jsonData))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, 500, resp.StatusCode)
	})
}
