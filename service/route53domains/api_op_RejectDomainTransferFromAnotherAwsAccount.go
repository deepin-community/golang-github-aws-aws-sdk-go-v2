// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Rejects the transfer of a domain from another Amazon Web Services account to
// the current Amazon Web Services account. You initiate a transfer betweenAmazon
// Web Services accounts using [TransferDomainToAnotherAwsAccount].
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
func (c *Client) RejectDomainTransferFromAnotherAwsAccount(ctx context.Context, params *RejectDomainTransferFromAnotherAwsAccountInput, optFns ...func(*Options)) (*RejectDomainTransferFromAnotherAwsAccountOutput, error) {
	if params == nil {
		params = &RejectDomainTransferFromAnotherAwsAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectDomainTransferFromAnotherAwsAccount", params, optFns, c.addOperationRejectDomainTransferFromAnotherAwsAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectDomainTransferFromAnotherAwsAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The RejectDomainTransferFromAnotherAwsAccount request includes the following
// element.
type RejectDomainTransferFromAnotherAwsAccountInput struct {

	// The name of the domain that was specified when another Amazon Web Services
	// account submitted a [TransferDomainToAnotherAwsAccount]request.
	//
	// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// The RejectDomainTransferFromAnotherAwsAccount response includes the following
// element.
type RejectDomainTransferFromAnotherAwsAccountOutput struct {

	// The identifier that TransferDomainToAnotherAwsAccount returned to track the
	// progress of the request. Because the transfer request was rejected, the value is
	// no longer valid, and you can't use GetOperationDetail to query the operation
	// status.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRejectDomainTransferFromAnotherAwsAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRejectDomainTransferFromAnotherAwsAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRejectDomainTransferFromAnotherAwsAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RejectDomainTransferFromAnotherAwsAccount"); err != nil {
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
	if err = addOpRejectDomainTransferFromAnotherAwsAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRejectDomainTransferFromAnotherAwsAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRejectDomainTransferFromAnotherAwsAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RejectDomainTransferFromAnotherAwsAccount",
	}
}