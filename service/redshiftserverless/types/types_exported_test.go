// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"time"
)

func ExampleSchedule_outputUsage() {
	var union types.Schedule
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ScheduleMemberAt:
		_ = v.Value // Value is time.Time

	case *types.ScheduleMemberCron:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *time.Time

func ExampleTargetAction_outputUsage() {
	var union types.TargetAction
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TargetActionMemberCreateSnapshot:
		_ = v.Value // Value is types.CreateSnapshotScheduleActionParameters

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.CreateSnapshotScheduleActionParameters