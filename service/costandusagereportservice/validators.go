// Code generated by smithy-go-codegen DO NOT EDIT.

package costandusagereportservice

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDeleteReportDefinition struct {
}

func (*validateOpDeleteReportDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteReportDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteReportDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteReportDefinitionInput(input); err != nil {
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

type validateOpModifyReportDefinition struct {
}

func (*validateOpModifyReportDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpModifyReportDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ModifyReportDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpModifyReportDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutReportDefinition struct {
}

func (*validateOpPutReportDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutReportDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutReportDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutReportDefinitionInput(input); err != nil {
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

func addOpDeleteReportDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteReportDefinition{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpModifyReportDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpModifyReportDefinition{}, middleware.After)
}

func addOpPutReportDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutReportDefinition{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateReportDefinition(v *types.ReportDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReportDefinition"}
	if v.ReportName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportName"))
	}
	if len(v.TimeUnit) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("TimeUnit"))
	}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if len(v.Compression) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Compression"))
	}
	if v.AdditionalSchemaElements == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdditionalSchemaElements"))
	}
	if v.S3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Bucket"))
	}
	if v.S3Prefix == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Prefix"))
	}
	if len(v.S3Region) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("S3Region"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteReportDefinitionInput(v *DeleteReportDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteReportDefinitionInput"}
	if v.ReportName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportName"))
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
	if v.ReportName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpModifyReportDefinitionInput(v *ModifyReportDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ModifyReportDefinitionInput"}
	if v.ReportName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportName"))
	}
	if v.ReportDefinition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportDefinition"))
	} else if v.ReportDefinition != nil {
		if err := validateReportDefinition(v.ReportDefinition); err != nil {
			invalidParams.AddNested("ReportDefinition", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutReportDefinitionInput(v *PutReportDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutReportDefinitionInput"}
	if v.ReportDefinition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportDefinition"))
	} else if v.ReportDefinition != nil {
		if err := validateReportDefinition(v.ReportDefinition); err != nil {
			invalidParams.AddNested("ReportDefinition", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
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
	if v.ReportName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportName"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
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
	if v.ReportName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportName"))
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