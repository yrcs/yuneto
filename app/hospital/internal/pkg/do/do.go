package do

import "time"

type HospitalSetting struct {
	Id                 uint64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Name               string
	RegistrationNumber string
	ContactPerson      string
	ContactMobile      string
	Locked             uint8
	ApiUrl             string
	SignatureKey       string
}
