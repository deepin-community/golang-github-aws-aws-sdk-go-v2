// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This operation is used with the Amazon GameLift containers feature, which is
//
// currently in public preview.
//
// Retrieves the properties of a container group definition, including all
// container definitions in the group.
//
// To retrieve a container group definition, provide a resource identifier. If
// successful, this operation returns the complete properties of the container
// group definition.
//
// # Learn more
//
// [Manage a container group definition]
//
// [Manage a container group definition]: https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-create-groups.html
func (c *Client) DescribeContainerGroupDefinition(ctx context.Context, params *DescribeContainerGroupDefinitionInput, optFns ...func(*Options)) (*DescribeContainerGroupDefinitionOutput, error) {
	if params == nil {
		params = &DescribeContainerGroupDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeContainerGroupDefinition", params, optFns, c.addOperationDescribeContainerGroupDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeContainerGroupDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeContainerGroupDefinitionInput struct {

	// The unique identifier for the container group definition to retrieve properties
	// for. You can use either the Name or ARN value.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeContainerGroupDefinitionOutput struct {

	// The properties of the requested container group definition resource.
	ContainerGroupDefinition *types.ContainerGroupDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeContainerGroupDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeContainerGroupDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeContainerGroupDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeContainerGroupDefinition"); err != nil {
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
	if err = addOpDescribeContainerGroupDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeContainerGroupDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeContainerGroupDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeContainerGroupDefinition",
	}
}