package testutils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"sync"
)

////////////////////////////////////////////////////////////////////////////////////////////
// RandomValues utility functions maintain a set of random values to be used by testing.
////////////////////////////////////////////////////////////////////////////////////////////

// we hold a thread-safe counter that increases every time we want new random number
var (
	counter      = uint(rand.Int31n(100))
	counterMutex sync.Mutex
)

func nextRandom() uint {
	counterMutex.Lock()
	defer counterMutex.Unlock()
	counter += uint(rand.Int31n(99) + 1)
	return counter
}

var randomTenantId = uint(rand.Int31n(2048) + 1)

// RandomValues is a set of random values, so tests can use them
type RandomValues struct {
	TenantId     uint
	TenantIdStr  string
	ClusterId    string
	DepartmentId string
	ProjectId    uint
	ProjectIdStr string

	IsForbidden bool // is the requesting API forbidden from seeing this?
}

func NewRandomValues() *RandomValues {
	result := RandomValues{
		TenantId:     randomTenantId,
		ClusterId:    uuid.New().String(),
		DepartmentId: fmt.Sprintf("%d", nextRandom()),
		ProjectId:    nextRandom(),
	}

	result.TenantIdStr = fmt.Sprintf("%v", result.TenantId)
	result.ProjectIdStr = fmt.Sprintf("%v", result.ProjectId)

	return &result
}

func (self RandomValues) Clone() *RandomValues {
	return &RandomValues{
		TenantId:     self.TenantId,
		TenantIdStr:  self.TenantIdStr,
		ProjectId:    self.ProjectId,
		ProjectIdStr: self.ProjectIdStr,
		DepartmentId: self.DepartmentId,
		ClusterId:    self.ClusterId,
		IsForbidden:  self.IsForbidden,
	}
}

// DifferentTenant returns the set of values with different tenant ID than whatever was originally chosen
func (self *RandomValues) CloneWithDifferentTenant() *RandomValues {
	result := self.Clone()
	result.TenantId = nextRandom()
	result.TenantIdStr = fmt.Sprintf("%v", result.TenantId)
	return result.CloneWithDifferentCluster()
}

// CloneWithDifferentCluster returns the set of values with different cluster ID than whatever was originally chosen
func (self *RandomValues) CloneWithDifferentCluster() *RandomValues {
	result := self.Clone()
	result.ClusterId = uuid.New().String()
	return result.CloneWithDifferentDepartment()
}

func (self *RandomValues) CloneWithSpecificCluster(clusterId string) *RandomValues {
	result := self.Clone()
	result.ClusterId = clusterId
	return result.CloneWithDifferentDepartment()
}

// CloneWithDifferentDepartment returns the set of values with different department ID than whatever was originally chosen
func (self *RandomValues) CloneWithDifferentDepartment() *RandomValues {
	result := self.Clone()
	result.DepartmentId = fmt.Sprintf("%d", nextRandom())
	return result.CloneWithDifferentProject()
}

// CloneWithDifferentProject returns the set of values with different cluster ID than whatever was originally chosen
func (self *RandomValues) CloneWithDifferentProject() *RandomValues {
	result := self.Clone()
	result.ProjectId = nextRandom()
	result.ProjectIdStr = fmt.Sprintf("%v", result.ProjectId)
	return result
}

func (self *RandomValues) CloneWithSpecificProject(projectId uint) *RandomValues {
	result := self.Clone()
	result.ProjectId = projectId
	result.ProjectIdStr = fmt.Sprintf("%v", result.ProjectId)
	return result
}

// Forbidden marks this set of values as forbidden for the API
func (self *RandomValues) Forbidden() *RandomValues {
	self.IsForbidden = true
	return self
}

// Allowed marks this set of values as forbidden for the API
func (self *RandomValues) Allowed() *RandomValues {
	self.IsForbidden = false
	return self
}

func (self *RandomValues) ProjectIdParam() string {
	return fmt.Sprintf("projectId=%s", self.ProjectIdStr)
}
