// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
)

func ExampleImportOptions_outputUsage() {
	var union types.ImportOptions
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ImportOptionsMemberNeptune:
		_ = v.Value // Value is types.NeptuneImportOptions

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.NeptuneImportOptions