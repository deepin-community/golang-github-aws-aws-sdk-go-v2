// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorad/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the group access control entries for a template.
func (c *Client) GetTemplateGroupAccessControlEntry(ctx context.Context, params *GetTemplateGroupAccessControlEntryInput, optFns ...func(*Options)) (*GetTemplateGroupAccessControlEntryOutput, error) {
	if params == nil {
		params = &GetTemplateGroupAccessControlEntryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTemplateGroupAccessControlEntry", params, optFns, c.addOperationGetTemplateGroupAccessControlEntryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTemplateGroupAccessControlEntryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTemplateGroupAccessControlEntryInput struct {

	// Security identifier (SID) of the group object from Active Directory. The SID
	// starts with "S-".
	//
	// This member is required.
	GroupSecurityIdentifier *string

	// The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate].
	//
	// [CreateTemplate]: https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html
	//
	// This member is required.
	TemplateArn *string

	noSmithyDocumentSerde
}

type GetTemplateGroupAccessControlEntryOutput struct {

	// An access control entry allows or denies an Active Directory group from
	// enrolling and/or autoenrolling with a template.
	AccessControlEntry *types.AccessControlEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTemplateGroupAccessControlEntryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTemplateGroupAccessControlEntry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTemplateGroupAccessControlEntry{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTemplateGroupAccessControlEntry"); err != nil {
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
	if err = addOpGetTemplateGroupAccessControlEntryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplateGroupAccessControlEntry(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTemplateGroupAccessControlEntry(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTemplateGroupAccessControlEntry",
	}
}