// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Evaluates a given template and returns the response. The mapping template can
// be a request or response template.
//
// Request templates take the incoming request after a GraphQL operation is parsed
// and convert it into a request configuration for the selected data source
// operation. Response templates interpret responses from the data source and map
// it to the shape of the GraphQL field output type.
//
// Mapping templates are written in the Apache Velocity Template Language (VTL).
func (c *Client) EvaluateMappingTemplate(ctx context.Context, params *EvaluateMappingTemplateInput, optFns ...func(*Options)) (*EvaluateMappingTemplateOutput, error) {
	if params == nil {
		params = &EvaluateMappingTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EvaluateMappingTemplate", params, optFns, c.addOperationEvaluateMappingTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EvaluateMappingTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EvaluateMappingTemplateInput struct {

	// The map that holds all of the contextual information for your resolver
	// invocation. A context is required for this action.
	//
	// This member is required.
	Context *string

	// The mapping template; this can be a request or response template. A template is
	// required for this action.
	//
	// This member is required.
	Template *string

	noSmithyDocumentSerde
}

type EvaluateMappingTemplateOutput struct {

	// The ErrorDetail object.
	Error *types.ErrorDetail

	// The mapping template; this can be a request or response template.
	EvaluationResult *string

	// A list of logs that were generated by calls to util.log.info and util.log.error
	// in the evaluated code.
	Logs []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEvaluateMappingTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEvaluateMappingTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEvaluateMappingTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EvaluateMappingTemplate"); err != nil {
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
	if err = addOpEvaluateMappingTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEvaluateMappingTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEvaluateMappingTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EvaluateMappingTemplate",
	}
}