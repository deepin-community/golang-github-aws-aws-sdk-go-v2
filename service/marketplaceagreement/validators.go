// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplaceagreement

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDescribeAgreement struct {
}

func (*validateOpDescribeAgreement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAgreement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAgreementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAgreementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAgreementTerms struct {
}

func (*validateOpGetAgreementTerms) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAgreementTerms) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAgreementTermsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAgreementTermsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDescribeAgreementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAgreement{}, middleware.After)
}

func addOpGetAgreementTermsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAgreementTerms{}, middleware.After)
}

func validateOpDescribeAgreementInput(v *DescribeAgreementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAgreementInput"}
	if v.AgreementId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgreementId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAgreementTermsInput(v *GetAgreementTermsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAgreementTermsInput"}
	if v.AgreementId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgreementId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}