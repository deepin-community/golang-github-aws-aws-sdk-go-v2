// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new backend environment for an Amplify app.
//
// This API is available only to Amplify Gen 1 applications where the backend is
// created using Amplify Studio or the Amplify command line interface (CLI). This
// API isn’t available to Amplify Gen 2 applications. When you deploy an
// application with Amplify Gen 2, you provision the app's backend infrastructure
// using Typescript code.
func (c *Client) CreateBackendEnvironment(ctx context.Context, params *CreateBackendEnvironmentInput, optFns ...func(*Options)) (*CreateBackendEnvironmentOutput, error) {
	if params == nil {
		params = &CreateBackendEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBackendEnvironment", params, optFns, c.addOperationCreateBackendEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBackendEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the backend environment create request.
type CreateBackendEnvironmentInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name for the backend environment.
	//
	// This member is required.
	EnvironmentName *string

	// The name of deployment artifacts.
	DeploymentArtifacts *string

	// The AWS CloudFormation stack name of a backend environment.
	StackName *string

	noSmithyDocumentSerde
}

// The result structure for the create backend environment request.
type CreateBackendEnvironmentOutput struct {

	// Describes the backend environment for an Amplify app.
	//
	// This member is required.
	BackendEnvironment *types.BackendEnvironment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBackendEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBackendEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBackendEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBackendEnvironment"); err != nil {
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
	if err = addOpCreateBackendEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackendEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBackendEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBackendEnvironment",
	}
}