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

// Accepts a subscription request to a specific asset.
func (c *Client) AcceptSubscriptionRequest(ctx context.Context, params *AcceptSubscriptionRequestInput, optFns ...func(*Options)) (*AcceptSubscriptionRequestOutput, error) {
	if params == nil {
		params = &AcceptSubscriptionRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptSubscriptionRequest", params, optFns, c.addOperationAcceptSubscriptionRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptSubscriptionRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptSubscriptionRequestInput struct {

	// The Amazon DataZone domain where the specified subscription request is being
	// accepted.
	//
	// This member is required.
	DomainIdentifier *string

	// The unique identifier of the subscription request that is to be accepted.
	//
	// This member is required.
	Identifier *string

	// A description that specifies the reason for accepting the specified
	// subscription request.
	DecisionComment *string

	noSmithyDocumentSerde
}

type AcceptSubscriptionRequestOutput struct {

	// The timestamp that specifies when the subscription request was accepted.
	//
	// This member is required.
	CreatedAt *time.Time

	// Specifies the Amazon DataZone user that accepted the specified subscription
	// request.
	//
	// This member is required.
	CreatedBy *string

	// The unique identifier of the Amazon DataZone domain where the specified
	// subscription request was accepted.
	//
	// This member is required.
	DomainId *string

	// The identifier of the subscription request.
	//
	// This member is required.
	Id *string

	// Specifies the reason for requesting a subscription to the asset.
	//
	// This member is required.
	RequestReason *string

	// Specifies the status of the subscription request.
	//
	// This member is required.
	Status types.SubscriptionRequestStatus

	// Specifies the asset for which the subscription request was created.
	//
	// This member is required.
	SubscribedListings []types.SubscribedListing

	// Specifies the Amazon DataZone users who are subscribed to the asset specified
	// in the subscription request.
	//
	// This member is required.
	SubscribedPrincipals []types.SubscribedPrincipal

	// Specifies the timestamp when subscription request was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// Specifies the reason for accepting the subscription request.
	DecisionComment *string

	// Specifes the ID of the Amazon DataZone user who reviewed the subscription
	// request.
	ReviewerId *string

	// Specifies the Amazon DataZone user who updated the subscription request.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptSubscriptionRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAcceptSubscriptionRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAcceptSubscriptionRequest{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AcceptSubscriptionRequest"); err != nil {
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
	if err = addOpAcceptSubscriptionRequestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptSubscriptionRequest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptSubscriptionRequest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AcceptSubscriptionRequest",
	}
}