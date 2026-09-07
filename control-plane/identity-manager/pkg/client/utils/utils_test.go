package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClusterIdAttr(t *testing.T) {
	res := GetClusterIdAttr()
	assert.Equal(t, "runai.cluster_id", res)
}

func TestGetClientTypeAttr(t *testing.T) {
	res := GetClientTypeAttr()
	assert.Equal(t, "runai.client_type", res)
}
