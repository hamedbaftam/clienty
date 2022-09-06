package clienty

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// DefaultBaseURL is the default API base url used by Client to send requests to Trello.
const DefaultBaseURL = "https://api.trello.com/1"

// Client is the central object for making API calls. It wraps a http client,
// context, logger and identity configuration (Key and Token) of the Trello member.
type Client struct {
	Client *http.Client
	//Logger   logger
	BaseURL string
	Key     string
	Token   string
	//throttle *rate.Limiter
	testMode bool
	ctx      context.Context
}

// Get takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a GET request on the Trello API endpoint and the path and uses the
// Arguments as URL parameters. Then it returns either the target interface
// updated from the response or an error.
func (c *Client) Get(path string, args Arguments, target interface{}) error {

	// Trello prohibits more than 10 seconds/second per token
	//c.Throttle()

	params := args.ToURLValues()
	//c.log("[trello] GET %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {

	}
	req = req.WithContext(c.ctx)

	return c.do(req, url, target)
}

// Put takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a PUT request on the Trello API endpoint with the path and uses
// the Arguments as URL parameters. Then it returns either the target interface
// updated from the response or an error.
func (c *Client) Put(path string, args Arguments, target interface{}) error {

	// Trello prohibits more than 10 seconds/second per token
	//c.Throttle()

	params := args.ToURLValues()
	//c.log("[trello] PUT %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("PUT", urlWithParams, nil)
	if err != nil {

	}

	return c.do(req, url, target)
}

// Post takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a POST request on the Trello API endpoint with the path and uses
// the Arguments as URL parameters. Then it returns either the target interface
// updated from the response or an error.
func (c *Client) Post(path string, args Arguments, target interface{}) error {

	// Trello prohibits more than 10 seconds/second per token
	//c.Throttle()

	params := args.ToURLValues()
	//c.log("[trello] POST %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("POST", urlWithParams, nil)
	if err != nil {
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.do(req, url, target)
}

// PostWithBody takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a POST request on the Trello API endpoint with the path and uses
// the Arguments as URL parameters, takes file io.Reader and put to multipart body.
// Then it returns either the target interface
// updated from the response or an error.
func (c *Client) PostWithBody(path string, args Arguments, target interface{}, filename string, file io.Reader) error {

	// Trello prohibits more than 10 seconds/second per token

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	params := args.ToURLValues()
	//c.log("[trello] POST %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("POST", urlWithParams, body)
	if err != nil {

	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return c.do(req, url, target)
}

// Delete takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a DELETE request on the Trello API endpoint with the path and uses
// the Arguments as URL parameters. Then it returns either the target interface
// updated from the response or an error.
func (c *Client) Delete(path string, args Arguments, target interface{}) error {

	//c.Throttle()

	params := args.ToURLValues()
	//c.log("[trello] DELETE %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("DELETE", urlWithParams, nil)
	if err != nil {
		//return errors.Wrapf(err, "Invalid DELETE request %s", url)
	}

	return c.do(req, url, target)
}

func (c *Client) do(req *http.Request, url string, target interface{}) error {
	resp, err := c.Client.Do(req)
	if err != nil {
		//return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		//return makeHTTPClientError(url, resp)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//return errors.Wrapf(err, "HTTP Read error on response for %s", url)
	}
	err = json.Unmarshal(b, target)
	if err != nil {
		//return errors.Wrapf(err, "JSON decode failed on %s:\n%s", url, string(b))
	}
	return nil
}
