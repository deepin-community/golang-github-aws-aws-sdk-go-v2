// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new, empty repository.
func (c *Client) CreateRepository(ctx context.Context, params *CreateRepositoryInput, optFns ...func(*Options)) (*CreateRepositoryOutput, error) {
	if params == nil {
		params = &CreateRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRepository", params, optFns, c.addOperationCreateRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a create repository operation.
type CreateRepositoryInput struct {

	// The name of the new repository to be created.
	//
	// The repository name must be unique across the calling Amazon Web Services
	// account. Repository names are limited to 100 alphanumeric, dash, and underscore
	// characters, and cannot include certain characters. For more information about
	// the limits on repository names, see [Quotas]in the CodeCommit User Guide. The suffix
	// .git is prohibited.
	//
	// [Quotas]: https://docs.aws.amazon.com/codecommit/latest/userguide/limits.html
	//
	// This member is required.
	RepositoryName *string

	// The ID of the encryption key. You can view the ID of an encryption key in the
	// KMS console, or use the KMS APIs to programmatically retrieve a key ID. For more
	// information about acceptable values for kmsKeyID, see [KeyId]in the Decrypt API
	// description in the Key Management Service API Reference.
	//
	// If no key is specified, the default aws/codecommit Amazon Web Services managed
	// key is used.
	//
	// [KeyId]: https://docs.aws.amazon.com/APIReference/API_Decrypt.html#KMS-Decrypt-request-KeyId
	KmsKeyId *string

	// A comment or description about the new repository.
	//
	// The description field for a repository accepts all HTML characters and all
	// valid Unicode characters. Applications that do not HTML-encode the description
	// and display it in a webpage can expose users to potentially malicious code. Make
	// sure that you HTML-encode the description field in any application that uses
	// this API to display the repository description on a webpage.
	RepositoryDescription *string

	// One or more tag key-value pairs to use when tagging this repository.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents the output of a create repository operation.
type CreateRepositoryOutput struct {

	// Information about the newly created repository.
	RepositoryMetadata *types.RepositoryMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRepository{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRepository"); err != nil {
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
	if err = addOpCreateRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRepository",
	}
}