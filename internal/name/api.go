package name

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ipApi         string = "https://api.myip.com"
	records       string = "https://api.name.com/v4/domains/%s/records"
	updateRecords string = "https://api.name.com/v4/domains/%s/records/%d"
)

type api struct{}

func newApi() *api {
	return &api{}
}

type IpApiResponse struct {
	Ip string `json:"ip"`
}

// TODO: probably move this into its own package
func (a *api) getIp(task Task) (string, error) {
	resp, err := http.Get(ipApi)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var data IpApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Ip == "" {
		return "", fmt.Errorf("Ip field is empty")
	}

	return data.Ip, nil
}

func (a *api) update(task Task, ip string) error {

	// Get record id
	record, err := a.getRecord(task)

	if err != nil {
		return err
	}

	if record == nil {
		log.Printf("No record found. Creating a new record %s.%s", task.Host, task.Domain)
		err = a.createRecord(task, ip)
	} else {
		log.Printf("Updating record %s.%s", record.Host, task.Domain)
		err = a.updateRecord(task, record, ip)
	}

	return err
}

type RecordsResponse struct {
	Records []Record `json:"records"`
}

type Record struct {
	Id     int64  `json:"id"`
	Type   string `json:"type"`
	Host   string `json:"host"`
	Answer string `json:"answer,omitempty"`
}

func (a *api) getRecord(task Task) (*Record, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf(records, task.Domain), nil)
	req.SetBasicAuth(task.User, task.Token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("Unexpected HTTP Response Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	var data RecordsResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	for _, record := range data.Records {
		if record.Type == "A" && record.Host == task.Host {
			return &record, nil
		}
	}

	return nil, nil
}

func (a *api) createRecord(task Task, ip string) error {
	payload, _ := json.Marshal(map[string]interface{}{"host": task.Host, "type": "A", "answer": ip, "ttl": 300})

	req, _ := http.NewRequest("POST", fmt.Sprintf(records, task.Domain), bytes.NewBuffer(payload))
	req.SetBasicAuth(task.User, task.Token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Unexpected HTTP Response Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return nil
}

func (a *api) updateRecord(task Task, record *Record, ip string) error {
	if record.Answer == ip {
		return nil
	}

	payload, _ := json.Marshal(map[string]interface{}{"host": record.Host, "type": record.Type, "answer": ip, "ttl": 300})

	req, _ := http.NewRequest("PUT", fmt.Sprintf(updateRecords, task.Domain, record.Id), bytes.NewBuffer(payload))
	req.SetBasicAuth(task.User, task.Token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Unexpected HTTP Response Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return nil
}
