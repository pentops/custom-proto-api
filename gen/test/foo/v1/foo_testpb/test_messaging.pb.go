// Code generated by protoc-gen-go-messaging. DO NOT EDIT.

package foo_testpb

// Service: FooService
// Method: GetFoo

func (msg *GetFooRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/test.v1.FooService/GetFoo",
		"grpc-message": "test.v1.GetFooRequest",
		"api-version":  "9ced96f272c9411019144324060b3be973b321b4",
	}
	return headers
}

// Method: PostFoo

func (msg *PostFooRequest) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/test.v1.FooService/PostFoo",
		"grpc-message": "test.v1.PostFooRequest",
		"api-version":  "9ced96f272c9411019144324060b3be973b321b4",
	}
	return headers
}

// Service: FooTopic
// Method: Foo

func (msg *FooMessage) MessagingHeaders() map[string]string {
	headers := map[string]string{
		"grpc-service": "/test.v1.FooTopic/Foo",
		"grpc-message": "test.v1.FooMessage",
		"api-version":  "9ced96f272c9411019144324060b3be973b321b4",
	}
	return headers
}
