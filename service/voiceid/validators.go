// Code generated by smithy-go-codegen DO NOT EDIT.

package voiceid

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/voiceid/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateFraudster struct {
}

func (*validateOpAssociateFraudster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateFraudster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateFraudsterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateFraudsterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDomain struct {
}

func (*validateOpCreateDomain) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDomainInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateWatchlist struct {
}

func (*validateOpCreateWatchlist) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateWatchlist) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateWatchlistInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateWatchlistInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDomain struct {
}

func (*validateOpDeleteDomain) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDomainInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteFraudster struct {
}

func (*validateOpDeleteFraudster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteFraudster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteFraudsterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteFraudsterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSpeaker struct {
}

func (*validateOpDeleteSpeaker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSpeaker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSpeakerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSpeakerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteWatchlist struct {
}

func (*validateOpDeleteWatchlist) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteWatchlist) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteWatchlistInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteWatchlistInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeDomain struct {
}

func (*validateOpDescribeDomain) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeDomainInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeFraudster struct {
}

func (*validateOpDescribeFraudster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeFraudster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeFraudsterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeFraudsterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeFraudsterRegistrationJob struct {
}

func (*validateOpDescribeFraudsterRegistrationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeFraudsterRegistrationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeFraudsterRegistrationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeFraudsterRegistrationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeSpeakerEnrollmentJob struct {
}

func (*validateOpDescribeSpeakerEnrollmentJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeSpeakerEnrollmentJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeSpeakerEnrollmentJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeSpeakerEnrollmentJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeSpeaker struct {
}

func (*validateOpDescribeSpeaker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeSpeaker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeSpeakerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeSpeakerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeWatchlist struct {
}

func (*validateOpDescribeWatchlist) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeWatchlist) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeWatchlistInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeWatchlistInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateFraudster struct {
}

func (*validateOpDisassociateFraudster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateFraudster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateFraudsterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateFraudsterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEvaluateSession struct {
}

func (*validateOpEvaluateSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEvaluateSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EvaluateSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEvaluateSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListFraudsterRegistrationJobs struct {
}

func (*validateOpListFraudsterRegistrationJobs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListFraudsterRegistrationJobs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListFraudsterRegistrationJobsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListFraudsterRegistrationJobsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListFraudsters struct {
}

func (*validateOpListFraudsters) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListFraudsters) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListFraudstersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListFraudstersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSpeakerEnrollmentJobs struct {
}

func (*validateOpListSpeakerEnrollmentJobs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSpeakerEnrollmentJobs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSpeakerEnrollmentJobsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSpeakerEnrollmentJobsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSpeakers struct {
}

func (*validateOpListSpeakers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSpeakers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSpeakersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSpeakersInput(input); err != nil {
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

type validateOpListWatchlists struct {
}

func (*validateOpListWatchlists) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListWatchlists) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListWatchlistsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListWatchlistsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpOptOutSpeaker struct {
}

func (*validateOpOptOutSpeaker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpOptOutSpeaker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*OptOutSpeakerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpOptOutSpeakerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartFraudsterRegistrationJob struct {
}

func (*validateOpStartFraudsterRegistrationJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartFraudsterRegistrationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartFraudsterRegistrationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartFraudsterRegistrationJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartSpeakerEnrollmentJob struct {
}

func (*validateOpStartSpeakerEnrollmentJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartSpeakerEnrollmentJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartSpeakerEnrollmentJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartSpeakerEnrollmentJobInput(input); err != nil {
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

type validateOpUpdateDomain struct {
}

func (*validateOpUpdateDomain) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDomainInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateWatchlist struct {
}

func (*validateOpUpdateWatchlist) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateWatchlist) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateWatchlistInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateWatchlistInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateFraudsterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateFraudster{}, middleware.After)
}

func addOpCreateDomainValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDomain{}, middleware.After)
}

func addOpCreateWatchlistValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateWatchlist{}, middleware.After)
}

func addOpDeleteDomainValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDomain{}, middleware.After)
}

func addOpDeleteFraudsterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteFraudster{}, middleware.After)
}

func addOpDeleteSpeakerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSpeaker{}, middleware.After)
}

func addOpDeleteWatchlistValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteWatchlist{}, middleware.After)
}

func addOpDescribeDomainValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeDomain{}, middleware.After)
}

func addOpDescribeFraudsterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeFraudster{}, middleware.After)
}

func addOpDescribeFraudsterRegistrationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeFraudsterRegistrationJob{}, middleware.After)
}

func addOpDescribeSpeakerEnrollmentJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeSpeakerEnrollmentJob{}, middleware.After)
}

func addOpDescribeSpeakerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeSpeaker{}, middleware.After)
}

func addOpDescribeWatchlistValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeWatchlist{}, middleware.After)
}

func addOpDisassociateFraudsterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateFraudster{}, middleware.After)
}

func addOpEvaluateSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEvaluateSession{}, middleware.After)
}

func addOpListFraudsterRegistrationJobsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListFraudsterRegistrationJobs{}, middleware.After)
}

func addOpListFraudstersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListFraudsters{}, middleware.After)
}

func addOpListSpeakerEnrollmentJobsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSpeakerEnrollmentJobs{}, middleware.After)
}

func addOpListSpeakersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSpeakers{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpListWatchlistsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListWatchlists{}, middleware.After)
}

func addOpOptOutSpeakerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpOptOutSpeaker{}, middleware.After)
}

func addOpStartFraudsterRegistrationJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartFraudsterRegistrationJob{}, middleware.After)
}

func addOpStartSpeakerEnrollmentJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartSpeakerEnrollmentJob{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateDomainValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDomain{}, middleware.After)
}

func addOpUpdateWatchlistValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateWatchlist{}, middleware.After)
}

func validateInputDataConfig(v *types.InputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InputDataConfig"}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOutputDataConfig(v *types.OutputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OutputDataConfig"}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateServerSideEncryptionConfiguration(v *types.ServerSideEncryptionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ServerSideEncryptionConfiguration"}
	if v.KmsKeyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KmsKeyId"))
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

func validateOpAssociateFraudsterInput(v *AssociateFraudsterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateFraudsterInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.WatchlistId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WatchlistId"))
	}
	if v.FraudsterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FraudsterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDomainInput(v *CreateDomainInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDomainInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ServerSideEncryptionConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServerSideEncryptionConfiguration"))
	} else if v.ServerSideEncryptionConfiguration != nil {
		if err := validateServerSideEncryptionConfiguration(v.ServerSideEncryptionConfiguration); err != nil {
			invalidParams.AddNested("ServerSideEncryptionConfiguration", err.(smithy.InvalidParamsError))
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

func validateOpCreateWatchlistInput(v *CreateWatchlistInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateWatchlistInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDomainInput(v *DeleteDomainInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDomainInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteFraudsterInput(v *DeleteFraudsterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteFraudsterInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.FraudsterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FraudsterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSpeakerInput(v *DeleteSpeakerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSpeakerInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.SpeakerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpeakerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteWatchlistInput(v *DeleteWatchlistInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteWatchlistInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.WatchlistId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WatchlistId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeDomainInput(v *DescribeDomainInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeDomainInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeFraudsterInput(v *DescribeFraudsterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeFraudsterInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.FraudsterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FraudsterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeFraudsterRegistrationJobInput(v *DescribeFraudsterRegistrationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeFraudsterRegistrationJobInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeSpeakerEnrollmentJobInput(v *DescribeSpeakerEnrollmentJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeSpeakerEnrollmentJobInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeSpeakerInput(v *DescribeSpeakerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeSpeakerInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.SpeakerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpeakerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeWatchlistInput(v *DescribeWatchlistInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeWatchlistInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.WatchlistId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WatchlistId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateFraudsterInput(v *DisassociateFraudsterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateFraudsterInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.WatchlistId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WatchlistId"))
	}
	if v.FraudsterId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FraudsterId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEvaluateSessionInput(v *EvaluateSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EvaluateSessionInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.SessionNameOrId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionNameOrId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListFraudsterRegistrationJobsInput(v *ListFraudsterRegistrationJobsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListFraudsterRegistrationJobsInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListFraudstersInput(v *ListFraudstersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListFraudstersInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSpeakerEnrollmentJobsInput(v *ListSpeakerEnrollmentJobsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSpeakerEnrollmentJobsInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSpeakersInput(v *ListSpeakersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSpeakersInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
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

func validateOpListWatchlistsInput(v *ListWatchlistsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListWatchlistsInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpOptOutSpeakerInput(v *OptOutSpeakerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OptOutSpeakerInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.SpeakerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpeakerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartFraudsterRegistrationJobInput(v *StartFraudsterRegistrationJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartFraudsterRegistrationJobInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartSpeakerEnrollmentJobInput(v *StartSpeakerEnrollmentJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartSpeakerEnrollmentJobInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if v.InputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputDataConfig"))
	} else if v.InputDataConfig != nil {
		if err := validateInputDataConfig(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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

func validateOpUpdateDomainInput(v *UpdateDomainInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDomainInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ServerSideEncryptionConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServerSideEncryptionConfiguration"))
	} else if v.ServerSideEncryptionConfiguration != nil {
		if err := validateServerSideEncryptionConfiguration(v.ServerSideEncryptionConfiguration); err != nil {
			invalidParams.AddNested("ServerSideEncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateWatchlistInput(v *UpdateWatchlistInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateWatchlistInput"}
	if v.DomainId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainId"))
	}
	if v.WatchlistId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WatchlistId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}