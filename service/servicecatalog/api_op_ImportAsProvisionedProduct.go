// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Requests the import of a resource as an Service Catalog provisioned product
//
// that is associated to an Service Catalog product and provisioning artifact. Once
// imported, all supported governance actions are supported on the provisioned
// product.
//
// Resource import only supports CloudFormation stack ARNs. CloudFormation
// StackSets, and non-root nested stacks, are not supported.
//
// The CloudFormation stack must have one of the following statuses to be
// imported: CREATE_COMPLETE , UPDATE_COMPLETE , UPDATE_ROLLBACK_COMPLETE ,
// IMPORT_COMPLETE , and IMPORT_ROLLBACK_COMPLETE .
//
// Import of the resource requires that the CloudFormation stack template matches
// the associated Service Catalog product provisioning artifact.
//
// When you import an existing CloudFormation stack into a portfolio, Service
// Catalog does not apply the product's associated constraints during the import
// process. Service Catalog applies the constraints after you call
// UpdateProvisionedProduct for the provisioned product.
//
// The user or role that performs this operation must have the
// cloudformation:GetTemplate and cloudformation:DescribeStacks IAM policy
// permissions.
//
// You can only import one provisioned product at a time. The product's
// CloudFormation stack must have the IMPORT_COMPLETE status before you import
// another.
func (c *Client) ImportAsProvisionedProduct(ctx context.Context, params *ImportAsProvisionedProductInput, optFns ...func(*Options)) (*ImportAsProvisionedProductOutput, error) {
	if params == nil {
		params = &ImportAsProvisionedProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportAsProvisionedProduct", params, optFns, c.addOperationImportAsProvisionedProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportAsProvisionedProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportAsProvisionedProductInput struct {

	// A unique identifier that you provide to ensure idempotency. If multiple
	// requests differ only by the idempotency token, the same response is returned for
	// each repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The unique identifier of the resource to be imported. It only currently
	// supports CloudFormation stack IDs.
	//
	// This member is required.
	PhysicalId *string

	// The product identifier.
	//
	// This member is required.
	ProductId *string

	// The user-friendly name of the provisioned product. The value must be unique for
	// the Amazon Web Services account. The name cannot be updated after the product is
	// provisioned.
	//
	// This member is required.
	ProvisionedProductName *string

	// The identifier of the provisioning artifact.
	//
	// This member is required.
	ProvisioningArtifactId *string

	// The language code.
	//
	//   - jp - Japanese
	//
	//   - zh - Chinese
	AcceptLanguage *string

	noSmithyDocumentSerde
}

type ImportAsProvisionedProductOutput struct {

	// Information about a request operation.
	RecordDetail *types.RecordDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportAsProvisionedProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportAsProvisionedProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportAsProvisionedProduct{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportAsProvisionedProduct"); err != nil {
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
	if err = addIdempotencyToken_opImportAsProvisionedProductMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpImportAsProvisionedProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportAsProvisionedProduct(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpImportAsProvisionedProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpImportAsProvisionedProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpImportAsProvisionedProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ImportAsProvisionedProductInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ImportAsProvisionedProductInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opImportAsProvisionedProductMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpImportAsProvisionedProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opImportAsProvisionedProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportAsProvisionedProduct",
	}
}