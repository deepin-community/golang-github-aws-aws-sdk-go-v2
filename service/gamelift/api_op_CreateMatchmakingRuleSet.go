// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new rule set for FlexMatch matchmaking. A rule set describes the type
// of match to create, such as the number and size of teams. It also sets the
// parameters for acceptable player matches, such as minimum skill level or
// character type.
//
// To create a matchmaking rule set, provide unique rule set name and the rule set
// body in JSON format. Rule sets must be defined in the same Region as the
// matchmaking configuration they are used with.
//
// Since matchmaking rule sets cannot be edited, it is a good idea to check the
// rule set syntax using [ValidateMatchmakingRuleSet]before creating a new rule set.
//
// # Learn more
//
// [Build a rule set]
//
// [Design a matchmaker]
//
// [Matchmaking with FlexMatch]
//
// [Matchmaking with FlexMatch]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-intro.html
// [Build a rule set]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html
// [Design a matchmaker]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html
// [ValidateMatchmakingRuleSet]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_ValidateMatchmakingRuleSet.html
func (c *Client) CreateMatchmakingRuleSet(ctx context.Context, params *CreateMatchmakingRuleSetInput, optFns ...func(*Options)) (*CreateMatchmakingRuleSetOutput, error) {
	if params == nil {
		params = &CreateMatchmakingRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMatchmakingRuleSet", params, optFns, c.addOperationCreateMatchmakingRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMatchmakingRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMatchmakingRuleSetInput struct {

	// A unique identifier for the matchmaking rule set. A matchmaking configuration
	// identifies the rule set it uses by this name value. Note that the rule set name
	// is different from the optional name field in the rule set body.
	//
	// This member is required.
	Name *string

	// A collection of matchmaking rules, formatted as a JSON string. Comments are not
	// allowed in JSON, but most elements support a description field.
	//
	// This member is required.
	RuleSetBody *string

	// A list of labels to assign to the new matchmaking rule set resource. Tags are
	// developer-defined key-value pairs. Tagging Amazon Web Services resources are
	// useful for resource management, access management and cost allocation. For more
	// information, see [Tagging Amazon Web Services Resources]in the Amazon Web Services General Reference.
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMatchmakingRuleSetOutput struct {

	// The newly created matchmaking rule set.
	//
	// This member is required.
	RuleSet *types.MatchmakingRuleSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMatchmakingRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMatchmakingRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMatchmakingRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMatchmakingRuleSet"); err != nil {
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
	if err = addOpCreateMatchmakingRuleSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMatchmakingRuleSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMatchmakingRuleSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMatchmakingRuleSet",
	}
}