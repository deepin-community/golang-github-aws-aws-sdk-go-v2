// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details about the Hub resource in your account, including the HubArn
// and the time when you enabled Security Hub.
func (c *Client) DescribeHub(ctx context.Context, params *DescribeHubInput, optFns ...func(*Options)) (*DescribeHubOutput, error) {
	if params == nil {
		params = &DescribeHubInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHub", params, optFns, c.addOperationDescribeHubMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHubOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHubInput struct {

	// The ARN of the Hub resource to retrieve.
	HubArn *string

	noSmithyDocumentSerde
}

type DescribeHubOutput struct {

	// Whether to automatically enable new controls when they are added to standards
	// that are enabled.
	//
	// If set to true , then new controls for enabled standards are enabled
	// automatically. If set to false , then new controls are not enabled.
	AutoEnableControls *bool

	// Specifies whether the calling account has consolidated control findings turned
	// on. If the value for this field is set to SECURITY_CONTROL , Security Hub
	// generates a single finding for a control check even when the check applies to
	// multiple enabled standards.
	//
	// If the value for this field is set to STANDARD_CONTROL , Security Hub generates
	// separate findings for a control check when the check applies to multiple enabled
	// standards.
	//
	// The value for this field in a member account matches the value in the
	// administrator account. For accounts that aren't part of an organization, the
	// default value of this field is SECURITY_CONTROL if you enabled Security Hub on
	// or after February 23, 2023.
	ControlFindingGenerator types.ControlFindingGenerator

	// The ARN of the Hub resource that was retrieved.
	HubArn *string

	// The date and time when Security Hub was enabled in the account.
	SubscribedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHubMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeHub{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeHub{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeHub"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHub(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHub(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeHub",
	}
}