// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
)

func ExampleCapabilityConfiguration_outputUsage() {
	var union types.CapabilityConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CapabilityConfigurationMemberEdi:
		_ = v.Value // Value is types.EdiConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EdiConfiguration

func ExampleEdiType_outputUsage() {
	var union types.EdiType
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EdiTypeMemberX12Details:
		_ = v.Value // Value is types.X12Details

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.X12Details