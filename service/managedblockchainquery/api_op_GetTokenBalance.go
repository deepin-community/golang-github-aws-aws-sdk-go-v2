// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the balance of a specific token, including native tokens, for a given
// address (wallet or contract) on the blockchain.
//
// Only the native tokens BTC and ETH, and the ERC-20, ERC-721, and ERC 1155 token
// standards are supported.
func (c *Client) GetTokenBalance(ctx context.Context, params *GetTokenBalanceInput, optFns ...func(*Options)) (*GetTokenBalanceOutput, error) {
	if params == nil {
		params = &GetTokenBalanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTokenBalance", params, optFns, c.addOperationGetTokenBalanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTokenBalanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTokenBalanceInput struct {

	// The container for the identifier for the owner.
	//
	// This member is required.
	OwnerIdentifier *types.OwnerIdentifier

	// The container for the identifier for the token, including the unique token ID
	// and its blockchain network.
	//
	// This member is required.
	TokenIdentifier *types.TokenIdentifier

	// The time for when the TokenBalance is requested or the current time if a time
	// is not provided in the request.
	//
	// This time will only be recorded up to the second.
	AtBlockchainInstant *types.BlockchainInstant

	noSmithyDocumentSerde
}

type GetTokenBalanceOutput struct {

	// The container for time.
	//
	// This member is required.
	AtBlockchainInstant *types.BlockchainInstant

	// The container for the token balance.
	//
	// This member is required.
	Balance *string

	// The container for time.
	LastUpdatedTime *types.BlockchainInstant

	// The container for the owner identifier.
	OwnerIdentifier *types.OwnerIdentifier

	// The container for the identifier for the token including the unique token ID
	// and its blockchain network.
	//
	// Only the native tokens BTC and ETH, and the ERC-20, ERC-721, and ERC 1155 token
	// standards are supported.
	TokenIdentifier *types.TokenIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTokenBalanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTokenBalance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTokenBalance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTokenBalance"); err != nil {
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
	if err = addOpGetTokenBalanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTokenBalance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTokenBalance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTokenBalance",
	}
}