// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a bulk import job request. For more information,
// see [Describe a bulk import job (CLI)]in the Amazon Simple Storage Service User Guide.
//
// [Describe a bulk import job (CLI)]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/DescribeBulkImportJob.html
func (c *Client) DescribeBulkImportJob(ctx context.Context, params *DescribeBulkImportJobInput, optFns ...func(*Options)) (*DescribeBulkImportJobOutput, error) {
	if params == nil {
		params = &DescribeBulkImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBulkImportJob", params, optFns, c.addOperationDescribeBulkImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBulkImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBulkImportJobInput struct {

	// The ID of the job.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeBulkImportJobOutput struct {

	// The Amazon S3 destination where errors associated with the job creation request
	// are saved.
	//
	// This member is required.
	ErrorReportLocation *types.ErrorReportLocation

	// The files in the specified Amazon S3 bucket that contain your data.
	//
	// This member is required.
	Files []types.File

	// Contains the configuration information of a job, such as the file format used
	// to save data in Amazon S3.
	//
	// This member is required.
	JobConfiguration *types.JobConfiguration

	// The date the job was created, in Unix epoch TIME.
	//
	// This member is required.
	JobCreationDate *time.Time

	// The ID of the job.
	//
	// This member is required.
	JobId *string

	// The date the job was last updated, in Unix epoch time.
	//
	// This member is required.
	JobLastUpdateDate *time.Time

	// The unique name that helps identify the job request.
	//
	// This member is required.
	JobName *string

	// The [ARN] of the IAM role that allows IoT SiteWise to read Amazon S3 data.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	JobRoleArn *string

	// The status of the bulk import job can be one of following values:
	//
	//   - PENDING – IoT SiteWise is waiting for the current bulk import job to finish.
	//
	//   - CANCELLED – The bulk import job has been canceled.
	//
	//   - RUNNING – IoT SiteWise is processing your request to import your data from
	//   Amazon S3.
	//
	//   - COMPLETED – IoT SiteWise successfully completed your request to import data
	//   from Amazon S3.
	//
	//   - FAILED – IoT SiteWise couldn't process your request to import data from
	//   Amazon S3. You can use logs saved in the specified error report location in
	//   Amazon S3 to troubleshoot issues.
	//
	//   - COMPLETED_WITH_FAILURES – IoT SiteWise completed your request to import data
	//   from Amazon S3 with errors. You can use logs saved in the specified error report
	//   location in Amazon S3 to troubleshoot issues.
	//
	// This member is required.
	JobStatus types.JobStatus

	// If set to true, ingest new data into IoT SiteWise storage. Measurements with
	// notifications, metrics and transforms are computed. If set to false, historical
	// data is ingested into IoT SiteWise as is.
	AdaptiveIngestion *bool

	// If set to true, your data files is deleted from S3, after ingestion into IoT
	// SiteWise storage.
	DeleteFilesAfterImport *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBulkImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBulkImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBulkImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBulkImportJob"); err != nil {
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
	if err = addEndpointPrefix_opDescribeBulkImportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeBulkImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBulkImportJob(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeBulkImportJobMiddleware struct {
}

func (*endpointPrefix_opDescribeBulkImportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeBulkImportJobMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDescribeBulkImportJobMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDescribeBulkImportJobMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDescribeBulkImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBulkImportJob",
	}
}