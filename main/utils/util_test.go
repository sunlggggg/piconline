package utils

import (
	"testing"
)

func TestHash(t *testing.T) {
	t.Log("hash", Hash("ddd", 10))
}
