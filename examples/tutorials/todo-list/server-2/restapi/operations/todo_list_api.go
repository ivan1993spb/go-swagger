package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	httpkit "github.com/go-swagger/go-swagger/httpkit"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-2/restapi/operations/todos"
)

// NewTodoListAPI creates a new TodoList instance
func NewTodoListAPI(spec *loads.Document) *TodoListAPI {
	o := &TodoListAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*TodoListAPI The product of a tutorial on goswagger.io */
type TodoListAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/io.goswagger.examples.todo-list.v1+json" mime type
	JSONConsumer httpkit.Consumer

	// JSONProducer registers a producer for a "application/io.goswagger.examples.todo-list.v1+json" mime type
	JSONProducer httpkit.Producer

	// TodosAddOneHandler sets the operation handler for the add one operation
	TodosAddOneHandler todos.AddOneHandler
	// TodosDestroyOneHandler sets the operation handler for the destroy one operation
	TodosDestroyOneHandler todos.DestroyOneHandler
	// TodosFindTodosHandler sets the operation handler for the find todos operation
	TodosFindTodosHandler todos.FindTodosHandler
	// TodosUpdateOneHandler sets the operation handler for the update one operation
	TodosUpdateOneHandler todos.UpdateOneHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup
}

// SetDefaultProduces sets the default produces media type
func (o *TodoListAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TodoListAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *TodoListAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TodoListAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TodoListAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TodoListAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TodoListAPI
func (o *TodoListAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.TodosAddOneHandler == nil {
		unregistered = append(unregistered, "todos.AddOneHandler")
	}

	if o.TodosDestroyOneHandler == nil {
		unregistered = append(unregistered, "todos.DestroyOneHandler")
	}

	if o.TodosFindTodosHandler == nil {
		unregistered = append(unregistered, "todos.FindTodosHandler")
	}

	if o.TodosUpdateOneHandler == nil {
		unregistered = append(unregistered, "todos.UpdateOneHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TodoListAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TodoListAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]httpkit.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *TodoListAPI) ConsumersFor(mediaTypes []string) map[string]httpkit.Consumer {

	result := make(map[string]httpkit.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/io.goswagger.examples.todo-list.v1+json":
			result["application/io.goswagger.examples.todo-list.v1+json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TodoListAPI) ProducersFor(mediaTypes []string) map[string]httpkit.Producer {

	result := make(map[string]httpkit.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/io.goswagger.examples.todo-list.v1+json":
			result["application/io.goswagger.examples.todo-list.v1+json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TodoListAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *TodoListAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/"] = todos.NewAddOne(o.context, o.TodosAddOneHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/{id}"] = todos.NewDestroyOne(o.context, o.TodosDestroyOneHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/"] = todos.NewFindTodos(o.context, o.TodosFindTodosHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/{id}"] = todos.NewUpdateOne(o.context, o.TodosUpdateOneHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TodoListAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
