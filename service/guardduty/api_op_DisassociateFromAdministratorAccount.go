// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates the current GuardDuty member account from its administrator
// account.
//
// When you disassociate an invited member from a GuardDuty delegated
// administrator, the member account details obtained from the [CreateMembers]API, including the
// associated email addresses, are retained. This is done so that the delegated
// administrator can invoke the [InviteMembers]API without the need to invoke the CreateMembers
// API again. To remove the details associated with a member account, the delegated
// administrator must invoke the [DeleteMembers]API.
//
// With autoEnableOrganizationMembers configuration for your organization set to
// ALL , you'll receive an error if you attempt to disable GuardDuty in a member
// account.
//
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [CreateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
func (c *Client) DisassociateFromAdministratorAccount(ctx context.Context, params *DisassociateFromAdministratorAccountInput, optFns ...func(*Options)) (*DisassociateFromAdministratorAccountOutput, error) {
	if params == nil {
		params = &DisassociateFromAdministratorAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateFromAdministratorAccount", params, optFns, c.addOperationDisassociateFromAdministratorAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateFromAdministratorAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateFromAdministratorAccountInput struct {

	// The unique ID of the detector of the GuardDuty member account.
	//
	// This member is required.
	DetectorId *string

	noSmithyDocumentSerde
}

type DisassociateFromAdministratorAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateFromAdministratorAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateFromAdministratorAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateFromAdministratorAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateFromAdministratorAccount"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDisassociateFromAdministratorAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateFromAdministratorAccount(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDisassociateFromAdministratorAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateFromAdministratorAccount",
	}
}