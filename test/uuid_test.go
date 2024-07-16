package test

import (
	uuid2 "github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid, _ := uuid2.NewUUID()
	t.Log(uuid.String())
	return
}
