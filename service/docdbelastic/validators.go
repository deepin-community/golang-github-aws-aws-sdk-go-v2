// Code generated by smithy-go-codegen DO NOT EDIT.

package docdbelastic

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCopyClusterSnapshot struct {
}

func (*validateOpCopyClusterSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCopyClusterSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CopyClusterSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCopyClusterSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCluster struct {
}

func (*validateOpCreateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateClusterSnapshot struct {
}

func (*validateOpCreateClusterSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateClusterSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateClusterSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateClusterSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCluster struct {
}

func (*validateOpDeleteCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteClusterSnapshot struct {
}

func (*validateOpDeleteClusterSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteClusterSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteClusterSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteClusterSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCluster struct {
}

func (*validateOpGetCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetClusterSnapshot struct {
}

func (*validateOpGetClusterSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetClusterSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetClusterSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetClusterSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRestoreClusterFromSnapshot struct {
}

func (*validateOpRestoreClusterFromSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRestoreClusterFromSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RestoreClusterFromSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRestoreClusterFromSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartCluster struct {
}

func (*validateOpStartCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopCluster struct {
}

func (*validateOpStopCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCluster struct {
}

func (*validateOpUpdateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCopyClusterSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCopyClusterSnapshot{}, middleware.After)
}

func addOpCreateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCluster{}, middleware.After)
}

func addOpCreateClusterSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateClusterSnapshot{}, middleware.After)
}

func addOpDeleteClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCluster{}, middleware.After)
}

func addOpDeleteClusterSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteClusterSnapshot{}, middleware.After)
}

func addOpGetClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCluster{}, middleware.After)
}

func addOpGetClusterSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetClusterSnapshot{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpRestoreClusterFromSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRestoreClusterFromSnapshot{}, middleware.After)
}

func addOpStartClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartCluster{}, middleware.After)
}

func addOpStopClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopCluster{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCluster{}, middleware.After)
}

func validateOpCopyClusterSnapshotInput(v *CopyClusterSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CopyClusterSnapshotInput"}
	if v.SnapshotArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotArn"))
	}
	if v.TargetSnapshotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetSnapshotName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateClusterInput(v *CreateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateClusterInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if len(v.AuthType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AuthType"))
	}
	if v.AdminUserName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdminUserName"))
	}
	if v.AdminUserPassword == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdminUserPassword"))
	}
	if v.ShardCapacity == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardCapacity"))
	}
	if v.ShardCount == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardCount"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateClusterSnapshotInput(v *CreateClusterSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateClusterSnapshotInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if v.SnapshotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteClusterInput(v *DeleteClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteClusterInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteClusterSnapshotInput(v *DeleteClusterSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteClusterSnapshotInput"}
	if v.SnapshotArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetClusterInput(v *GetClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetClusterInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetClusterSnapshotInput(v *GetClusterSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetClusterSnapshotInput"}
	if v.SnapshotArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRestoreClusterFromSnapshotInput(v *RestoreClusterFromSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RestoreClusterFromSnapshotInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.SnapshotArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartClusterInput(v *StartClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartClusterInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopClusterInput(v *StopClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopClusterInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateClusterInput(v *UpdateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateClusterInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}