// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a running app for the specified UserProfile. This operation is
// automatically invoked by Amazon SageMaker upon access to the associated Domain,
// and when new kernel configurations are selected by the user. A user may have
// multiple Apps active simultaneously.
func (c *Client) CreateApp(ctx context.Context, params *CreateAppInput, optFns ...func(*Options)) (*CreateAppOutput, error) {
	if params == nil {
		params = &CreateAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApp", params, optFns, c.addOperationCreateAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppInput struct {

	// The name of the app.
	//
	// This member is required.
	AppName *string

	// The type of app.
	//
	// This member is required.
	AppType types.AppType

	// The domain ID.
	//
	// This member is required.
	DomainId *string

	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image
	// created on the instance.
	//
	// The value of InstanceType passed as part of the ResourceSpec in the CreateApp
	// call overrides the value passed as part of the ResourceSpec configured for the
	// user profile or the domain. If InstanceType is not specified in any of those
	// three ResourceSpec values for a KernelGateway app, the CreateApp call fails
	// with a request validation error.
	ResourceSpec *types.ResourceSpec

	// The name of the space. If this value is not set, then UserProfileName must be
	// set.
	SpaceName *string

	// Each tag consists of a key and an optional value. Tag keys must be unique per
	// resource.
	Tags []types.Tag

	// The user profile name. If this value is not set, then SpaceName must be set.
	UserProfileName *string

	noSmithyDocumentSerde
}

type CreateAppOutput struct {

	// The Amazon Resource Name (ARN) of the app.
	AppArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateApp{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateApp"); err != nil {
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
	if err = addOpCreateAppValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateApp",
	}
}