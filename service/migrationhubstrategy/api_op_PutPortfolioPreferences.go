// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Saves the specified migration and modernization preferences.
func (c *Client) PutPortfolioPreferences(ctx context.Context, params *PutPortfolioPreferencesInput, optFns ...func(*Options)) (*PutPortfolioPreferencesOutput, error) {
	if params == nil {
		params = &PutPortfolioPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPortfolioPreferences", params, optFns, c.addOperationPutPortfolioPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPortfolioPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPortfolioPreferencesInput struct {

	// The classification for application component types.
	ApplicationMode types.ApplicationMode

	//  The transformation preferences for non-database applications.
	ApplicationPreferences *types.ApplicationPreferences

	//  The transformation preferences for database applications.
	DatabasePreferences *types.DatabasePreferences

	//  The rank of the business goals based on priority.
	PrioritizeBusinessGoals *types.PrioritizeBusinessGoals

	noSmithyDocumentSerde
}

type PutPortfolioPreferencesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPortfolioPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutPortfolioPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutPortfolioPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutPortfolioPreferences"); err != nil {
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
	if err = addOpPutPortfolioPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPortfolioPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPortfolioPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutPortfolioPreferences",
	}
}