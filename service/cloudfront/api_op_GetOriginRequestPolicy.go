// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets an origin request policy, including the following metadata:
//
//   - The policy's identifier.
//
//   - The date and time when the policy was last modified.
//
// To get an origin request policy, you must provide the policy's identifier. If
// the origin request policy is attached to a distribution's cache behavior, you
// can get the policy's identifier using ListDistributions or GetDistribution . If
// the origin request policy is not attached to a cache behavior, you can get the
// identifier using ListOriginRequestPolicies .
func (c *Client) GetOriginRequestPolicy(ctx context.Context, params *GetOriginRequestPolicyInput, optFns ...func(*Options)) (*GetOriginRequestPolicyOutput, error) {
	if params == nil {
		params = &GetOriginRequestPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOriginRequestPolicy", params, optFns, c.addOperationGetOriginRequestPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOriginRequestPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOriginRequestPolicyInput struct {

	// The unique identifier for the origin request policy. If the origin request
	// policy is attached to a distribution's cache behavior, you can get the policy's
	// identifier using ListDistributions or GetDistribution . If the origin request
	// policy is not attached to a cache behavior, you can get the identifier using
	// ListOriginRequestPolicies .
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetOriginRequestPolicyOutput struct {

	// The current version of the origin request policy.
	ETag *string

	// The origin request policy.
	OriginRequestPolicy *types.OriginRequestPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOriginRequestPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetOriginRequestPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetOriginRequestPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOriginRequestPolicy"); err != nil {
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
	if err = addOpGetOriginRequestPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOriginRequestPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOriginRequestPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOriginRequestPolicy",
	}
}