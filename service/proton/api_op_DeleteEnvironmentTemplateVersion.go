// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// If no other minor versions of an environment template exist, delete a major
// version of the environment template if it's not the Recommended version. Delete
// the Recommended version of the environment template if no other major versions
// or minor versions of the environment template exist. A major version of an
// environment template is a version that's not backward compatible.
//
// Delete a minor version of an environment template if it isn't the Recommended
// version. Delete a Recommended minor version of the environment template if no
// other minor versions of the environment template exist. A minor version of an
// environment template is a version that's backward compatible.
func (c *Client) DeleteEnvironmentTemplateVersion(ctx context.Context, params *DeleteEnvironmentTemplateVersionInput, optFns ...func(*Options)) (*DeleteEnvironmentTemplateVersionOutput, error) {
	if params == nil {
		params = &DeleteEnvironmentTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEnvironmentTemplateVersion", params, optFns, c.addOperationDeleteEnvironmentTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEnvironmentTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEnvironmentTemplateVersionInput struct {

	// The environment template major version to delete.
	//
	// This member is required.
	MajorVersion *string

	// The environment template minor version to delete.
	//
	// This member is required.
	MinorVersion *string

	// The name of the environment template.
	//
	// This member is required.
	TemplateName *string

	noSmithyDocumentSerde
}

type DeleteEnvironmentTemplateVersionOutput struct {

	// The detailed data of the environment template version being deleted.
	EnvironmentTemplateVersion *types.EnvironmentTemplateVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEnvironmentTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteEnvironmentTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteEnvironmentTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteEnvironmentTemplateVersion"); err != nil {
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
	if err = addOpDeleteEnvironmentTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEnvironmentTemplateVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEnvironmentTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteEnvironmentTemplateVersion",
	}
}