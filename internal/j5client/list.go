package j5client

import (
	"fmt"
	"strings"

	"github.com/pentops/j5/gen/j5/client/v1/client_j5pb"
	"github.com/pentops/j5/gen/j5/list/v1/list_j5pb"
	"github.com/pentops/j5/gen/j5/schema/v1/schema_j5pb"
	"github.com/pentops/j5/internal/j5reflect"
)

func buildListRequest(response j5reflect.RootSchema) (*client_j5pb.ListRequest, error) {

	responseObj, ok := response.(*j5reflect.ObjectSchema)
	if !ok {
		return nil, fmt.Errorf("expected object schema, got %T", response)
	}

	var foundArray *j5reflect.ArrayField

	for _, field := range responseObj.Properties {
		asArray, ok := field.Schema.(*j5reflect.ArrayField)
		if !ok {
			continue
		}
		if foundArray != nil {
			return nil, fmt.Errorf("found multiple arrays in response")
		}

		foundArray = asArray
	}
	if foundArray == nil {
		return nil, fmt.Errorf("no array found in response")
	}

	rootSchema, ok := foundArray.Schema.(*j5reflect.ObjectField)
	if !ok {
		return nil, fmt.Errorf("expected object schema, got %T", foundArray.Schema)
	}

	out := &client_j5pb.ListRequest{}

	addSearch := func(schema j5reflect.WalkProperty, searching *list_j5pb.SearchingConstraint) {
		if searching == nil {
			return
		}
		out.SearchableFields = append(out.SearchableFields, &client_j5pb.ListRequest_SearchField{
			Name: strings.Join(schema.Path, "."),
		})
	}

	addFilter := func(schema j5reflect.WalkProperty, filtering *list_j5pb.FilteringConstraint) {
		if filtering == nil {
			return
		}
		out.FilterableFields = append(out.FilterableFields, &client_j5pb.ListRequest_FilterField{
			Name:           strings.Join(schema.Path, "."),
			DefaultFilters: filtering.DefaultFilters,
		})
	}

	addSort := func(schema j5reflect.WalkProperty, sorting *list_j5pb.SortingConstraint) {
		if sorting == nil {
			return
		}
		var ds *client_j5pb.ListRequest_SortField_Direction
		if sorting.DefaultSort {
			direction := client_j5pb.ListRequest_SortField_DIRECTION_ASC
			ds = &direction
		}
		out.SortableFields = append(out.SortableFields, &client_j5pb.ListRequest_SortField{
			Name:        strings.Join(schema.Path, "."),
			DefaultSort: ds,
		})
	}

	if err := j5reflect.WalkSchemaFields(rootSchema.Schema(), func(schema j5reflect.WalkProperty) error {

		switch st := schema.Schema.(type) {
		case *j5reflect.EnumField:
			if st.ListRules != nil {
				addFilter(schema, st.ListRules.Filtering)
			}

		case *j5reflect.ScalarSchema:
			switch scalar := st.Proto.Type.(type) {
			case *schema_j5pb.Field_Any:
				// do nothing

			case *schema_j5pb.Field_Array:
				// do nothing

			case *schema_j5pb.Field_Boolean:
				if scalar.Boolean.ListRules != nil {
					addFilter(schema, scalar.Boolean.ListRules.Filtering)
				}

			case *schema_j5pb.Field_Bytes:
				// do nothing

			case *schema_j5pb.Field_Date:
				// do nothing

			case *schema_j5pb.Field_Decimal:
				// do nothing

			case *schema_j5pb.Field_Float:
				if scalar.Float.ListRules != nil {
					addFilter(schema, scalar.Float.ListRules.Filtering)
					addSort(schema, scalar.Float.ListRules.Sorting)
				}

			case *schema_j5pb.Field_Integer:
				if scalar.Integer.ListRules != nil {
					addFilter(schema, scalar.Integer.ListRules.Filtering)
					addSort(schema, scalar.Integer.ListRules.Sorting)
				}

			case *schema_j5pb.Field_Key:
				if scalar.Key.ListRules != nil {
					addFilter(schema, scalar.Key.ListRules.Filtering)
				}

			case *schema_j5pb.Field_Map:
				// do nothing

			case *schema_j5pb.Field_Object:
				// do nothing

			case *schema_j5pb.Field_Oneof:
				if scalar.Oneof.ListRules != nil {
					addFilter(schema, scalar.Oneof.ListRules.Filtering)
				}

			case *schema_j5pb.Field_Timestamp:
				if scalar.Timestamp.ListRules != nil {
					addFilter(schema, scalar.Timestamp.ListRules.Filtering)
					addSort(schema, scalar.Timestamp.ListRules.Sorting)
				}

			case *schema_j5pb.Field_String_:
				if scalar.String_.ListRules != nil {
					addSearch(schema, scalar.String_.ListRules.Searching)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("walk schema fields: %w", err)
	}

	return out, nil
}
