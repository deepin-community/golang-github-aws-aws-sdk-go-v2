// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates, changes, or deletes CIDR blocks within a collection. Contains
// authoritative IP information mapping blocks to one or multiple locations.
//
// A change request can update multiple locations in a collection at a time, which
// is helpful if you want to move one or more CIDR blocks from one location to
// another in one transaction, without downtime.
//
// # Limits
//
// The max number of CIDR blocks included in the request is 1000. As a result, big
// updates require multiple API calls.
//
// PUT and DELETE_IF_EXISTS
//
// Use ChangeCidrCollection to perform the following actions:
//
//   - PUT : Create a CIDR block within the specified collection.
//
//   - DELETE_IF_EXISTS : Delete an existing CIDR block from the collection.
func (c *Client) ChangeCidrCollection(ctx context.Context, params *ChangeCidrCollectionInput, optFns ...func(*Options)) (*ChangeCidrCollectionOutput, error) {
	if params == nil {
		params = &ChangeCidrCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangeCidrCollection", params, optFns, c.addOperationChangeCidrCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangeCidrCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ChangeCidrCollectionInput struct {

	//  Information about changes to a CIDR collection.
	//
	// This member is required.
	Changes []types.CidrCollectionChange

	// The UUID of the CIDR collection to update.
	//
	// This member is required.
	Id *string

	// A sequential counter that Amazon Route 53 sets to 1 when you create a
	// collection and increments it by 1 each time you update the collection.
	//
	// We recommend that you use ListCidrCollection to get the current value of
	// CollectionVersion for the collection that you want to update, and then include
	// that value with the change request. This prevents Route 53 from overwriting an
	// intervening update:
	//
	//   - If the value in the request matches the value of CollectionVersion in the
	//   collection, Route 53 updates the collection.
	//
	//   - If the value of CollectionVersion in the collection is greater than the
	//   value in the request, the collection was changed after you got the version
	//   number. Route 53 does not update the collection, and it returns a
	//   CidrCollectionVersionMismatch error.
	CollectionVersion *int64

	noSmithyDocumentSerde
}

type ChangeCidrCollectionOutput struct {

	// The ID that is returned by ChangeCidrCollection . You can use it as input to
	// GetChange to see if a CIDR collection change has propagated or not.
	//
	// This member is required.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationChangeCidrCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpChangeCidrCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpChangeCidrCollection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ChangeCidrCollection"); err != nil {
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
	if err = addOpChangeCidrCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChangeCidrCollection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opChangeCidrCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ChangeCidrCollection",
	}
}