package utils

import "github.com/google/uuid"

func ParseUUIDPtr(s string) *uuid.UUID {
	uuid, err := uuid.Parse(s)
	if err != nil {
		return nil
	}
	return &uuid
}
