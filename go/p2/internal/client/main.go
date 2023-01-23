package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ActivityClient struct {
	URL string
}

type Activity struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

type ActivityDocument struct {
	Activity Activity `json:"activity"`
}

type IDDocument struct {
	ID int `json:"id"`
}

func (a Activity) String() string {
	return fmt.Sprintf("\"%s\" added at %d-%d-%d %d:%d:%d", a.Description,
		a.Time.Year(), a.Time.Month(), a.Time.Day(),
		a.Time.Hour(), a.Time.Minute(), a.Time.Second())
}

func (c *ActivityClient) Insert(a Activity) (int, error) {
	activityDoc := ActivityDocument{Activity: a}
	activityDocBytes, err := json.Marshal(activityDoc)

	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewReader(activityDocBytes))

	if err != nil {
		return -1, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return -1, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return -1, fmt.Errorf("response body: %v", string(body))
	}

	var id IDDocument
	json.Unmarshal(body, &id)

	if err != nil {
		return -1, err
	}

	return id.ID, nil
}

func (c *ActivityClient) Retrieve(id int) (*Activity, error) {

	idStr := strconv.Itoa(id)
	url := c.URL + "/" + idStr
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if res.Body != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response error body: %v", body)
	}

	var a Activity

	err = json.Unmarshal(body, &a)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (c *ActivityClient) List() (*[]Activity, error) {
	url := c.URL
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if res.Body != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response error body: %v", body)
	}

	var a []Activity

	err = json.Unmarshal(body, &a)

	if err != nil {
		return nil, err
	}

	return &a, nil
}
