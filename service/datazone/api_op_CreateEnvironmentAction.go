// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an action for the environment, for example, creates a console link for
// an analytics tool that is available in this environment.
func (c *Client) CreateEnvironmentAction(ctx context.Context, params *CreateEnvironmentActionInput, optFns ...func(*Options)) (*CreateEnvironmentActionOutput, error) {
	if params == nil {
		params = &CreateEnvironmentActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentAction", params, optFns, c.addOperationCreateEnvironmentActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentActionInput struct {

	// The ID of the Amazon DataZone domain in which the environment action is created.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the environment in which the environment action is created.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The name of the environment action.
	//
	// This member is required.
	Name *string

	// The parameters of the environment action.
	//
	// This member is required.
	Parameters types.ActionParameters

	// The description of the environment action that is being created in the
	// environment.
	Description *string

	noSmithyDocumentSerde
}

type CreateEnvironmentActionOutput struct {

	// The ID of the domain in which the environment action is created.
	//
	// This member is required.
	DomainId *string

	// The ID of the environment in which the environment is created.
	//
	// This member is required.
	EnvironmentId *string

	// The ID of the environment action.
	//
	// This member is required.
	Id *string

	// The name of the environment action.
	//
	// This member is required.
	Name *string

	// The parameters of the environment action.
	//
	// This member is required.
	Parameters types.ActionParameters

	// The description of the environment action.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEnvironmentAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEnvironmentAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEnvironmentAction"); err != nil {
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
	if err = addOpCreateEnvironmentActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEnvironmentAction",
	}
}