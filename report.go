package nexmo

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Report represents the Report API functions for sending text messages.
// https://developer.vonage.com/en/api/reports
type Report struct {
	client *Client
}

// Product types.
const (
	ProductSMS                 = "SMS"
	ProductSMSTrafficControl   = "SMS-TRAFFIC-CONTROL"
	ProductVoiceCall           = "VOICE-CALL"
	ProductVoiceFailed         = "VOICE-FAILED"
	ProductVoiceTTS            = "VOICE-TTS"
	ProductInAppVoice          = "IN-APP-VOICE"
	ProductWebSocketCall       = "WEBSOCKET-CALL"
	ProductASR                 = "ASR"
	ProductAMD                 = "AMD"
	ProductVerifyAPI           = "VERIFY-API"
	ProductVerifyV2            = "VERIFY-V2"
	ProductNumberInsight       = "NUMBER-INSIGHT"
	ProductNumberInsightV2     = "NUMBER-INSIGHT-V2"
	ProductConversationEvent   = "CONVERSATION-EVENT"
	ProductConversationMessage = "CONVERSATION-MESSAGE"
	ProductMessages            = "MESSAGES"
	ProductVideoAPI            = "VIDEO-API"
	ProductNetworkAPIEvent     = "NETWORK-API-EVENT"
	ProductReportsUsage        = "REPORTS-USAGE"
)

// Direction types.
const (
	DirectionInbound  = "inbound"
	DirectionOutbound = "outbound"
)

// RecordsRequest defines a records request message.
type RecordsRequest struct {
	AccountID string
	ID        string
	Product   string
	Direction string
}

// RecordsResponse defines a records response message.
type RecordsResponse struct {
	Links struct {
		Self struct {
			Href time.Time `json:"href"`
		} `json:"self"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
	} `json:"_links"`
	Cursor        string    `json:"cursor"`
	Iv            string    `json:"iv"`
	RequestID     string    `json:"request_id"`
	RequestStatus string    `json:"request_status"`
	ReceivedAt    time.Time `json:"received_at"`
	ItemsCount    int       `json:"items_count"`
	IdsNotFound   string    `json:"ids_not_found"`
	Product       string    `json:"product"`
	Records       []struct {
		AccountID            string    `json:"account_id"`
		MessageID            string    `json:"message_id"`
		AccountRef           string    `json:"account_ref"`
		ClientRef            string    `json:"client_ref"`
		Direction            string    `json:"direction"`
		From                 string    `json:"from"`
		To                   string    `json:"to"`
		ForcedFrom           string    `json:"forced_from"`
		ChangedFrom          string    `json:"changed_from"`
		Concatenated         string    `json:"concatenated"`
		MessageBody          string    `json:"message_body"`
		Network              string    `json:"network"`
		NetworkName          string    `json:"network_name"`
		Country              string    `json:"country"`
		CountryName          string    `json:"country_name"`
		DateReceived         time.Time `json:"date_received"`
		DateFinalized        time.Time `json:"date_finalized"`
		Latency              string    `json:"latency"`
		Status               string    `json:"status"`
		ErrorCode            string    `json:"error_code"`
		ErrorCodeDescription string    `json:"error_code_description"`
		Currency             string    `json:"currency"`
		TotalPrice           string    `json:"total_price"`
		ID                   string    `json:"id"`
		Dcs                  string    `json:"dcs"`
		ValidityPeriod       string    `json:"validity_period"`
		IPAddress            string    `json:"ip_address"`
		Udh                  string    `json:"udh"`
		WorkflowID           string    `json:"workflow_id"`
	} `json:"records"`
}

// Send the message using the specified Nexmo API client.
func (c *Report) Send(req *RecordsRequest) (*RecordsResponse, error) {
	var r, err = http.NewRequest("GET", apiRootv2+"/v2/reports/records", nil)
	if err != nil {
		return nil, err
	}

	var q = r.URL.Query()
	q.Add("account_id", req.AccountID)
	q.Add("id", req.ID)
	q.Add("product", req.Product)
	q.Add("direction", req.Direction)

	var auth = base64.RawURLEncoding.EncodeToString([]byte(c.client.apiKey + ":" + c.client.apiSecret))

	r.Header.Add("Authorization", "Basic "+auth)
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.client.HTTPClient.Do(r); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var records RecordsResponse
	err = json.Unmarshal(body, &records)
	if err != nil {
		return nil, err
	}
	return &records, nil
}
