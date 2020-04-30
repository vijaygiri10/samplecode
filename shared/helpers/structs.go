package helpers

import (
	"regexp"
	"time"
)

//Event ...
type Event struct {
	AccountID   string      `json:"account_id"`
	MessageID   string      `json:"message_id"`
	EventType   string      `json:"event_type"`
	MetaData    interface{} `json:"metadata"`
	MessageType string      `json:"message_type"`
	CreatedAT   time.Time   `json:"created_at"`
}

type AcceptedMetaData struct {
	SuppressionListID string   `json:"suppression_list_id"`
	TrackingDomain    string   `json:"tracking_domain"`
	Subject           string   `json:"subject"`
	RcptTo            string   `json:"to_address"`
	MailFrom          string   `json:"from_address"`
	CC                []string `json:"cc"`
	BCC               []string `json:"bcc"`
	ReturnPath        string   `json:"return_path"`
	SendingIP         string   `json:"sending_ip"`
	Status            int64    `json:"status"`
	MessageSize       int64    `json:"message_size"`
	Tags              []string `json:"tags"`
}

type EnqueuedMetaData struct {
	RcptTo              string    `json:"to_address"`
	Subject             string    `json:"subject"`
	EmailProcessStartAT time.Time `json:"email_process_start_at"`
	GlusterFSWriteAT    time.Time `json:"glusterFS_write_at"`
}

type SuppressedMetaData struct {
	RcptTo   string `json:"to_address"`
	Subject  string `json:"subject"`
	MailFrom string `json:"from_address"`
}

type ClickedMetaData struct {
	IPAddress   string `json:"ip"`
	DigestURL   string `json:"url_digest"`
	Browser     string `json:"browser"`
	OriginalURL string `json:"url"`
}
type OpenedMetaData struct {
	EmailUUID string `json:"email_uuid"`
	Browser   string `json:"browser"`
	IPAddress string `json:"ip"`
}

//Email Common Struct
type Email struct {
	ID             string                 `json:"id"`
	AccountID      string                 `json:"account_id"`
	MailFrom       string                 `json:"from_email"`
	RcptTo         interface{}            `json:"rcpt_to"`
	Cc             []string               `json:"cc"`
	Bcc            []string               `json:"bcc"`
	Archive        []map[string]string    `json:"archive"`
	Tags           []string               `json:"tags"`
	Options        map[string]interface{} `json:"options"`
	MetaData       map[string]interface{} `json:"metadata"`
	RemoteIP       string                 `json:"remote_ip"`
	Body           string                 `json:"body"`
	TrackingDomain string                 `json:"tracking_domain"`
	Subject        string                 `json:"subject"`
	// Changed INT64 to Time.Time as per Postgres
	ReceivedAT          time.Time `json:"received_at"`
	GlusterFSWriteAT    time.Time `json:"glusterFS_write_at"`
	EmailProcessStartAT time.Time `json:"email_process_start"`
	JMTASentAT          time.Time `json:"sent_to_jmta"`
	SuppressedAT        time.Time `json:"suppressed_at"`

	// Newly Added Fields
	SuppressionListID string `json:"suppression_list_id"` //not null
	CategoryID        string `json:"category_id"`         //not null
	ReturnPath        string `json:"return_path"`         //not null
	SendingIP         string `json:"sending_ip"`          //not null
	MessageSize       int64  `json:"message_size"`        //not null
	TrackingDomainID  string `json:"tracking_domain_id"`
}

type Results struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
type errs struct {
	Errors []map[string]string
}
type Responce struct {
	Result        Results `json:"result"`
	ErrorResponse errs    `json:"errors"`
}

func validateEmail(email interface{}) bool {
	switch email.(type) {
	case string:
		return regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email.(string))
	case []string:
		emails := email.([]string)
		for _, email := range emails {
			if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email) {
				return false
			}
		}
	default:
		return false
	}
	return true
}

// Validate will validate email msg Attribute
func (m *Email) Validate() (status bool, errorResponse *Responce) {
	resp := Responce{}
	status = true
	if m.MailFrom == "" {
		status = false
		resp.ErrorResponse.Errors = append(resp.ErrorResponse.Errors, map[string]string{"Message": "Valid From address is required.", "Code": "4002"})
	} else if !validateEmail(m.MailFrom) {
		status = false
		resp.ErrorResponse.Errors = append(resp.ErrorResponse.Errors, map[string]string{"Message": "From address is invalid.", "Code": "4003"})
	}
	if m.RcptTo.(string) == "" {
		status = false
		resp.ErrorResponse.Errors = append(resp.ErrorResponse.Errors, map[string]string{"Message": "Valid recipient address is required.", "Code": "4004"})
	} else if !validateEmail(m.RcptTo) {

		status = false
		resp.ErrorResponse.Errors = append(resp.ErrorResponse.Errors, map[string]string{"Message": "Recipient address is invalid.", "Code": "4005"})
	}

	return status, &resp
}

type ClicksInfo struct {
	AccountID   string    `json:"account_id"`
	EmailUUID   string    `json:"email_uuid"`
	Browser     string    `json:"browser"`
	IPAddress   string    `json:"ip_address"`
	LinkKeyUUID string    `json:"link_key_uuid"`
	OriginalURL string    `json:"url"`
	RecordedAT  time.Time `json:"recorded_at"`
}

type Opens struct {
	AccountID  string    `json:"account_id"`
	EmailUUID  string    `json:"email_uuid"`
	Browser    string    `json:"browser"`
	IPAddress  string    `json:"ip_address"`
	RecordedAT time.Time `json:"recorded_at"`
}
