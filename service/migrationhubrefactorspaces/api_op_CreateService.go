// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubrefactorspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an Amazon Web Services Migration Hub Refactor Spaces service. The
// account owner of the service is always the environment owner, regardless of
// which account in the environment creates the service. Services have either a URL
// endpoint in a virtual private cloud (VPC), or a Lambda function endpoint.
//
// If an Amazon Web Services resource is launched in a service VPC, and you want
// it to be accessible to all of an environment’s services with VPCs and routes,
// apply the RefactorSpacesSecurityGroup to the resource. Alternatively, to add
// more cross-account constraints, apply your own security group.
func (c *Client) CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) {
	if params == nil {
		params = &CreateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateService", params, optFns, c.addOperationCreateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceInput struct {

	// The ID of the application which the service is created.
	//
	// This member is required.
	ApplicationIdentifier *string

	// The type of endpoint to use for the service. The type can be a URL in a VPC or
	// an Lambda function.
	//
	// This member is required.
	EndpointType types.ServiceEndpointType

	// The ID of the environment in which the service is created.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The name of the service.
	//
	// This member is required.
	Name *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The description of the service.
	Description *string

	// The configuration for the Lambda endpoint type.
	LambdaEndpoint *types.LambdaEndpointInput

	// The tags to assign to the service. A tag is a label that you assign to an
	// Amazon Web Services resource. Each tag consists of a key-value pair..
	Tags map[string]string

	// The configuration for the URL endpoint type. When creating a route to a
	// service, Refactor Spaces automatically resolves the address in the
	// UrlEndpointInput object URL when the Domain Name System (DNS) time-to-live (TTL)
	// expires, or every 60 seconds for TTLs less than 60 seconds.
	UrlEndpoint *types.UrlEndpointInput

	// The ID of the VPC.
	VpcId *string

	noSmithyDocumentSerde
}

type CreateServiceOutput struct {

	// The ID of the application that the created service belongs to.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the service.
	Arn *string

	// The Amazon Web Services account ID of the service creator.
	CreatedByAccountId *string

	// A timestamp that indicates when the service is created.
	CreatedTime *time.Time

	// The description of the created service.
	Description *string

	// The endpoint type of the service.
	EndpointType types.ServiceEndpointType

	// The unique identifier of the environment.
	EnvironmentId *string

	// The configuration for the Lambda endpoint type.
	LambdaEndpoint *types.LambdaEndpointInput

	// A timestamp that indicates when the service was last updated.
	LastUpdatedTime *time.Time

	// The name of the service.
	Name *string

	// The Amazon Web Services account ID of the service owner.
	OwnerAccountId *string

	// The unique identifier of the service.
	ServiceId *string

	// The current state of the service.
	State types.ServiceState

	// The tags assigned to the created service. A tag is a label that you assign to
	// an Amazon Web Services resource. Each tag consists of a key-value pair..
	Tags map[string]string

	// The configuration for the URL endpoint type.
	UrlEndpoint *types.UrlEndpointInput

	// The ID of the VPC.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateService{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateService"); err != nil {
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
	if err = addIdempotencyToken_opCreateServiceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateService(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateService struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateService) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateServiceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateService{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateService",
	}
}