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

// Updates the specified user profile in Amazon DataZone.
func (c *Client) UpdateUserProfile(ctx context.Context, params *UpdateUserProfileInput, optFns ...func(*Options)) (*UpdateUserProfileOutput, error) {
	if params == nil {
		params = &UpdateUserProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserProfile", params, optFns, c.addOperationUpdateUserProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserProfileInput struct {

	// The identifier of the Amazon DataZone domain in which a user profile is updated.
	//
	// This member is required.
	DomainIdentifier *string

	// The status of the user profile that are to be updated.
	//
	// This member is required.
	Status types.UserProfileStatus

	// The identifier of the user whose user profile is to be updated.
	//
	// This member is required.
	UserIdentifier *string

	// The type of the user profile that are to be updated.
	Type types.UserProfileType

	noSmithyDocumentSerde
}

type UpdateUserProfileOutput struct {

	// The details of the user profile in Amazon DataZone.
	Details types.UserProfileDetails

	// The identifier of the Amazon DataZone domain in which a user profile is updated.
	DomainId *string

	// The identifier of the user profile.
	Id *string

	// The status of the user profile.
	Status types.UserProfileStatus

	// The type of the user profile.
	Type types.UserProfileType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateUserProfile"); err != nil {
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
	if err = addOpUpdateUserProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUserProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateUserProfile",
	}
}