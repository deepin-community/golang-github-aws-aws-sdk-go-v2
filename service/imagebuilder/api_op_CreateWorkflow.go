// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new workflow or a new version of an existing workflow.
func (c *Client) CreateWorkflow(ctx context.Context, params *CreateWorkflowInput, optFns ...func(*Options)) (*CreateWorkflowOutput, error) {
	if params == nil {
		params = &CreateWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkflow", params, optFns, c.addOperationCreateWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkflowInput struct {

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see [Ensuring idempotency]in the Amazon EC2 API Reference.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	//
	// This member is required.
	ClientToken *string

	// The name of the workflow to create.
	//
	// This member is required.
	Name *string

	// The semantic version of this workflow resource. The semantic version syntax
	// adheres to the following rules.
	//
	// The semantic version has four nodes: ../. You can assign values for the first
	// three, and can filter on all of them.
	//
	// Assignment: For the first three nodes you can assign any positive integer
	// value, including zero, with an upper limit of 2^30-1, or 1073741823 for each
	// node. Image Builder automatically assigns the build number to the fourth node.
	//
	// Patterns: You can use any numeric pattern that adheres to the assignment
	// requirements for the nodes that you can assign. For example, you might choose a
	// software version pattern, such as 1.0.0, or a date, such as 2021.01.01.
	//
	// This member is required.
	SemanticVersion *string

	// The phase in the image build process for which the workflow resource is
	// responsible.
	//
	// This member is required.
	Type types.WorkflowType

	// Describes what change has been made in this version of the workflow, or what
	// makes this version different from other versions of the workflow.
	ChangeDescription *string

	// Contains the UTF-8 encoded YAML document content for the workflow.
	// Alternatively, you can specify the uri of a YAML document file stored in Amazon
	// S3. However, you cannot specify both properties.
	Data *string

	// Describes the workflow.
	Description *string

	// The ID of the KMS key that is used to encrypt this workflow resource.
	KmsKeyId *string

	// Tags that apply to the workflow resource.
	Tags map[string]string

	// The uri of a YAML component document file. This must be an S3 URL (
	// s3://bucket/key ), and the requester must have permission to access the S3
	// bucket it points to. If you use Amazon S3, you can specify component content up
	// to your service quota.
	//
	// Alternatively, you can specify the YAML document inline, using the component
	// data property. You cannot specify both properties.
	Uri *string

	noSmithyDocumentSerde
}

type CreateWorkflowOutput struct {

	// The client token that uniquely identifies the request.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the workflow resource that the request
	// created.
	WorkflowBuildVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorkflow"); err != nil {
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
	if err = addIdempotencyToken_opCreateWorkflowMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkflow(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWorkflow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWorkflow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWorkflow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWorkflowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWorkflowInput ")
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
func addIdempotencyToken_opCreateWorkflowMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWorkflow{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorkflow",
	}
}