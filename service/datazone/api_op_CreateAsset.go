// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an asset in Amazon DataZone catalog.
func (c *Client) CreateAsset(ctx context.Context, params *CreateAssetInput, optFns ...func(*Options)) (*CreateAssetOutput, error) {
	if params == nil {
		params = &CreateAssetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAsset", params, optFns, c.addOperationCreateAssetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssetInput struct {

	// Amazon DataZone domain where the asset is created.
	//
	// This member is required.
	DomainIdentifier *string

	// Asset name.
	//
	// This member is required.
	Name *string

	// The unique identifier of the project that owns this asset.
	//
	// This member is required.
	OwningProjectIdentifier *string

	// The unique identifier of this asset's type.
	//
	// This member is required.
	TypeIdentifier *string

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// Asset description.
	Description *string

	// The external identifier of the asset.
	ExternalIdentifier *string

	// Metadata forms attached to the asset.
	FormsInput []types.FormInput

	// Glossary terms attached to the asset.
	GlossaryTerms []string

	// The configuration of the automatically generated business-friendly metadata for
	// the asset.
	PredictionConfiguration *types.PredictionConfiguration

	// The revision of this asset's type.
	TypeRevision *string

	noSmithyDocumentSerde
}

type CreateAssetOutput struct {

	// The ID of the Amazon DataZone domain in which the asset was created.
	//
	// This member is required.
	DomainId *string

	// The metadata forms that are attached to the created asset.
	//
	// This member is required.
	FormsOutput []types.FormOutput

	// The unique identifier of the created asset.
	//
	// This member is required.
	Id *string

	// The name of the created asset.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon DataZone project that owns the created asset.
	//
	// This member is required.
	OwningProjectId *string

	// The revision of the asset.
	//
	// This member is required.
	Revision *string

	// The identifier of the created asset type.
	//
	// This member is required.
	TypeIdentifier *string

	// The revision type of the asset.
	//
	// This member is required.
	TypeRevision *string

	// The timestamp of when the asset was created.
	CreatedAt *time.Time

	// The Amazon DataZone user that created this asset in the catalog.
	CreatedBy *string

	// The description of the created asset.
	Description *string

	// The external identifier of the asset.
	ExternalIdentifier *string

	// The timestamp of when the first revision of the asset took place.
	FirstRevisionCreatedAt *time.Time

	// The Amazon DataZone user that made the first revision of the asset.
	FirstRevisionCreatedBy *string

	// The glossary terms that are attached to the created asset.
	GlossaryTerms []string

	// The latest data point that was imported into the time series form for the
	// asset.
	LatestTimeSeriesDataPointFormsOutput []types.TimeSeriesDataPointSummaryFormOutput

	// The details of an asset published in an Amazon DataZone catalog.
	Listing *types.AssetListingDetails

	// The configuration of the automatically generated business-friendly metadata for
	// the asset.
	PredictionConfiguration *types.PredictionConfiguration

	// The read-only metadata forms that are attached to the created asset.
	ReadOnlyFormsOutput []types.FormOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAsset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAsset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAsset"); err != nil {
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
	if err = addIdempotencyToken_opCreateAssetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAssetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAsset(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAsset struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAsset) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAsset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAssetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAssetInput ")
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
func addIdempotencyToken_opCreateAssetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAsset{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAsset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAsset",
	}
}