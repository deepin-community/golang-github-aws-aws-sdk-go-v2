// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a Neptune bulk loader job to load data from an Amazon S3 bucket into a
// Neptune DB instance. See [Using the Amazon Neptune Bulk Loader to Ingest Data].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartLoaderJob]IAM action in that cluster.
//
// [neptune-db:StartLoaderJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startloaderjob
// [Using the Amazon Neptune Bulk Loader to Ingest Data]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load.html
func (c *Client) StartLoaderJob(ctx context.Context, params *StartLoaderJobInput, optFns ...func(*Options)) (*StartLoaderJobOutput, error) {
	if params == nil {
		params = &StartLoaderJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartLoaderJob", params, optFns, c.addOperationStartLoaderJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartLoaderJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartLoaderJobInput struct {

	// The format of the data. For more information about data formats for the Neptune
	// Loader command, see [Load Data Formats].
	//
	// Allowed values
	//
	//   - csv for the [Gremlin CSV data format].
	//
	//   - opencypher for the [openCypher CSV data format].
	//
	//   - ntriples for the [N-Triples RDF data format].
	//
	//   - nquads for the [N-Quads RDF data format].
	//
	//   - rdfxml for the [RDF\XML RDF data format].
	//
	//   - turtle for the [Turtle RDF data format].
	//
	// [RDF\XML RDF data format]: https://www.w3.org/TR/rdf-syntax-grammar/
	// [Gremlin CSV data format]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-gremlin.html
	// [N-Triples RDF data format]: https://www.w3.org/TR/n-triples/
	// [openCypher CSV data format]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-opencypher.html
	// [Turtle RDF data format]: https://www.w3.org/TR/turtle/
	// [Load Data Formats]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format.html
	// [N-Quads RDF data format]: https://www.w3.org/TR/n-quads/
	//
	// This member is required.
	Format types.Format

	// The Amazon Resource Name (ARN) for an IAM role to be assumed by the Neptune DB
	// instance for access to the S3 bucket. The IAM role ARN provided here should be
	// attached to the DB cluster (see [Adding the IAM Role to an Amazon Neptune Cluster].
	//
	// [Adding the IAM Role to an Amazon Neptune Cluster]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-IAM-add-role-cluster.html
	//
	// This member is required.
	IamRoleArn *string

	// The Amazon region of the S3 bucket. This must match the Amazon Region of the DB
	// cluster.
	//
	// This member is required.
	S3BucketRegion types.S3BucketRegion

	// The source parameter accepts an S3 URI that identifies a single file, multiple
	// files, a folder, or multiple folders. Neptune loads every data file in any
	// folder that is specified.
	//
	// The URI can be in any of the following formats.
	//
	//   - s3://(bucket_name)/(object-key-name)
	//
	//   - https://s3.amazonaws.com/(bucket_name)/(object-key-name)
	//
	//   - https://s3.us-east-1.amazonaws.com/(bucket_name)/(object-key-name)
	//
	// The object-key-name element of the URI is equivalent to the [prefix] parameter in an S3 [ListObjects]
	// API call. It identifies all the objects in the specified S3 bucket whose names
	// begin with that prefix. That can be a single file or folder, or multiple files
	// and/or folders.
	//
	// The specified folder or folders can contain multiple vertex files and multiple
	// edge files.
	//
	// [ListObjects]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html
	// [prefix]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html#API_ListObjects_RequestParameters
	//
	// This member is required.
	Source *string

	// This is an optional parameter that can make a queued load request contingent on
	// the successful completion of one or more previous jobs in the queue.
	//
	// Neptune can queue up as many as 64 load requests at a time, if their
	// queueRequest parameters are set to "TRUE" . The dependencies parameter lets you
	// make execution of such a queued request dependent on the successful completion
	// of one or more specified previous requests in the queue.
	//
	// For example, if load Job-A and Job-B are independent of each other, but load
	// Job-C needs Job-A and Job-B to be finished before it begins, proceed as follows:
	//
	//   - Submit load-job-A and load-job-B one after another in any order, and save
	//   their load-ids.
	//
	//   - Submit load-job-C with the load-ids of the two jobs in its dependencies
	//   field:
	//
	// Because of the dependencies parameter, the bulk loader will not start Job-C
	// until Job-A and Job-B have completed successfully. If either one of them fails,
	// Job-C will not be executed, and its status will be set to
	// LOAD_FAILED_BECAUSE_DEPENDENCY_NOT_SATISFIED .
	//
	// You can set up multiple levels of dependency in this way, so that the failure
	// of one job will cause all requests that are directly or indirectly dependent on
	// it to be cancelled.
	Dependencies []string

	// failOnError – A flag to toggle a complete stop on an error.
	//
	// Allowed values: "TRUE" , "FALSE" .
	//
	// Default value: "TRUE" .
	//
	// When this parameter is set to "FALSE" , the loader tries to load all the data in
	// the location specified, skipping any entries with errors.
	//
	// When this parameter is set to "TRUE" , the loader stops as soon as it encounters
	// an error. Data loaded up to that point persists.
	FailOnError *bool

	// The load job mode.
	//
	// Allowed values: RESUME , NEW , AUTO .
	//
	// Default value: AUTO .
	//
	//   - RESUME – In RESUME mode, the loader looks for a previous load from this
	//   source, and if it finds one, resumes that load job. If no previous load job is
	//   found, the loader stops.
	//
	// The loader avoids reloading files that were successfully loaded in a previous
	//   job. It only tries to process failed files. If you dropped previously loaded
	//   data from your Neptune cluster, that data is not reloaded in this mode. If a
	//   previous load job loaded all files from the same source successfully, nothing is
	//   reloaded, and the loader returns success.
	//
	//   - NEW – In NEW mode, the creates a new load request regardless of any previous
	//   loads. You can use this mode to reload all the data from a source after dropping
	//   previously loaded data from your Neptune cluster, or to load new data available
	//   at the same source.
	//
	//   - AUTO – In AUTO mode, the loader looks for a previous load job from the same
	//   source, and if it finds one, resumes that job, just as in RESUME mode.
	//
	// If the loader doesn't find a previous load job from the same source, it loads
	//   all data from the source, just as in NEW mode.
	Mode types.Mode

	// The optional parallelism parameter can be set to reduce the number of threads
	// used by the bulk load process.
	//
	// Allowed values:
	//
	//   - LOW – The number of threads used is the number of available vCPUs divided by
	//   8.
	//
	//   - MEDIUM – The number of threads used is the number of available vCPUs divided
	//   by 2.
	//
	//   - HIGH – The number of threads used is the same as the number of available
	//   vCPUs.
	//
	//   - OVERSUBSCRIBE – The number of threads used is the number of available vCPUs
	//   multiplied by 2. If this value is used, the bulk loader takes up all available
	//   resources.
	//
	// This does not mean, however, that the OVERSUBSCRIBE setting results in 100% CPU
	//   utilization. Because the load operation is I/O bound, the highest CPU
	//   utilization to expect is in the 60% to 70% range.
	//
	// Default value: HIGH
	//
	// The parallelism setting can sometimes result in a deadlock between threads when
	// loading openCypher data. When this happens, Neptune returns the
	// LOAD_DATA_DEADLOCK error. You can generally fix the issue by setting parallelism
	// to a lower setting and retrying the load command.
	Parallelism types.Parallelism

	// parserConfiguration – An optional object with additional parser configuration
	// values. Each of the child parameters is also optional:
	//
	//   - namedGraphUri – The default graph for all RDF formats when no graph is
	//   specified (for non-quads formats and NQUAD entries with no graph).
	//
	// The default is https://aws.amazon.com/neptune/vocab/v01/DefaultNamedGraph .
	//
	//   - baseUri – The base URI for RDF/XML and Turtle formats.
	//
	// The default is https://aws.amazon.com/neptune/default .
	//
	//   - allowEmptyStrings – Gremlin users need to be able to pass empty string
	//   values("") as node and edge properties when loading CSV data. If
	//   allowEmptyStrings is set to false (the default), such empty strings are
	//   treated as nulls and are not loaded.
	//
	// If allowEmptyStrings is set to true , the loader treats empty strings as valid
	//   property values and loads them accordingly.
	ParserConfiguration map[string]string

	// This is an optional flag parameter that indicates whether the load request can
	// be queued up or not.
	//
	// You don't have to wait for one load job to complete before issuing the next
	// one, because Neptune can queue up as many as 64 jobs at a time, provided that
	// their queueRequest parameters are all set to "TRUE" . The queue order of the
	// jobs will be first-in-first-out (FIFO).
	//
	// If the queueRequest parameter is omitted or set to "FALSE" , the load request
	// will fail if another load job is already running.
	//
	// Allowed values: "TRUE" , "FALSE" .
	//
	// Default value: "FALSE" .
	QueueRequest *bool

	// updateSingleCardinalityProperties is an optional parameter that controls how
	// the bulk loader treats a new value for single-cardinality vertex or edge
	// properties. This is not supported for loading openCypher data.
	//
	// Allowed values: "TRUE" , "FALSE" .
	//
	// Default value: "FALSE" .
	//
	// By default, or when updateSingleCardinalityProperties is explicitly set to
	// "FALSE" , the loader treats a new value as an error, because it violates single
	// cardinality.
	//
	// When updateSingleCardinalityProperties is set to "TRUE" , on the other hand, the
	// bulk loader replaces the existing value with the new one. If multiple edge or
	// single-cardinality vertex property values are provided in the source file(s)
	// being loaded, the final value at the end of the bulk load could be any one of
	// those new values. The loader only guarantees that the existing value has been
	// replaced by one of the new ones.
	UpdateSingleCardinalityProperties *bool

	// This parameter is required only when loading openCypher data that contains
	// relationship IDs. It must be included and set to True when openCypher
	// relationship IDs are explicitly provided in the load data (recommended).
	//
	// When userProvidedEdgeIds is absent or set to True , an :ID column must be
	// present in every relationship file in the load.
	//
	// When userProvidedEdgeIds is present and set to False , relationship files in the
	// load must not contain an :ID column. Instead, the Neptune loader automatically
	// generates an ID for each relationship.
	//
	// It's useful to provide relationship IDs explicitly so that the loader can
	// resume loading after error in the CSV data have been fixed, without having to
	// reload any relationships that have already been loaded. If relationship IDs have
	// not been explicitly assigned, the loader cannot resume a failed load if any
	// relationship file has had to be corrected, and must instead reload all the
	// relationships.
	UserProvidedEdgeIds *bool

	noSmithyDocumentSerde
}

type StartLoaderJobOutput struct {

	// Contains a loadId name-value pair that provides an identifier for the load
	// operation.
	//
	// This member is required.
	Payload map[string]string

	// The HTTP return code indicating the status of the load job.
	//
	// This member is required.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartLoaderJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartLoaderJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartLoaderJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartLoaderJob"); err != nil {
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
	if err = addOpStartLoaderJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartLoaderJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartLoaderJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartLoaderJob",
	}
}