package core

import (
	_ "github.com/99designs/gqlgen/graphql/introspection"   // required by protoc
	_ "github.com/gogo/protobuf/gogoproto"                  // required by protoc
	_ "github.com/gogo/protobuf/types"                      // required by protoc
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor" // required by protoc
)
