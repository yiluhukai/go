package util

import (
	"github.com/satori/go.uuid"
)

func Gen_uuid() string {
	u1 := uuid.Must(uuid.NewV4())

	return u1.String()
}
