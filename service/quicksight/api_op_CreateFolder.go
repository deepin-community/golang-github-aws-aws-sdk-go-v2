// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an empty shared folder.
func (c *Client) CreateFolder(ctx context.Context, params *CreateFolderInput, optFns ...func(*Options)) (*CreateFolderOutput, error) {
	if params == nil {
		params = &CreateFolderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFolder", params, optFns, c.addOperationCreateFolderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFolderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFolderInput struct {

	// The ID for the Amazon Web Services account where you want to create the folder.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the folder.
	//
	// This member is required.
	FolderId *string

	// The type of folder. By default, folderType is SHARED .
	FolderType types.FolderType

	// The name of the folder.
	Name *string

	// The Amazon Resource Name (ARN) for the parent folder.
	//
	// ParentFolderArn can be null. An empty parentFolderArn creates a root-level
	// folder.
	ParentFolderArn *string

	// A structure that describes the principals and the resource-level permissions of
	// a folder.
	//
	// To specify no permissions, omit Permissions .
	Permissions []types.ResourcePermission

	// An optional parameter that determines the sharing scope of the folder. The
	// default value for this parameter is ACCOUNT .
	SharingModel types.SharingModel

	// Tags for the folder.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFolderOutput struct {

	// The Amazon Resource Name (ARN) for the newly created folder.
	Arn *string

	// The folder ID for the newly created folder.
	FolderId *string

	// The request ID for the newly created folder.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFolderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFolder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFolder{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFolder"); err != nil {
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
	if err = addOpCreateFolderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFolder(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFolder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFolder",
	}
}