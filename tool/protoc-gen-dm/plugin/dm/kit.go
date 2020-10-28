package dm

import (
	"strings"
)

import (
	"github.com/dotamixer/doom/tool/protoc-gen-dm/generator"
)

func init() {
	generator.RegisterPlugin(new(dm))
}

// dm is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for go-dm support.
type dm struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "dm".
func (g *dm) Name() string {
	return "kit"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.
var (
	pkgImports map[generator.GoPackageName]bool
)

// Init initializes the plugin.
func (g *dm) Init(gen *generator.Generator) {
	g.gen = gen
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *dm) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *dm) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *dm) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *dm) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}

	for i, service := range file.FileDescriptorProto.Service {
		g.generateNewClient(file, service, i)
	}
}

// reservedClientName records whether a client name is reserved on the client side.
var reservedClientName = map[string]bool{
	// TODO: do we need any in go-dm?
}

func unexport(s string) string {
	if len(s) == 0 {
		return ""
	}
	name := strings.ToLower(s[:1]) + s[1:]
	if pkgImports[generator.GoPackageName(name)] {
		return name + "_"
	}
	return name
}
