package helpers

import (
	"context"
	"fmt"
	"jetsend_opens/shared/log"
	"strconv"
	"strings"
	"time"
)

// TODO: Pick these constants from config
const (
	// Account Info Table: It's a hash table and serves as cached user info
	//
	// FORMAT:
	// account_info_[account_id] : { status: 1 }
	// status : Integer representation of valid Account Status, like : Active, Suspended etc...

	// UserAccountInfoTable : this is used to identify user's basic info in redis, it's a hash table
	UserAccountInfoTable = "account_info"
	// StatusKeyInUserAccountInfoTable :This resides in UserAccountInfoTable
	StatusKeyInUserAccountInfoTable = "status"

	// HourlyEventCounterTable : It's a hash table and is used to identify counters in redis
	//
	// FORMAT:
	// date : dd-mm-yyyy  -> 31-01-2020
	// hour : 1   -> inidicates 1st hour of the day, similarly there could be 2, 3 ... 24
	// event_type : clicks, complaints etc... NOTE: It's string representation of these events, and not any integer
	//
	// Accountwise counters:
	// Refers to account specific hourly event counter, this gives us insights into the complaints, clicks etc... on account level
	// hourly_event_counter_[account_id] : { [event_type]_[date]_[hour]: 50, [event_type]_[date]_[hour]: 3 }
	//
	// Global counters:
	// Refers to global hourly event counter, this gives us insights into the complaints, clicks etc... at our product level
	// hourly_event_counter : { [event_type]_[date]_[hour]: 50, [event_type]_[date]_[hour]: 3 }
	HourlyEventCounterTable = "hourly_event_counter"

	// RejectedKeyInHoursTable : This resides in  HourlyEventCounterTable
	RejectedKeyInHoursTable = "rejected"
	// OpenedKeyInHoursTable : This resides in  HourlyEventCounterTable
	OpenedKeyInHoursTable = "opened"
	// ClickedKeyInHoursTable : This resides in  HourlyEventCounterTable
	ClickedKeyInHoursTable = "clicked"
	// AcceptedKeyInHoursTable : This resides in  HourlyEventCounterTable
	AcceptedKeyInHoursTable = "accepted"
	// EnqueuedKeyInHoursTable : This resides in  HourlyEventCounterTable
	EnqueuedKeyInHoursTable = "enqueued"
	// SuppressedKeyInHoursTable : This resides in  HourlyEventCounterTable
	SuppressedKeyInHoursTable = "suppressed"
	// ReportedKeyInHoursTable : This resides in  HourlyEventCounterTable
	ReportedKeyInHoursTable = "reported"
	// DeliveredKeyInHoursTable : This resides in  HourlyEventCounterTable
	DeliveredKeyInHoursTable = "delivered"
	// SoftBounceKeyInHoursTable : This resides in  HourlyEventCounterTable
	SoftBounceKeyInHoursTable = "soft_bounced"
	// HardBounceKeyInHoursTable : This resides in  HourlyEventCounterTable
	HardBounceKeyInHoursTable = "hard_bounced"
)

// PartitionTableKey :
// can't do much, the implementation was discussed and was agreed upon something which resulted into such a poor implementation of this function : /
func PartitionTableKey(ctx context.Context, key string) (eventType EmailEventType, date time.Time, hour int, err error) {
	tokens := strings.Split(key, "_")
	if len(tokens) == 3 {
		// 0th part: rejected, open, clicked, accepted, enqueued, suppressed, complaints, delivered
		// 1st part: date
		// 2nd part: time
		eventType = GetEventTypeFromName(tokens[0])
		date, err = time.Parse(LayoutISO, tokens[1])
		if err != nil {
			log.Error(ctx, "Failed to extract date from the key: ", key, " error: ", err)
			return
		}
		hour, err = strconv.Atoi(tokens[2])
		if err != nil {
			log.Error(ctx, "Failed to extract date from the key: ", key, " error: ", err)
			return
		}

	} else if len(tokens) == 4 {
		// 0th part: soft, hard
		// 1st part: bounce
		// 2nd part: date
		// 3rd part: time
		eventType = GetEventTypeFromName(tokens[0] + "_" + tokens[1])
		date, err = time.Parse(LayoutISO, tokens[2])
		if err != nil {
			log.Error(ctx, "Failed to extract date from the key: ", key, " error: ", err)
			return
		}
		hour, err = strconv.Atoi(tokens[3])
		if err != nil {
			log.Error(ctx, "Failed to extract date from the key: ", key, " error: ", err)
			return
		}
	} else {
		return eventType, date, hour, fmt.Errorf("PartitionKey(): doesn't know how to extract date form the key: %s", key)
	}

	return

}

// GetAccountWiseHourlyCounterTableName :
func GetAccountWiseHourlyCounterTableName(accountID string) string {
	return HourlyEventCounterTable + "_" + accountID
}

// GetGlobalHourlyCounterTableName :
func GetGlobalHourlyCounterTableName() string {
	return HourlyEventCounterTable
}

// GetHourlyComplaintKey :
func GetHourlyComplaintKey(date string, hour string) string {
	return ReportedKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlyDeliveredKey :
func GetHourlyDeliveredKey(date string, hour string) string {
	return DeliveredKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlySoftBounceKey :
func GetHourlySoftBounceKey(date string, hour string) string {
	return SoftBounceKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlyHardBounceKey :
func GetHourlyHardBounceKey(date string, hour string) string {
	return HardBounceKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlyOpenedKey :
func GetHourlyOpenedKey(date string, hour string) string {
	return OpenedKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlyClickedKey :
func GetHourlyClickedKey(date string, hour string) string {
	return ClickedKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlyAcceptedKey :
func GetHourlyAcceptedKey(date string, hour string) string {
	return AcceptedKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlyEnqueuedKey :
func GetHourlyEnqueuedKey(date string, hour string) string {
	return EnqueuedKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlySuppressedKey :
func GetHourlySuppressedKey(date string, hour string) string {
	return SuppressedKeyInHoursTable + "_" + date + "_" + hour
}

// GetHourlyRejectedKey :
func GetHourlyRejectedKey(date string, hour string) string {
	return RejectedKeyInHoursTable + "_" + date + "_" + hour
}

// GetAccountInfoTable :
func GetAccountInfoTable(accountID string) string {
	return UserAccountInfoTable + "_" + accountID
}

// GetAccountStatusKey :
func GetAccountStatusKey() string {
	return StatusKeyInUserAccountInfoTable
}
