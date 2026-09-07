package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetClusterIdAttr(t *testing.T) {
	res := GetClusterIdAttr()
	assert.Equal(t, "runai.cluster_id", res)
}

func TestGetClientTypeAttr(t *testing.T) {
	res := GetClientTypeAttr()
	assert.Equal(t, "runai.client_type", res)
}
