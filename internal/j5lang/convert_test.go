package j5lang

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pentops/j5/gen/j5/schema/v1/schema_j5pb"
	"github.com/pentops/j5/gen/j5/source/v1/source_j5pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestConvert(t *testing.T) {

	input := strings.Join([]string{
		`package pentops.j5lang.example`,
		`version = "v1"`,
		`object Foo {`,
		`| Foo Object Description`,
		`  `,
		`  `,
		`  field bar_field string {`,
		`    validate.required = true`,
		`    validate.min_length = 1`,
		`  }`,
		`  `,
		`  field baz_field object {`,
		`	 ref path.to.Type`,
		`  }`,
		`}`,
	}, "\n")

	file, err := ParseFile(input)
	if err != nil {
		t.Fatal(err)
	}

	source, err := ConvertFile(file)
	if err != nil {
		t.Fatal(err)
	}

	want := &source_j5pb.SourceFile{
		Package: "pentops.j5lang.example",
		Version: "v1",
		Schemas: []*schema_j5pb.RootSchema{{
			Type: &schema_j5pb.RootSchema_Object{
				Object: &schema_j5pb.Object{
					Name:        "Foo",
					Description: "Foo Object Description",
					Properties: []*schema_j5pb.ObjectProperty{{
						Name:     "bar_field",
						Required: true,
						Schema: &schema_j5pb.Field{
							Type: &schema_j5pb.Field_String_{
								String_: &schema_j5pb.StringField{
									Rules: &schema_j5pb.StringField_Rules{
										MinLength: ptr(uint64(1)),
									},
								},
							},
						},
					}, {
						Name: "baz_field",
						Schema: &schema_j5pb.Field{
							Type: &schema_j5pb.Field_Object{
								Object: &schema_j5pb.ObjectField{
									Schema: &schema_j5pb.ObjectField_Ref{
										Ref: &schema_j5pb.Ref{
											Package: "path.to",
											Schema:  "Type",
										},
									},
								},
							},
						},
					}},
				},
			},
		}},
	}

	cmpProto(t, source, want)

}

func ptr[T any](v T) *T {
	return &v
}

func cmpProto(t testing.TB, a, b proto.Message) {
	diff := cmp.Diff(a, b, protocmp.Transform())
	if diff != "" {
		t.Error(diff)
	}

}