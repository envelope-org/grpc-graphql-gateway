package builder

import (
	"fmt"
	"strings"

	"github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql/spec"
)

// Message builder generates common object definition.
// type {...} on GraphQL, and Object on Go program.
type Message struct {
	m *spec.Message
}

func NewMessage(m *spec.Message) *Message {
	return &Message{
		m: m,
	}
}

func (b *Message) BuildQuery() (string, error) {
	lines := []string{}
	if c := b.m.Comment(spec.GraphqlComment); c != "" {
		lines = append(lines, c)
	}
	lines = append(lines, fmt.Sprintf("type %s {", b.m.Name()))

	for _, f := range b.m.Fields() {
		fieldType := f.GraphqlType()
		if f.IsRepeated() {
			fieldType = "[" + fieldType + "]"
		}
		if !f.IsOptional() {
			fieldType += "!"
		}

		if c := f.Comment(spec.GraphqlComment); c != "" {
			lines = append(lines, c)
		}

		lines = append(lines, fmt.Sprintf(
			"  %s: %s",
			f.Name(),
			fieldType,
		))
	}

	lines = append(lines, fmt.Sprintf("} # message %s in %s\n", b.m.SingleName(), b.m.Filename()))
	return strings.Join(lines, "\n"), nil
}

func (b *Message) BuildProgram() (string, error) {
	var fields []string

	for _, f := range b.m.Fields() {
		fieldType := f.GraphqlGoType()
		if f.IsRepeated() {
			fieldType = "graphql.NewList(" + fieldType + ")"
		}
		if !f.IsOptional() {
			fieldType = "graphql.NewNonNull(" + fieldType + ")"
		}

		fields = append(fields, strings.TrimSpace(fmt.Sprintf(`
			%s
			"%s": &graphql.Field{
				Type: %s,
			},`,
			f.Comment(spec.GoComment),
			f.Name(),
			fieldType,
		)))
	}

	return strings.TrimSpace(fmt.Sprintf(`
%s
var %s = graphql.NewObject(graphql.ObjectConfig{
	Name: "%s",
	Fields: graphql.Fields{
		%s
	},
}) // message %s in %s`,
		b.m.Comment(spec.GoComment),
		spec.PrefixType(b.m.Name()),
		b.m.Name(),
		strings.TrimSpace(strings.Join(fields, "\n")),
		b.m.SingleName(),
		b.m.Filename(),
	)) + "\n", nil
}