// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAcceptInvitation struct {
}

func (*validateOpAcceptInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAcceptInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AcceptInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAcceptInvitationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchGetGraphMemberDatasources struct {
}

func (*validateOpBatchGetGraphMemberDatasources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetGraphMemberDatasources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetGraphMemberDatasourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetGraphMemberDatasourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchGetMembershipDatasources struct {
}

func (*validateOpBatchGetMembershipDatasources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetMembershipDatasources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetMembershipDatasourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetMembershipDatasourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMembers struct {
}

func (*validateOpCreateMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteGraph struct {
}

func (*validateOpDeleteGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMembers struct {
}

func (*validateOpDeleteMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeOrganizationConfiguration struct {
}

func (*validateOpDescribeOrganizationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeOrganizationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeOrganizationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeOrganizationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateMembership struct {
}

func (*validateOpDisassociateMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEnableOrganizationAdminAccount struct {
}

func (*validateOpEnableOrganizationAdminAccount) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEnableOrganizationAdminAccount) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EnableOrganizationAdminAccountInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEnableOrganizationAdminAccountInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetInvestigation struct {
}

func (*validateOpGetInvestigation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetInvestigation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetInvestigationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetInvestigationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMembers struct {
}

func (*validateOpGetMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDatasourcePackages struct {
}

func (*validateOpListDatasourcePackages) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDatasourcePackages) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDatasourcePackagesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDatasourcePackagesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListIndicators struct {
}

func (*validateOpListIndicators) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListIndicators) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListIndicatorsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListIndicatorsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListInvestigations struct {
}

func (*validateOpListInvestigations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListInvestigations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListInvestigationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListInvestigationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListMembers struct {
}

func (*validateOpListMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListMembersInput(input); err != nil {
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

type validateOpRejectInvitation struct {
}

func (*validateOpRejectInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRejectInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RejectInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRejectInvitationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartInvestigation struct {
}

func (*validateOpStartInvestigation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartInvestigation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartInvestigationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartInvestigationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartMonitoringMember struct {
}

func (*validateOpStartMonitoringMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartMonitoringMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartMonitoringMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartMonitoringMemberInput(input); err != nil {
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

type validateOpUpdateDatasourcePackages struct {
}

func (*validateOpUpdateDatasourcePackages) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDatasourcePackages) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDatasourcePackagesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDatasourcePackagesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateInvestigationState struct {
}

func (*validateOpUpdateInvestigationState) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateInvestigationState) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateInvestigationStateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateInvestigationStateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateOrganizationConfiguration struct {
}

func (*validateOpUpdateOrganizationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateOrganizationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateOrganizationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateOrganizationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAcceptInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAcceptInvitation{}, middleware.After)
}

func addOpBatchGetGraphMemberDatasourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetGraphMemberDatasources{}, middleware.After)
}

func addOpBatchGetMembershipDatasourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetMembershipDatasources{}, middleware.After)
}

func addOpCreateMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMembers{}, middleware.After)
}

func addOpDeleteGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteGraph{}, middleware.After)
}

func addOpDeleteMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMembers{}, middleware.After)
}

func addOpDescribeOrganizationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeOrganizationConfiguration{}, middleware.After)
}

func addOpDisassociateMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateMembership{}, middleware.After)
}

func addOpEnableOrganizationAdminAccountValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEnableOrganizationAdminAccount{}, middleware.After)
}

func addOpGetInvestigationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetInvestigation{}, middleware.After)
}

func addOpGetMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMembers{}, middleware.After)
}

func addOpListDatasourcePackagesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDatasourcePackages{}, middleware.After)
}

func addOpListIndicatorsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListIndicators{}, middleware.After)
}

func addOpListInvestigationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListInvestigations{}, middleware.After)
}

func addOpListMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListMembers{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpRejectInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRejectInvitation{}, middleware.After)
}

func addOpStartInvestigationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartInvestigation{}, middleware.After)
}

func addOpStartMonitoringMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartMonitoringMember{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateDatasourcePackagesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDatasourcePackages{}, middleware.After)
}

func addOpUpdateInvestigationStateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateInvestigationState{}, middleware.After)
}

func addOpUpdateOrganizationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateOrganizationConfiguration{}, middleware.After)
}

func validateAccount(v *types.Account) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Account"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if v.EmailAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EmailAddress"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAccountList(v []types.Account) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AccountList"}
	for i := range v {
		if err := validateAccount(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDateFilter(v *types.DateFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DateFilter"}
	if v.StartInclusive == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartInclusive"))
	}
	if v.EndInclusive == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndInclusive"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilterCriteria(v *types.FilterCriteria) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FilterCriteria"}
	if v.Severity != nil {
		if err := validateStringFilter(v.Severity); err != nil {
			invalidParams.AddNested("Severity", err.(smithy.InvalidParamsError))
		}
	}
	if v.Status != nil {
		if err := validateStringFilter(v.Status); err != nil {
			invalidParams.AddNested("Status", err.(smithy.InvalidParamsError))
		}
	}
	if v.State != nil {
		if err := validateStringFilter(v.State); err != nil {
			invalidParams.AddNested("State", err.(smithy.InvalidParamsError))
		}
	}
	if v.EntityArn != nil {
		if err := validateStringFilter(v.EntityArn); err != nil {
			invalidParams.AddNested("EntityArn", err.(smithy.InvalidParamsError))
		}
	}
	if v.CreatedTime != nil {
		if err := validateDateFilter(v.CreatedTime); err != nil {
			invalidParams.AddNested("CreatedTime", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateStringFilter(v *types.StringFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StringFilter"}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAcceptInvitationInput(v *AcceptInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AcceptInvitationInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetGraphMemberDatasourcesInput(v *BatchGetGraphMemberDatasourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetGraphMemberDatasourcesInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.AccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetMembershipDatasourcesInput(v *BatchGetMembershipDatasourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetMembershipDatasourcesInput"}
	if v.GraphArns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArns"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMembersInput(v *CreateMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMembersInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.Accounts == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Accounts"))
	} else if v.Accounts != nil {
		if err := validateAccountList(v.Accounts); err != nil {
			invalidParams.AddNested("Accounts", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteGraphInput(v *DeleteGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteGraphInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMembersInput(v *DeleteMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMembersInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.AccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeOrganizationConfigurationInput(v *DescribeOrganizationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeOrganizationConfigurationInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateMembershipInput(v *DisassociateMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateMembershipInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEnableOrganizationAdminAccountInput(v *EnableOrganizationAdminAccountInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EnableOrganizationAdminAccountInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetInvestigationInput(v *GetInvestigationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetInvestigationInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.InvestigationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvestigationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMembersInput(v *GetMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMembersInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.AccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDatasourcePackagesInput(v *ListDatasourcePackagesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDatasourcePackagesInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListIndicatorsInput(v *ListIndicatorsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListIndicatorsInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.InvestigationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvestigationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListInvestigationsInput(v *ListInvestigationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListInvestigationsInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.FilterCriteria != nil {
		if err := validateFilterCriteria(v.FilterCriteria); err != nil {
			invalidParams.AddNested("FilterCriteria", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListMembersInput(v *ListMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListMembersInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
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

func validateOpRejectInvitationInput(v *RejectInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RejectInvitationInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartInvestigationInput(v *StartInvestigationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartInvestigationInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.EntityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EntityArn"))
	}
	if v.ScopeStartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScopeStartTime"))
	}
	if v.ScopeEndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScopeEndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartMonitoringMemberInput(v *StartMonitoringMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartMonitoringMemberInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
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

func validateOpUpdateDatasourcePackagesInput(v *UpdateDatasourcePackagesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDatasourcePackagesInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.DatasourcePackages == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasourcePackages"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateInvestigationStateInput(v *UpdateInvestigationStateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateInvestigationStateInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.InvestigationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvestigationId"))
	}
	if len(v.State) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("State"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateOrganizationConfigurationInput(v *UpdateOrganizationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateOrganizationConfigurationInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}