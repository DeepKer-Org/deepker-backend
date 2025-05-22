package enum

type DeviceStatus string
type Sex string
type AlertStatus string

const (
	DeviceStatusInUse       DeviceStatus = "In Use"
	DeviceStatusFree        DeviceStatus = "Free"
	DeviceStatusUnavailable DeviceStatus = "Unavailable"
	DeviceStatusConnecting  DeviceStatus = "Connecting"

	SexMale   Sex = "M"
	SexFemale Sex = "F"

	AlertStatusAttended   AlertStatus = "Attended"
	AlertStatusUnattended AlertStatus = "Unattended"
)
