package control

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

var endpoints = []Endpoint{}

type EndpointHandler func(interface{}, *http.ResponseWriter, *http.Request) error
type Endpoint struct {
	Data    interface{}
	Path    string
	Method  string
	Handler *EndpointHandler
}

func (eh EndpointHandler) ServeHTTP(w *http.ResponseWriter, req *http.Request) error {
	// TODO: Define standard logic that tests request scope based, on method and
	// path
	error := eh(eh.Data, w, req)
	if error != nil {
		return error
	}
}

// Name defines a short name for our Endpoint
func (e Endpoint) Name() string {
	return fmt.Sprintf("%v %v", e.Method, e.Path)
}

// Register an endpoint into a composite type referencing all global handlers
func (e Endpoint) Register() {
	endpoint := FindEndpoint(method, path)

	if endpoint != nil {
		return errors.New("Endpoint already registered by '%s'", e.Name)
	}

	log.Debugf("control: register endpoint '%v %v'", e.Method, e.Path)
	append(endpoints, e)
}

// RegisterEndpoint adds a new endpoint to the list of endpoints
func (e *Endpoint) RegisterHandler(data interface{}, eh *EndpointHandler) error {
	log.Debugf("control: register handler for '%s'", e.Name)
	e.Data = data
	e.Handler = eh
	e.Register()
	return e
}

// NewEndpoint creates a new endpoint by validating uniqueness of its fields
func NewEndpoint(method string, path string) (*Endpoint, error) {
	endpoint := FindEndpoint(method, path)
	if endpoint != nil {
		return errors.New("Endpoint already registered by '%s'", e.Name)
	}

	return &Endpoint{
		Path: path,
		Method: method,
	}
}

// GetEndpoints returns a slice of all private endpoints
func GetEndpoints() []Endpoint {
	return endpoints
}

// FindEndpoint returns an endpoint from within our global slice of all
// endpoints
func FindEndpoint(method string, path string) *Endpoint {
	for i, e := range endpoints {
		if e.Path == path && e.Method == method {
			return e
		}
	}
}
