// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an existing custom vocabulary filter with a new list of words. The new
// list you provide overwrites all previous entries; you cannot append new terms
// onto an existing custom vocabulary filter.
func (c *Client) UpdateVocabularyFilter(ctx context.Context, params *UpdateVocabularyFilterInput, optFns ...func(*Options)) (*UpdateVocabularyFilterOutput, error) {
	if params == nil {
		params = &UpdateVocabularyFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVocabularyFilter", params, optFns, c.addOperationUpdateVocabularyFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVocabularyFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVocabularyFilterInput struct {

	// The name of the custom vocabulary filter you want to update. Custom vocabulary
	// filter names are case sensitive.
	//
	// This member is required.
	VocabularyFilterName *string

	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access
	// the Amazon S3 bucket that contains your input files (in this case, your custom
	// vocabulary filter). If the role that you specify doesn’t have the appropriate
	// permissions to access the specified Amazon S3 location, your request fails.
	//
	// IAM role ARNs have the format
	// arn:partition:iam::account:role/role-name-with-path . For example:
	// arn:aws:iam::111122223333:role/Admin .
	//
	// For more information, see [IAM ARNs].
	//
	// [IAM ARNs]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns
	DataAccessRoleArn *string

	// The Amazon S3 location of the text file that contains your custom vocabulary
	// filter terms. The URI must be located in the same Amazon Web Services Region as
	// the resource you're calling.
	//
	// Here's an example URI path: s3://DOC-EXAMPLE-BUCKET/my-vocab-filter-file.txt
	//
	// Note that if you include VocabularyFilterFileUri in your request, you cannot
	// use Words ; you must choose one or the other.
	VocabularyFilterFileUri *string

	// Use this parameter if you want to update your custom vocabulary filter by
	// including all desired terms, as comma-separated values, within your request. The
	// other option for updating your vocabulary filter is to save your entries in a
	// text file and upload them to an Amazon S3 bucket, then specify the location of
	// your file using the VocabularyFilterFileUri parameter.
	//
	// Note that if you include Words in your request, you cannot use
	// VocabularyFilterFileUri ; you must choose one or the other.
	//
	// Each language has a character set that contains all allowed characters for that
	// specific language. If you use unsupported characters, your custom vocabulary
	// filter request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
	//
	// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
	Words []string

	noSmithyDocumentSerde
}

type UpdateVocabularyFilterOutput struct {

	// The language code you selected for your custom vocabulary filter.
	LanguageCode types.LanguageCode

	// The date and time the specified custom vocabulary filter was last updated.
	//
	// Timestamps are in the format YYYY-MM-DD'T'HH:MM:SS.SSSSSS-UTC . For example,
	// 2022-05-04T12:32:58.761000-07:00 represents 12:32 PM UTC-7 on May 4, 2022.
	LastModifiedTime *time.Time

	// The name of the updated custom vocabulary filter.
	VocabularyFilterName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVocabularyFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateVocabularyFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateVocabularyFilter{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateVocabularyFilter"); err != nil {
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
	if err = addOpUpdateVocabularyFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVocabularyFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVocabularyFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateVocabularyFilter",
	}
}