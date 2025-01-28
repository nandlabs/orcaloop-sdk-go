package utils

import "oss.nandlabs.io/golly/uuid"

func GenerateId() string {
	uid, _ := uuid.V4()
	if uid != nil {

	}
	return uid.String()
}
