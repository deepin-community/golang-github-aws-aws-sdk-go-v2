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

// Creates an Amazon DataZone environment profile.
func (c *Client) CreateEnvironmentProfile(ctx context.Context, params *CreateEnvironmentProfileInput, optFns ...func(*Options)) (*CreateEnvironmentProfileOutput, error) {
	if params == nil {
		params = &CreateEnvironmentProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentProfile", params, optFns, c.addOperationCreateEnvironmentProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentProfileInput struct {

	// The ID of the Amazon DataZone domain in which this environment profile is
	// created.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the blueprint with which this environment profile is created.
	//
	// This member is required.
	EnvironmentBlueprintIdentifier *string

	// The name of this Amazon DataZone environment profile.
	//
	// This member is required.
	Name *string

	// The identifier of the project in which to create the environment profile.
	//
	// This member is required.
	ProjectIdentifier *string

	// The Amazon Web Services account in which the Amazon DataZone environment is
	// created.
	AwsAccountId *string

	// The Amazon Web Services region in which this environment profile is created.
	AwsAccountRegion *string

	// The description of this Amazon DataZone environment profile.
	Description *string

	// The user parameters of this Amazon DataZone environment profile.
	UserParameters []types.EnvironmentParameter

	noSmithyDocumentSerde
}

type CreateEnvironmentProfileOutput struct {

	// The Amazon DataZone user who created this environment profile.
	//
	// This member is required.
	CreatedBy *string

	// The ID of the Amazon DataZone domain in which this environment profile is
	// created.
	//
	// This member is required.
	DomainId *string

	// The ID of the blueprint with which this environment profile is created.
	//
	// This member is required.
	EnvironmentBlueprintId *string

	// The ID of this Amazon DataZone environment profile.
	//
	// This member is required.
	Id *string

	// The name of this Amazon DataZone environment profile.
	//
	// This member is required.
	Name *string

	// The Amazon Web Services account ID in which this Amazon DataZone environment
	// profile is created.
	AwsAccountId *string

	// The Amazon Web Services region in which this Amazon DataZone environment
	// profile is created.
	AwsAccountRegion *string

	// The timestamp of when this environment profile was created.
	CreatedAt *time.Time

	// The description of this Amazon DataZone environment profile.
	Description *string

	// The ID of the Amazon DataZone project in which this environment profile is
	// created.
	ProjectId *string

	// The timestamp of when this environment profile was updated.
	UpdatedAt *time.Time

	// The user parameters of this Amazon DataZone environment profile.
	UserParameters []types.CustomParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEnvironmentProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEnvironmentProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEnvironmentProfile"); err != nil {
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
	if err = addOpCreateEnvironmentProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEnvironmentProfile",
	}
}