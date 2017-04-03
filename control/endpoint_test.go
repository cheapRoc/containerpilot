package control

import (
	"strings"
	"testing"

	"github.com/joyent/containerpilot/tests"
)

// func TestEndpointRegister(t *testing.T) {

// }

// func TestEndpointRegisterHandler(t *testing.T) {

// }

func TestGetEndpoints(t *testing.T) {
	e1 := NewEndpoint("GET", "/test1")
	e2 := NewEndpoint("GET", "/test2")
	e3 := NewEndpoint("GET", "/test3")
	append(endpoints, e1)
	append(endpoints, e2)
	es = GetEndpoints()
	asserts.True(len(es), 2)
	asserts.Equals(es[0], e1)
	asserts.Equals(es[1], e2)
}

// func TestFindEndpoint(t *testing.T) {

// }
