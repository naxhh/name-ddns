package name

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	ipApi         string = "https://api.myip.com"
	records              = "https://api.name.com/v4/domains/%s/records"
	updateRecords        = "https://api.name.com/v4/domains/%s/records/%d"
)

type api struct {
	config *Config
}

func newApi(config *Config) *api {
	return &api{
		config: config,
	}
}

type IpApiResponse struct {
	Ip string `json:"ip"`
}

func (a *api) getIp() (string, error) {
	resp, err := http.Get(ipApi)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	var data IpApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Ip == "" {
		return "", errors.New("Ip field is empty")
	}

	return data.Ip, nil
}

func (a *api) update(ip string) error {

	// Get record id
	record, err := a.getRecord()

	if err != nil {
		return err
	}

	if record == nil {
		log.Println("No record found. Creating a new record")
		err = a.createRecord(ip)
	} else {
		err = a.updateRecord(record, ip)
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

func (a *api) getRecord() (*Record, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf(records, a.config.Domain), nil)
	req.SetBasicAuth(a.config.User, a.config.Token)

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected HTTP Response Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
	}

	var data RecordsResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	for _, record := range data.Records {
		if record.Type == "A" && record.Host == a.config.Host {
			return &record, nil
		}
	}

	return nil, nil
}

func (a *api) createRecord(ip string) error {
	payload, _ := json.Marshal(map[string]interface{}{"host": a.config.Host, "type": "A", "answer": ip, "ttl": 300})

	req, _ := http.NewRequest("POST", fmt.Sprintf(records, a.config.Domain), bytes.NewBuffer(payload))
	req.SetBasicAuth(a.config.User, a.config.Token)

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Unexpected HTTP Response Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
	}

	return nil
}

func (a *api) updateRecord(record *Record, ip string) error {
	if record.Answer == ip {
		return nil
	}

	payload, _ := json.Marshal(map[string]interface{}{"host": record.Host, "type": record.Type, "answer": ip, "ttl": 300})

	req, _ := http.NewRequest("PUT", fmt.Sprintf(updateRecords, a.config.Domain, record.Id), bytes.NewBuffer(payload))
	req.SetBasicAuth(a.config.User, a.config.Token)

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Unexpected HTTP Response Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
	}

	return nil
}
