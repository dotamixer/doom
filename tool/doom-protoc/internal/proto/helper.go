package proto

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/jhump/protoreflect/dynamic"
	"strings"
)

type sourceWithFiles interface {
	GetAllFiles() ([]*desc.FileDescriptor, error)
}

var _ sourceWithFiles = (*fileSource)(nil)

var base64Codecs = []*base64.Encoding{base64.StdEncoding, base64.URLEncoding, base64.RawStdEncoding, base64.RawURLEncoding}

var printer = &protoprint.Printer{
	Compact:                  true,
	OmitComments:             protoprint.CommentsNonDoc,
	SortElements:             true,
	ForceFullyQualifiedNames: true,
}

// GetDescriptorText returns a string representation of the given descriptor.
// This returns a snippet of proto source that describes the given element.
func GetDescriptorText(dsc desc.Descriptor, _ DescriptorSource) (string, error) {
	// Note: DescriptorSource is not used, but remains an argument for backwards
	// compatibility with previous implementation.
	txt, err := printer.PrintProtoToString(dsc)
	if err != nil {
		return "", err
	}
	// callers don't expect trailing newlines
	if txt[len(txt)-1] == '\n' {
		txt = txt[:len(txt)-1]
	}
	return txt, nil
}

// fetchAllExtensions recursively fetches from the server extensions for the given message type as well as
// for all message types of nested fields. The extensions are added to the given dynamic registry of extensions
// so that all server-known extensions can be correctly parsed by grpcurl.
func fetchAllExtensions(source DescriptorSource, ext *dynamic.ExtensionRegistry, md *desc.MessageDescriptor, alreadyFetched map[string]bool) error {
	msgTypeName := md.GetFullyQualifiedName()
	if alreadyFetched[msgTypeName] {
		return nil
	}
	alreadyFetched[msgTypeName] = true
	if len(md.GetExtensionRanges()) > 0 {
		fds, err := source.AllExtensionsForType(msgTypeName)
		if err != nil {
			return fmt.Errorf("failed to query for extensions of type %s: %v", msgTypeName, err)
		}
		for _, fd := range fds {
			if err := ext.AddExtension(fd); err != nil {
				return fmt.Errorf("could not register extension %s of type %s: %v", fd.GetFullyQualifiedName(), msgTypeName, err)
			}
		}
	}
	// recursively fetch extensions for the types of any message fields
	for _, fd := range md.GetFields() {
		if fd.GetMessageType() != nil {
			err := fetchAllExtensions(source, ext, fd.GetMessageType(), alreadyFetched)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type notFoundError string

func notFound(kind, name string) error {
	return notFoundError(fmt.Sprintf("%s not found: %s", kind, name))
}

func (e notFoundError) Error() string {
	return string(e)
}

func isNotFoundError(err error) bool {
	_, ok := err.(notFoundError)
	return ok
}

func parseSymbol(svcAndMethod string) (string, string) {
	pos := strings.LastIndex(svcAndMethod, "/")
	if pos < 0 {
		pos = strings.LastIndex(svcAndMethod, ".")
		if pos < 0 {
			return "", ""
		}
	}
	return svcAndMethod[:pos], svcAndMethod[pos+1:]
}

func parseSvc(svc string) (string, error) {
	pos := strings.LastIndex(svc, ".")
	if pos < 0 {
		return "", errors.New("invalid svc. ")
	}

	return svc[:pos], nil
}
