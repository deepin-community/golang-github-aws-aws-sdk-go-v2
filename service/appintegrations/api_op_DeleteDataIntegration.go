// Code generated by smithy-go-codegen DO NOT EDIT.

package appintegrations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the DataIntegration. Only DataIntegrations that don't have any
// DataIntegrationAssociations can be deleted. Deleting a DataIntegration also
// deletes the underlying Amazon AppFlow flow and service linked role.
//
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the [CreateDataIntegration]API.
//
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func (c *Client) DeleteDataIntegration(ctx context.Context, params *DeleteDataIntegrationInput, optFns ...func(*Options)) (*DeleteDataIntegrationOutput, error) {
	if params == nil {
		params = &DeleteDataIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataIntegration", params, optFns, c.addOperationDeleteDataIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDataIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDataIntegrationInput struct {

	// A unique identifier for the DataIntegration.
	//
	// This member is required.
	DataIntegrationIdentifier *string

	noSmithyDocumentSerde
}

type DeleteDataIntegrationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDataIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteDataIntegration"); err != nil {
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
	if err = addOpDeleteDataIntegrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataIntegration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDataIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteDataIntegration",
	}
}