// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
)

func ExampleAggregationRequest_outputUsage() {
	var union types.AggregationRequest
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AggregationRequestMemberAccountAggregation:
		_ = v.Value // Value is types.AccountAggregation

	case *types.AggregationRequestMemberAmiAggregation:
		_ = v.Value // Value is types.AmiAggregation

	case *types.AggregationRequestMemberAwsEcrContainerAggregation:
		_ = v.Value // Value is types.AwsEcrContainerAggregation

	case *types.AggregationRequestMemberEc2InstanceAggregation:
		_ = v.Value // Value is types.Ec2InstanceAggregation

	case *types.AggregationRequestMemberFindingTypeAggregation:
		_ = v.Value // Value is types.FindingTypeAggregation

	case *types.AggregationRequestMemberImageLayerAggregation:
		_ = v.Value // Value is types.ImageLayerAggregation

	case *types.AggregationRequestMemberLambdaFunctionAggregation:
		_ = v.Value // Value is types.LambdaFunctionAggregation

	case *types.AggregationRequestMemberLambdaLayerAggregation:
		_ = v.Value // Value is types.LambdaLayerAggregation

	case *types.AggregationRequestMemberPackageAggregation:
		_ = v.Value // Value is types.PackageAggregation

	case *types.AggregationRequestMemberRepositoryAggregation:
		_ = v.Value // Value is types.RepositoryAggregation

	case *types.AggregationRequestMemberTitleAggregation:
		_ = v.Value // Value is types.TitleAggregation

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.LambdaFunctionAggregation
var _ *types.LambdaLayerAggregation
var _ *types.AwsEcrContainerAggregation
var _ *types.TitleAggregation
var _ *types.ImageLayerAggregation
var _ *types.AmiAggregation
var _ *types.FindingTypeAggregation
var _ *types.AccountAggregation
var _ *types.Ec2InstanceAggregation
var _ *types.PackageAggregation
var _ *types.RepositoryAggregation

func ExampleAggregationResponse_outputUsage() {
	var union types.AggregationResponse
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AggregationResponseMemberAccountAggregation:
		_ = v.Value // Value is types.AccountAggregationResponse

	case *types.AggregationResponseMemberAmiAggregation:
		_ = v.Value // Value is types.AmiAggregationResponse

	case *types.AggregationResponseMemberAwsEcrContainerAggregation:
		_ = v.Value // Value is types.AwsEcrContainerAggregationResponse

	case *types.AggregationResponseMemberEc2InstanceAggregation:
		_ = v.Value // Value is types.Ec2InstanceAggregationResponse

	case *types.AggregationResponseMemberFindingTypeAggregation:
		_ = v.Value // Value is types.FindingTypeAggregationResponse

	case *types.AggregationResponseMemberImageLayerAggregation:
		_ = v.Value // Value is types.ImageLayerAggregationResponse

	case *types.AggregationResponseMemberLambdaFunctionAggregation:
		_ = v.Value // Value is types.LambdaFunctionAggregationResponse

	case *types.AggregationResponseMemberLambdaLayerAggregation:
		_ = v.Value // Value is types.LambdaLayerAggregationResponse

	case *types.AggregationResponseMemberPackageAggregation:
		_ = v.Value // Value is types.PackageAggregationResponse

	case *types.AggregationResponseMemberRepositoryAggregation:
		_ = v.Value // Value is types.RepositoryAggregationResponse

	case *types.AggregationResponseMemberTitleAggregation:
		_ = v.Value // Value is types.TitleAggregationResponse

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AccountAggregationResponse
var _ *types.AmiAggregationResponse
var _ *types.PackageAggregationResponse
var _ *types.RepositoryAggregationResponse
var _ *types.ImageLayerAggregationResponse
var _ *types.LambdaLayerAggregationResponse
var _ *types.TitleAggregationResponse
var _ *types.FindingTypeAggregationResponse
var _ *types.LambdaFunctionAggregationResponse
var _ *types.AwsEcrContainerAggregationResponse
var _ *types.Ec2InstanceAggregationResponse

func ExampleSchedule_outputUsage() {
	var union types.Schedule
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ScheduleMemberDaily:
		_ = v.Value // Value is types.DailySchedule

	case *types.ScheduleMemberMonthly:
		_ = v.Value // Value is types.MonthlySchedule

	case *types.ScheduleMemberOneTime:
		_ = v.Value // Value is types.OneTimeSchedule

	case *types.ScheduleMemberWeekly:
		_ = v.Value // Value is types.WeeklySchedule

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DailySchedule
var _ *types.WeeklySchedule
var _ *types.OneTimeSchedule
var _ *types.MonthlySchedule