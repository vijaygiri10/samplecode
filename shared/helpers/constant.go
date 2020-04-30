package helpers

import "strings"

// NOTE: Please don't use these values, and try to replace them with EmailEventType values
// they are defined below
// Remove this comment if the note is fixed
var (
	Accepted   int64 = 0
	Enqueued   int64 = 1
	Deferred   int64 = 2
	Rejected   int64 = 3
	Delivered  int64 = 4
	Suppressed int64 = 5
	Bounced    int64 = 6
	Reported   int64 = 7
	Opened     int64 = 8
	Clicked    int64 = 9

	EmailMessageType int64 = 0
	SMSMessageType   int64 = 1
)

// EmailEventType : click, complaint etc...
type EmailEventType int16

const (
	EmailEventTypeInvalid EmailEventType = -1 + iota
	EmailEventTypeAccepted
	EmailEventTypeEnqueued
	EmailEventTypeDeferred
	EmailEventTypeRejected
	EmailEventTypeDelivered
	EmailEventTypeSuppressed
	EmailEventTypeBounced
	EmailEventTypeReported
	EmailEventTypeOpened
	EmailEventTypeClicked
	// the values for soft and hard bounce
	// are not defined by any other system as of now, but our module needs it
	// we are setting it to a higher value, so that they don't clash
	// with any other event type in near future
	// please fix them, if other modules also consider them as constants
	EmailEventTypeSoftBounce EmailEventType = 9997
	EmailEventTypeHardBounce EmailEventType = 9998
	EmailEventTypeMax        EmailEventType = 10000
)

var emailEventTypeMap = map[EmailEventType]string{
	EmailEventTypeInvalid:    "invalid",
	EmailEventTypeAccepted:   "accepted",
	EmailEventTypeEnqueued:   "enqueued",
	EmailEventTypeDeferred:   "deferred",
	EmailEventTypeRejected:   "rejected",
	EmailEventTypeDelivered:  "delivered",
	EmailEventTypeSuppressed: "suppressed",
	EmailEventTypeBounced:    "bounced",
	EmailEventTypeReported:   "reported",
	EmailEventTypeOpened:     "opened",
	EmailEventTypeClicked:    "clicked",
	EmailEventTypeSoftBounce: "soft_bounced",
	EmailEventTypeHardBounce: "hard_bounced",
}

var emailEventTypeReverseMap = map[string]EmailEventType{
	"invalid":      EmailEventTypeInvalid,
	"accepted":     EmailEventTypeAccepted,
	"enqueued":     EmailEventTypeEnqueued,
	"deferred":     EmailEventTypeDeferred,
	"rejected":     EmailEventTypeRejected,
	"delivered":    EmailEventTypeDelivered,
	"suppressed":   EmailEventTypeSuppressed,
	"bounced":      EmailEventTypeBounced,
	"reported":     EmailEventTypeReported,
	"opened":       EmailEventTypeOpened,
	"clicked":      EmailEventTypeClicked,
	"soft_bounced": EmailEventTypeSoftBounce,
	"hard_bounced": EmailEventTypeHardBounce,
}

// GetEventTypeFromName : Returns typed email event
func GetEventTypeFromName(event string) EmailEventType {
	return emailEventTypeReverseMap[strings.ToLower(event)]
}

// IntVal : Returns integer representation of the eventType
func (eventType EmailEventType) IntVal() int {
	return int(eventType)
}

// StringVal : Returns string representation of the eventType
func (eventType EmailEventType) StringVal() string {
	return emailEventTypeMap[eventType]
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// AccountStatus : Typed construct to refer to account status
type AccountStatus int16

const (
	AccountStatusInvalid AccountStatus = 0 + iota
	AccountStatusActive
	AccountStatusWarned
	AccountStatusSuspended
)

var accountStatusMap = map[AccountStatus]string{
	AccountStatusInvalid:   "Invalid",
	AccountStatusActive:    "Active",
	AccountStatusWarned:    "Warned",
	AccountStatusSuspended: "Suspended",
}

var accountStatusReverseMap = map[string]AccountStatus{
	"Invalid":   AccountStatusInvalid,
	"Active":    AccountStatusActive,
	"Warned":    AccountStatusWarned,
	"Suspended": AccountStatusSuspended,
}

// GetAccountStatusFromName : Returns typed account status
func GetAccountStatusFromName(status string) AccountStatus {
	return accountStatusReverseMap[status]
}

// GetAccountStatusFromInt : Returns typed account status
func GetAccountStatusFromInt(status int) AccountStatus {
	return AccountStatus(status)
}

// IntVal : Returns integer representation of the AccountStatus
func (status AccountStatus) IntVal() int {
	return int(status)
}

// StringVal : Returns string representation of the AccountStatus
func (status AccountStatus) StringVal() string {
	return accountStatusMap[status]
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// UserCriticalInfoType : Typed construct to refer to message types that are related to
// user critical info
type UserCriticalInfoType int16

const (
	UserCriticalInfoInvalid UserCriticalInfoType = iota
	AccountStatusUpdate
)

var userCriticalInfoTypeMap = map[UserCriticalInfoType]string{
	UserCriticalInfoInvalid: "Invalid",
	AccountStatusUpdate:     "AccountStatusUpdate",
}

// IntVal : Returns integer representation of the UserCriticalInfoType
func (status UserCriticalInfoType) IntVal() int {
	return int(status)
}

// StringVal : Returns string representation of the UserCriticalInfoType
func (status UserCriticalInfoType) StringVal() string {
	return userCriticalInfoTypeMap[status]
}

const (
	// LayoutISO : YYYY-MM-DD
	LayoutISO = "2006-01-02"
	// LayoutUS : MonthName Date, YYYY
	LayoutUS = "January 2, 2006"
)
