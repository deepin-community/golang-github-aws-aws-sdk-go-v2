// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
)

func ExampleProviderEndpointConfiguration_outputUsage() {
	var union types.ProviderEndpointConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ProviderEndpointConfigurationMemberMarketplaceConfiguration:
		_ = v.Value // Value is types.ProviderMarketplaceConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ProviderMarketplaceConfiguration