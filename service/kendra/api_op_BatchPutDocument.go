// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more documents to an index.
//
// The BatchPutDocument API enables you to ingest inline documents or a set of
// documents stored in an Amazon S3 bucket. Use this API to ingest your text and
// unstructured text into an index, add custom attributes to the documents, and to
// attach an access control list to the documents added to the index.
//
// The documents are indexed asynchronously. You can see the progress of the batch
// using Amazon Web Services CloudWatch. Any error messages related to processing
// the batch are sent to your Amazon Web Services CloudWatch log. You can also use
// the BatchGetDocumentStatus API to monitor the progress of indexing your
// documents.
//
// For an example of ingesting inline documents using Python and Java SDKs, see [Adding files directly to an index].
//
// [Adding files directly to an index]: https://docs.aws.amazon.com/kendra/latest/dg/in-adding-binary-doc.html
func (c *Client) BatchPutDocument(ctx context.Context, params *BatchPutDocumentInput, optFns ...func(*Options)) (*BatchPutDocumentOutput, error) {
	if params == nil {
		params = &BatchPutDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutDocument", params, optFns, c.addOperationBatchPutDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutDocumentInput struct {

	// One or more documents to add to the index.
	//
	// Documents have the following file size limits.
	//
	//   - 50 MB total size for any file
	//
	//   - 5 MB extracted text for any file
	//
	// For more information, see [Quotas].
	//
	// [Quotas]: https://docs.aws.amazon.com/kendra/latest/dg/quotas.html
	//
	// This member is required.
	Documents []types.Document

	// The identifier of the index to add the documents to. You need to create the
	// index first using the CreateIndex API.
	//
	// This member is required.
	IndexId *string

	// Configuration information for altering your document metadata and content
	// during the document ingestion process when you use the BatchPutDocument API.
	//
	// For more information on how to create, modify and delete document metadata, or
	// make other content alterations when you ingest documents into Amazon Kendra, see
	// [Customizing document metadata during the ingestion process].
	//
	// [Customizing document metadata during the ingestion process]: https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html
	CustomDocumentEnrichmentConfiguration *types.CustomDocumentEnrichmentConfiguration

	// The Amazon Resource Name (ARN) of an IAM role with permission to access your S3
	// bucket. For more information, see [IAM access roles for Amazon Kendra].
	//
	// [IAM access roles for Amazon Kendra]: https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html
	RoleArn *string

	noSmithyDocumentSerde
}

type BatchPutDocumentOutput struct {

	// A list of documents that were not added to the index because the document
	// failed a validation check. Each document contains an error message that
	// indicates why the document couldn't be added to the index.
	//
	// If there was an error adding a document to an index the error is reported in
	// your Amazon Web Services CloudWatch log. For more information, see [Monitoring Amazon Kendra with Amazon CloudWatch logs].
	//
	// [Monitoring Amazon Kendra with Amazon CloudWatch logs]: https://docs.aws.amazon.com/kendra/latest/dg/cloudwatch-logs.html
	FailedDocuments []types.BatchPutDocumentResponseFailedDocument

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchPutDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchPutDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchPutDocument{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchPutDocument"); err != nil {
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
	if err = addOpBatchPutDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutDocument(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchPutDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchPutDocument",
	}
}