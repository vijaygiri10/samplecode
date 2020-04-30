package helpers

import (
	"encoding/json"
	"fmt"
	"jetsend_opens/shared/jwt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func WriteResponse(write http.ResponseWriter, response interface{}, statusCode int) {
	write.Header().Set("Content-type", "application/json")
	write.WriteHeader(statusCode)
	json.NewEncoder(write).Encode(response)
}

func GetAccountIDFromJWT(token string) (string, error) {
	token = strings.TrimSpace(strings.Replace(token, "Bearer", "", -1))

	claims, err := jwt.GetClaimFromJWTToken(token)
	if err != nil {
		log.Println("Signup Invalid JWT Token err : ", err, token)
		return "", err
	}

	return claims.AccountID, nil
}

func ExecptionHandler(functionName string) {
	if err := recover(); err != nil {
		fmt.Println("Expection: ", functionName, " error info : ", err)
	}
}

func CreateLoggerFile(name string, file_path string) (*log.Logger, error) {
	log_file, err := os.OpenFile(file_path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := log.New(log_file, name+":", log.Ldate|log.Ltime)
	return logger, nil
}

func ConvertUTC(from string, to string, timezone string) (string, string) {
	StartTime, err := time.Parse(time.RFC3339, from+"Z")
	if err != nil {
		fmt.Println(err)
	}
	EndTime, err := time.Parse(time.RFC3339, to+"Z")
	if err != nil {
		fmt.Println(err)
	}
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println("Location is not valid")
	}
	// Add location to time
	StartTime = StartTime.In(loc)
	EndTime = EndTime.In(loc)
	if err != nil {
		fmt.Println(err)
	}
	return StartTime.Format(time.RFC3339), EndTime.Format(time.RFC3339)
}

//GetEventMataData ...
func GetEventMataData(eventType int64, eventData interface{}) (event Event, err error) {
	event = Event{}

	switch eventType {
	case Accepted:
		msg := eventData.(Email)
		mataData := AcceptedMetaData{}
		mataData.SuppressionListID = msg.SuppressionListID
		mataData.TrackingDomain = msg.TrackingDomain
		mataData.Subject = msg.Subject
		mataData.RcptTo = msg.RcptTo.(string)
		mataData.MailFrom = msg.MailFrom

		if len(msg.Cc) <= 0 {
			mataData.CC = []string{}
		} else {
			mataData.CC = msg.Cc
		}

		if len(msg.Bcc) <= 0 {
			mataData.BCC = []string{}
		} else {
			mataData.BCC = msg.Bcc
		}

		mataData.ReturnPath = msg.ReturnPath
		mataData.SendingIP = msg.SendingIP
		mataData.Status = Accepted
		mataData.MessageSize = msg.MessageSize

		if len(msg.Tags) <= 0 {
			mataData.Tags = []string{}
		} else {
			mataData.Tags = msg.Tags
		}

		event.AccountID = msg.AccountID
		event.MessageID = msg.ID
		event.MetaData = mataData
		event.EventType = "Accepted"
		event.CreatedAT = msg.ReceivedAT
	case Enqueued:
		msg := eventData.(Email)
		mataData := EnqueuedMetaData{}
		mataData.Subject = msg.Subject
		mataData.RcptTo = msg.RcptTo.(string)
		mataData.EmailProcessStartAT = msg.EmailProcessStartAT
		mataData.GlusterFSWriteAT = msg.GlusterFSWriteAT
		event.AccountID = msg.AccountID
		event.MessageID = msg.ID
		event.MetaData = mataData
		event.EventType = "Enqueued"
		event.CreatedAT = msg.JMTASentAT
	case Deferred:
		return event, fmt.Errorf("Deferred Event Not Supported")
	case Rejected:
	case Delivered:
		return event, fmt.Errorf("Deferred Event Not Supported")
	case Suppressed:
		msg := eventData.(Email)
		mataData := SuppressedMetaData{}
		mataData.RcptTo = msg.RcptTo.(string)
		mataData.MailFrom = msg.MailFrom
		mataData.Subject = msg.Subject
		event.AccountID = msg.AccountID
		event.MessageID = msg.ID
		event.MetaData = mataData
		event.EventType = "Suppressed"
		event.CreatedAT = msg.SuppressedAT
	case Bounced:
		return event, fmt.Errorf("Deferred Event Not Supported")
	case Reported:
		return event, fmt.Errorf("Reported Event Not Supported")
	case Opened:
		open := eventData.(Opens)
		mataData := OpenedMetaData{}
		mataData.Browser = open.Browser
		mataData.EmailUUID = open.EmailUUID
		mataData.IPAddress = open.IPAddress
		event.MetaData = mataData
		event.AccountID = open.AccountID
		event.MessageID = open.EmailUUID
		event.EventType = "Opened"
		event.CreatedAT = open.RecordedAT
	case Clicked:
		cl := eventData.(ClicksInfo)
		mataData := ClickedMetaData{}
		mataData.Browser = cl.Browser
		mataData.DigestURL = cl.LinkKeyUUID
		mataData.IPAddress = cl.IPAddress
		mataData.OriginalURL = cl.OriginalURL
		event.MetaData = mataData
		event.AccountID = cl.AccountID
		event.MessageID = cl.EmailUUID
		event.EventType = "Clicked"
		event.CreatedAT = cl.RecordedAT
	default:
		return event, fmt.Errorf("Invalid eventType : %d", eventType)

	}

	//mataData, _ := json.Marshal(mataDataMap)
	event.MessageType = strconv.Itoa(int(EmailMessageType))
	fmt.Println("GetEventMataData : ", event)
	return event, nil
}

// ParseEvent : Extracts event structure from bytes
func ParseEvent(msg []byte) (*Event, error) {
	var event Event
	if err := json.Unmarshal(msg, &event); err != nil {
		return nil, fmt.Errorf("parseEvent: error encountered while unmarshalling event  %v", err)
	}
	return &event, nil
}

//ParseInterfaceToStringArr ...
func ParseInterfaceToStringArr(data interface{}) (ArrString []string, err error) {

	switch data.(type) {
	case []interface{}:
		for _, val := range data.([]interface{}) {
			ArrString = append(ArrString, val.(string))
		}
	case interface{}:
		ArrString = append(ArrString, data.(string))
	default:
		fmt.Println("Invalid Data Type of Tags in Options : ", data)
		return ArrString, fmt.Errorf("Invalid Data Type of Tags in Options : ", data)
	}
	return
}
