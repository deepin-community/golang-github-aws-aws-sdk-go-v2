// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a policy statement object. To retrieve a list of existing policy
// statements, use the GetPolicy API.
func (c *Client) AddPolicyStatement(ctx context.Context, params *AddPolicyStatementInput, optFns ...func(*Options)) (*AddPolicyStatementOutput, error) {
	if params == nil {
		params = &AddPolicyStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddPolicyStatement", params, optFns, c.addOperationAddPolicyStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddPolicyStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddPolicyStatementInput struct {

	// The action that the principal can use on the resource.
	//
	// For example, entityresolution:GetIdMappingJob , entityresolution:GetMatchingJob .
	//
	// This member is required.
	Action []string

	// The Amazon Resource Name (ARN) of the resource that will be accessed by the
	// principal.
	//
	// This member is required.
	Arn *string

	// Determines whether the permissions specified in the policy are to be allowed (
	// Allow ) or denied ( Deny ).
	//
	// This member is required.
	Effect types.StatementEffect

	// The Amazon Web Services service or Amazon Web Services account that can access
	// the resource defined as ARN.
	//
	// This member is required.
	Principal []string

	// A statement identifier that differentiates the statement from others in the
	// same policy.
	//
	// This member is required.
	StatementId *string

	// A set of condition keys that you can use in key policies.
	Condition *string

	noSmithyDocumentSerde
}

type AddPolicyStatementOutput struct {

	// The Amazon Resource Name (ARN) of the resource that will be accessed by the
	// principal.
	//
	// This member is required.
	Arn *string

	// A unique identifier for the current revision of the policy.
	//
	// This member is required.
	Token *string

	// The resource-based policy.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddPolicyStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddPolicyStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddPolicyStatement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddPolicyStatement"); err != nil {
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
	if err = addOpAddPolicyStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddPolicyStatement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddPolicyStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddPolicyStatement",
	}
}