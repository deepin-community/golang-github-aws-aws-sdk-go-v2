// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a presigned URL that you can use to connect to the MLflow UI attached
// to your tracking server. For more information, see [Launch the MLflow UI using a presigned URL].
//
// [Launch the MLflow UI using a presigned URL]: https://docs.aws.amazon.com/sagemaker/latest/dg/mlflow-launch-ui.html
func (c *Client) CreatePresignedMlflowTrackingServerUrl(ctx context.Context, params *CreatePresignedMlflowTrackingServerUrlInput, optFns ...func(*Options)) (*CreatePresignedMlflowTrackingServerUrlOutput, error) {
	if params == nil {
		params = &CreatePresignedMlflowTrackingServerUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePresignedMlflowTrackingServerUrl", params, optFns, c.addOperationCreatePresignedMlflowTrackingServerUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePresignedMlflowTrackingServerUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePresignedMlflowTrackingServerUrlInput struct {

	// The name of the tracking server to connect to your MLflow UI.
	//
	// This member is required.
	TrackingServerName *string

	// The duration in seconds that your presigned URL is valid. The presigned URL can
	// be used only once.
	ExpiresInSeconds *int32

	// The duration in seconds that your MLflow UI session is valid.
	SessionExpirationDurationInSeconds *int32

	noSmithyDocumentSerde
}

type CreatePresignedMlflowTrackingServerUrlOutput struct {

	// A presigned URL with an authorization token.
	AuthorizedUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePresignedMlflowTrackingServerUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePresignedMlflowTrackingServerUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePresignedMlflowTrackingServerUrl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePresignedMlflowTrackingServerUrl"); err != nil {
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
	if err = addOpCreatePresignedMlflowTrackingServerUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePresignedMlflowTrackingServerUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePresignedMlflowTrackingServerUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePresignedMlflowTrackingServerUrl",
	}
}