package dm

import (
	"path"
	"strconv"
)

import (
	"github.com/dotamixer/doom/tool/protoc-gen-dm/generator"
	pb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// Paths for packages used by code generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
	grpcPkgPath = "google.golang.org/grpc"
)

// GenerateImports generates the import declaration for this file.
func (g *dm) GenerateImports(file *generator.FileDescriptor, imports map[generator.GoImportPath]generator.GoPackageName) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}

		g.P("import (")
		g.P(grpcPkgPath, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, grpcPkgPath)))
		g.P(")")
		g.P()

		// We need to keep track of imported packages to make sure we don't produce
		// a name collision when generating types.
		pkgImports = make(map[generator.GoPackageName]bool)
		for _, name := range imports {
			pkgImports[name] = true
		}


}

func (g *dm) generateNewClient(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	g.P("var (")
	g.P("ServiceName = ", strconv.Quote(*file.Package))
	g.P(")")
	g.P()

	origServName := service.GetName()
	servName := generator.CamelCase(origServName)


	g.P("func New", servName, "Client() (client ", servName,"Client, err error) {")
	g.P("var (")
	g.P("conn *grpc.ClientConn")
	g.P(")")
	g.P("conn, err = grpc.Dial(",  strconv.Quote("consul://default/" + *file.Package))
	g.P("grpc.WithInsecure())")
	g.P("if err != nil {")
	g.P("return")
	g.P("}")
	g.P()

	g.P("client = New", servName, "Client(conn)")
	g.P("return")
	g.P("}")
}
