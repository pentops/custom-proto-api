package j5reflect

import (
	"fmt"

	"github.com/pentops/j5/internal/j5schema"
	"github.com/pentops/j5/internal/patherr"
	"google.golang.org/protobuf/reflect/protoreflect"
)

/*** Interface ***/

type RangeValuesCallback func(Field) error
type RangePropertiesCallback func(Property) error

type Property interface {
	IsSet() bool
	Schema() *j5schema.ObjectProperty
	Field() (Field, error)

	IsArray() bool
	IsMap() bool
}

// PropertySet is implemented by Oneofs, Objects and Maps with String keys.
type PropertySet interface {
	SchemaName() string // Returns the full name of the entity wrapping the properties.

	RangeProperties(RangePropertiesCallback) error
	RangeValues(RangeValuesCallback) error

	// HasProperty returns true if there is a property with the given name in
	// the *schema* for the property set.
	HasProperty(name string) bool

	// GetProperty returns the property with the given name in the schema for
	// the property set. The property may not be set to a value.
	GetProperty(name string) (Property, error)

	GetValue(name string) (Field, bool, error)
	NewValue(name string) (Field, error)
	GetOrCreateValue(name string) (Field, error)

	ListPropertyNames() []string

	implementsPropertySet()
}

type ContainerField interface {
	PropertySet
	Field
}

/*** Implementation ***/

type property struct {
	schema    *j5schema.ObjectProperty
	protoPath []protoreflect.FieldDescriptor

	hasValue bool
	value    Field
}

func (p *property) newEmpty() *property {
	return &property{
		schema:    p.schema,
		protoPath: p.protoPath,
	}
}

var _ Property = &property{}

func (p *property) IsSet() bool {
	return p.hasValue
}

func (p *property) IsArray() bool {
	_, ok := p.schema.Schema.(*j5schema.ArrayField)
	return ok
}

func (p *property) IsMap() bool {
	_, ok := p.schema.Schema.(*j5schema.MapField)
	return ok
}

func (p *property) Schema() *j5schema.ObjectProperty {
	return p.schema
}

func (p *property) Field() (Field, error) {
	if !p.hasValue {
		return nil, fmt.Errorf("field %s is not set", p.schema.JSONName)
	}
	return p.value, nil
}

type propSetFactory struct {
	properties []*property
	fullName   string
}

func (factory propSetFactory) newMessage(msg protoreflect.Message) *propSet {
	ps := &propSet{
		asMap:    map[string]*property{},
		fullName: factory.fullName,
		value:    msg,
	}

	for _, prop := range factory.properties {
		cloned := prop.newEmpty()

		ps.asMap[prop.schema.JSONName] = cloned
		ps.asSlice = append(ps.asSlice, cloned)
	}

	return ps
}

type hasProps interface {
	FullName() string
	ClientProperties() []*j5schema.ObjectProperty
}

var _ PropertySet = &propSet{}

type propSet struct {
	asMap    map[string]*property
	asSlice  []*property
	fullName string

	value protoreflect.Message
}

// ProtoMessage returns the underlying protoreflect message. From there you are
// on your own, the schema may not match.
func (ps *propSet) ProtoReflect() protoreflect.Message {
	return ps.value
}

func newPropSet(schema hasProps, rootDesc protoreflect.MessageDescriptor) (propSetFactory, error) {

	if rootDesc == nil {
		return propSetFactory{}, fmt.Errorf("propSet root is not a message")
	}

	props := schema.ClientProperties()

	builtProps := make([]*property, 0, len(props))
	for _, propSchema := range props {
		prop := &property{
			schema:    propSchema,
			protoPath: make([]protoreflect.FieldDescriptor, len(propSchema.ProtoField)),
		}

		walk := rootDesc
		for idx, fieldNumber := range propSchema.ProtoField {
			fieldDesc := walk.Fields().ByNumber(fieldNumber)
			prop.protoPath[idx] = fieldDesc
			if idx < len(propSchema.ProtoField)-1 {
				if fieldDesc.Kind() != protoreflect.MessageKind {
					return propSetFactory{}, fmt.Errorf("field %s is not a message but has nested types", fieldDesc.FullName())
				}
				walk = fieldDesc.Message()
			}
		}

		builtProps = append(builtProps, prop)
	}

	return propSetFactory{
		fullName:   schema.FullName(),
		properties: builtProps,
	}, nil

}

func (*propSet) implementsPropertySet() {}

func (fs *propSet) SchemaName() string {
	return fs.fullName
}

func (fs *propSet) ListPropertyNames() []string {
	// in order, not using map.
	names := make([]string, 0, len(fs.asSlice))
	for _, prop := range fs.asSlice {
		names = append(names, prop.schema.JSONName)
	}
	return names
}

func (fs *propSet) HasProperty(name string) bool {
	_, ok := fs.asMap[name]
	return ok
}

func (fs *propSet) GetProperty(name string) (Property, error) {
	prop, ok := fs.asMap[name]
	if !ok {
		return nil, fmt.Errorf("%s has no property %s", fs.fullName, name)
	}
	return prop, nil
}

func (fs *propSet) GetOrCreateValue(name string) (Field, error) {
	prop, ok := fs.asMap[name]
	if !ok {
		return nil, fmt.Errorf("%s has no property %s", fs.fullName, name)
	}
	if prop.value != nil {
		return prop.value, nil
	}
	return fs.buildOrCreate(prop)
}

func (fs *propSet) GetValue(name string) (Field, bool, error) {
	prop, ok := fs.asMap[name]
	if !ok {
		return nil, false, fmt.Errorf("%q has no property %q", fs.fullName, name)
	}
	if prop.value != nil {
		return prop.value, true, nil
	}
	return fs.buildValue(prop, false)

}

func (fs *propSet) NewValue(name string) (Field, error) {
	prop, ok := fs.asMap[name]
	if !ok {
		return nil, fmt.Errorf("%q has no property %q", fs.fullName, name)
	}
	if prop.value != nil {
		return prop.value, fmt.Errorf("field %s is already set", name)
	}
	return fs.buildOrCreate(prop)
}

func (fs *propSet) RangeProperties(callback RangePropertiesCallback) error {
	var err error
	for _, prop := range fs.asSlice {
		err = callback(prop)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *propSet) RangeValues(callback RangeValuesCallback) error {
	for _, prop := range fs.asSlice {
		val, has, err := fs.GetValue(prop.schema.JSONName)
		if err != nil {
			return err
		}
		if !has {
			continue
		}
		err = callback(val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *propSet) buildOrCreate(prop *property) (Field, error) {
	val, _, err := fs.buildValue(prop, true)
	return val, err
}

func (fs *propSet) GetOne() (Field, bool, error) {
	var property Field
	var found bool

	for _, search := range fs.asSlice {
		field, has, err := fs.GetValue(search.schema.JSONName)
		if err != nil {
			return nil, false, err
		}
		if !has {
			continue
		}
		if found {
			return nil, true, fmt.Errorf("multiple values set for oneof")
		}
		property = field
		found = true
	}
	return property, found, nil
}

func (fs *propSet) buildValue(prop *property, create bool) (Field, bool, error) {

	msg := fs.value

	fieldContext := &propertyContext{
		schema: prop.schema,
	}

	if len(prop.protoPath) == 0 {
		wrapper, ok := prop.schema.Schema.(*j5schema.OneofField)
		if !ok {
			return nil, false, fmt.Errorf("Reflection Bug: no proto field and not a oneof")
		}

		descriptor := msg.Descriptor()

		propSetFactory, err := newPropSet(wrapper.Schema(), descriptor)
		if err != nil {
			return nil, false, patherr.Wrap(err, prop.schema.JSONName)
		}

		propSet := propSetFactory.newMessage(fs.value)
		oneof := &oneofImpl{
			schema:  wrapper.Schema(),
			propSet: propSet,
		}

		built := newOneofField(fieldContext, wrapper, oneof)
		return built, built.IsSet(), nil
	}

	var walkField protoreflect.FieldDescriptor
	walkMessage := msg
	walkPath := prop.protoPath[:]
	for len(walkPath) > 1 {
		walkField, walkPath = walkPath[0], walkPath[1:]
		fieldContext.walkedProtoPath = append(fieldContext.walkedProtoPath, walkField.JSONName())
		if !create {
			if !walkMessage.Has(walkField) {
				return nil, false, nil
			}
		}
		fieldValue := walkMessage.Mutable(walkField)
		if !fieldValue.IsValid() {
			panic(fmt.Sprintf("Reflection Bug: field %s is not valid", walkField.FullName()))
		}
		walkMessage = fieldValue.Message()
		if walkMessage == nil {
			return nil, false, fmt.Errorf("Reflection Bug: field %s is not a message", walkField.FullName())
		}
	}

	finalField := walkPath[0]
	if !create {
		if !walkMessage.Has(finalField) {
			return nil, false, nil
		}
	}
	fieldContext.walkedProtoPath = append(fieldContext.walkedProtoPath, finalField.JSONName())

	protoVal := newProtoPair(walkMessage, finalField)

	built, err := buildProperty(fieldContext, prop.schema, protoVal)
	if err != nil {
		return nil, false, err
	}
	prop.value = built

	prop.hasValue = true

	return prop.value, true, nil
}

func buildProperty(context fieldContext, schema *j5schema.ObjectProperty, value *protoPair) (Field, error) {

	switch st := schema.Schema.(type) {

	case *j5schema.ArrayField:
		if !value.fieldInParent.IsList() {
			return nil, fmt.Errorf("Reflection Bug: ArrayField is not a list")
		}

		valVal, err := value.getMutableValue(true)
		if err != nil {
			return nil, err
		}
		listVal := valVal.List()

		if st.Schema.Mutable() {

			ff, err := newMessageFieldFactory(st.Schema, value.fieldInParent.Message())
			if err != nil {
				return nil, err
			}

			field, err := newMessageArrayField(context, st, listVal, ff)
			if err != nil {
				return nil, err
			}

			return field, nil
		}

		ff, err := newFieldFactory(st.Schema, value.fieldInParent)
		if err != nil {
			return nil, err
		}

		field, err := newLeafArrayField(context, st, listVal, ff)
		if err != nil {
			return nil, err
		}

		return field, nil

	case *j5schema.MapField:
		if !value.fieldInParent.IsMap() {
			return nil, fmt.Errorf("MapField is not a map")
		}
		valVal, err := value.getMutableValue(true)
		if err != nil {
			return nil, err
		}
		mapVal := valVal.Map()

		if st.Schema.Mutable() {
			ff, err := newMessageFieldFactory(st.Schema, value.fieldInParent.Message())
			if err != nil {
				return nil, err
			}
			field, err := newMessageMapField(context, st, mapVal, ff)
			if err != nil {
				return nil, err
			}
			return field, nil
		}

		ff, err := newFieldFactory(st.Schema, value.fieldInParent.MapValue())
		if err != nil {
			return nil, err
		}

		field, err := newLeafMapField(context, st, mapVal, ff)
		if err != nil {
			return nil, err
		}
		return field, nil
	}

	if schema.Schema.Mutable() {
		messageVal, err := value.getMutableValue(true)
		if err != nil {
			return nil, err
		}
		message := messageVal.Message()
		ff, err := newMessageFieldFactory(schema.Schema, message.Descriptor())
		if err != nil {
			return nil, err
		}
		field := ff.buildField(context, message)
		return field, nil
	}

	ff, err := newFieldFactory(schema.Schema, value.fieldDescriptor())
	if err != nil {
		return nil, err
	}

	field := ff.buildField(context, value)
	return field, nil

}

func copyReflect(a, b protoreflect.Message) {
	a.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		b.Set(fd, val)
		return true
	})
}

type messageFieldFactory interface {
	buildField(schema fieldContext, value protoreflect.Message) Field
}

type fieldFactory interface {
	buildField(schema fieldContext, value protoContext) Field
}

func newMessageFieldFactory(schema j5schema.FieldSchema, desc protoreflect.MessageDescriptor) (messageFieldFactory, error) {

	switch st := schema.(type) {
	case *j5schema.ObjectField:
		propSetFactory, err := newPropSet(st.Schema(), desc)
		if err != nil {
			return nil, err
		}
		return &objectFieldFactory{
			schema:  st,
			propSet: propSetFactory,
		}, nil

	case *j5schema.OneofField:
		propSetFactory, err := newPropSet(st.Schema(), desc)
		if err != nil {
			return nil, err
		}
		return &oneofFieldFactory{
			schema:  st,
			propSet: propSetFactory,
		}, nil

	default:
		panic(fmt.Sprintf("invalid schema for message field: %T", schema))
	}
}

func newFieldFactory(schema j5schema.FieldSchema, field protoreflect.FieldDescriptor) (fieldFactory, error) {
	switch st := schema.(type) {
	case *j5schema.EnumField:
		if field.Kind() != protoreflect.EnumKind {
			return nil, fmt.Errorf("EnumField is kind %s", field.Kind())
		}
		return &enumFieldFactory{schema: st}, nil

	case *j5schema.ScalarSchema:
		if st.WellKnownTypeName != "" {
			if field.Kind() != protoreflect.MessageKind {
				return nil, fmt.Errorf("ScalarField is proto kind %s, want message for %T", field.Kind(), st.Proto.Type)
			}
			if string(field.Message().FullName()) != string(st.WellKnownTypeName) {
				return nil, fmt.Errorf("ScalarField message is %s, want %s for %T", field.Message().FullName(), st.WellKnownTypeName, st.Proto.Type)
			}
		} else if field.Kind() != st.Kind {
			return nil, fmt.Errorf("ScalarField is proto kind %s, want schema %q for %T", field.Kind(), st.Kind, st.Proto.Type)
		}
		return &scalarFieldFactory{schema: st}, nil

	default:
		panic(fmt.Sprintf("invalid schema for leaf field: %T", schema))
	}
}