// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtains the inline policy assigned to the permission set.
func (c *Client) GetInlinePolicyForPermissionSet(ctx context.Context, params *GetInlinePolicyForPermissionSetInput, optFns ...func(*Options)) (*GetInlinePolicyForPermissionSetOutput, error) {
	if params == nil {
		params = &GetInlinePolicyForPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInlinePolicyForPermissionSet", params, optFns, c.addOperationGetInlinePolicyForPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInlinePolicyForPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInlinePolicyForPermissionSetInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services Service Namespacesin the Amazon Web Services
	// General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the permission set.
	//
	// This member is required.
	PermissionSetArn *string

	noSmithyDocumentSerde
}

type GetInlinePolicyForPermissionSetOutput struct {

	// The inline policy that is attached to the permission set.
	//
	// For Length Constraints , if a valid ARN is provided for a permission set, it is
	// possible for an empty inline policy to be returned.
	InlinePolicy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInlinePolicyForPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInlinePolicyForPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInlinePolicyForPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetInlinePolicyForPermissionSet"); err != nil {
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
	if err = addOpGetInlinePolicyForPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInlinePolicyForPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInlinePolicyForPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetInlinePolicyForPermissionSet",
	}
}