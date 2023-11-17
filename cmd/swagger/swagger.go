package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/pentops/custom-proto-api/jsonapi"
	"github.com/pentops/custom-proto-api/structure"
)

func main() {
	src := flag.String("proto-src", "-", "Protobuf binary input file (- for stdin)")
	flag.Parse()

	codecOptions := jsonapi.Options{
		ShortEnums: &jsonapi.ShortEnumsOption{
			UnspecifiedSuffix: "UNSPECIFIED",
			StrictUnmarshal:   true,
		},
		WrapOneof: true,
	}

	descriptors, err := structure.ReadFileDescriptorSet(*src)
	if err != nil {
		log.Fatal(err.Error())
	}

	document, err := structure.BuildFromDescriptors(codecOptions, descriptors)
	if err != nil {
		log.Fatal(err.Error())
	}

	json, err := json.Marshal(document)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(json))

}
