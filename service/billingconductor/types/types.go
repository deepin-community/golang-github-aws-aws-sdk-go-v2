// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A representation of a linked account.
type AccountAssociationsListElement struct {

	// The Amazon Web Services account email.
	AccountEmail *string

	// The associating array of account IDs.
	AccountId *string

	// The Amazon Web Services account name.
	AccountName *string

	// The Billing Group Arn that the linked account is associated to.
	BillingGroupArn *string

	noSmithyDocumentSerde
}

// The set of accounts that will be under the billing group. The set of accounts
// resemble the linked accounts in a consolidated billing family.
type AccountGrouping struct {

	// The account IDs that make up the billing group. Account IDs must be a part of
	// the consolidated billing family, and not associated with another billing group.
	//
	// This member is required.
	LinkedAccountIds []string

	// Specifies if this billing group will automatically associate newly added Amazon
	// Web Services accounts that join your consolidated billing family.
	AutoAssociate *bool

	noSmithyDocumentSerde
}

// A representation of a resource association error.
type AssociateResourceError struct {

	// The reason why the resource association failed.
	Message *string

	// A static error code that's used to classify the type of failure.
	Reason AssociateResourceErrorReason

	noSmithyDocumentSerde
}

// A resource association result for a percentage custom line item.
type AssociateResourceResponseElement struct {

	// The resource ARN that was associated to the custom line item.
	Arn *string

	// An AssociateResourceError that will populate if the resource association fails.
	Error *AssociateResourceError

	noSmithyDocumentSerde
}

// The key-value pair that represents the attribute by which the
// BillingGroupCostReportResults are grouped. For example, if you want a
// service-level breakdown for Amazon Simple Storage Service (Amazon S3) of the
// billing group, the attribute will be a key-value pair of "PRODUCT_NAME" and "S3"
// .
type Attribute struct {

	// The key in a key-value pair that describes the margin summary.
	Key *string

	// The value in a key-value pair that describes the margin summary.
	Value *string

	noSmithyDocumentSerde
}

// A summary report of actual Amazon Web Services charges and calculated Amazon
// Web Services charges, based on the associated pricing plan of a billing group.
type BillingGroupCostReportElement struct {

	// The actual Amazon Web Services charges for the billing group.
	AWSCost *string

	// The Amazon Resource Name (ARN) of a billing group.
	Arn *string

	// The displayed currency.
	Currency *string

	// The billing group margin.
	Margin *string

	// The percentage of billing group margin.
	MarginPercentage *string

	// The hypothetical Amazon Web Services charges based on the associated pricing
	// plan of a billing group.
	ProformaCost *string

	noSmithyDocumentSerde
}

// A paginated call to retrieve a list of summary reports of actual Amazon Web
// Services charges and the calculated Amazon Web Services charges, broken down by
// attributes.
type BillingGroupCostReportResultElement struct {

	// The actual Amazon Web Services charges for the billing group.
	AWSCost *string

	// The Amazon Resource Number (ARN) that uniquely identifies the billing group.
	Arn *string

	// The list of key-value pairs that represent the attributes by which the
	// BillingGroupCostReportResults are grouped. For example, if you want the Amazon
	// S3 service-level breakdown of a billing group for November 2023, the attributes
	// list will contain a key-value pair of "PRODUCT_NAME" and "S3" and a key-value
	// pair of "BILLING_PERIOD" and "Nov 2023" .
	Attributes []Attribute

	// The displayed currency.
	Currency *string

	// The billing group margin.
	Margin *string

	// The percentage of the billing group margin.
	MarginPercentage *string

	// The hypothetical Amazon Web Services charges based on the associated pricing
	// plan of a billing group.
	ProformaCost *string

	noSmithyDocumentSerde
}

// A representation of a billing group.
type BillingGroupListElement struct {

	// Specifies if the billing group has automatic account association ( AutoAssociate
	// ) enabled.
	AccountGrouping *ListBillingGroupAccountGrouping

	// The Amazon Resource Number (ARN) that can be used to uniquely identify the
	// billing group.
	Arn *string

	// The preferences and settings that will be used to compute the Amazon Web
	// Services charges for a billing group.
	ComputationPreference *ComputationPreference

	// The time when the billing group was created.
	CreationTime int64

	// The description of the billing group.
	Description *string

	// The most recent time when the billing group was modified.
	LastModifiedTime int64

	// The name of the billing group.
	Name *string

	// The account ID that serves as the main account in a billing group.
	PrimaryAccountId *string

	// The number of accounts in the particular billing group.
	Size int64

	// The billing group status. Only one of the valid values can be used.
	Status BillingGroupStatus

	// The reason why the billing group is in its current status.
	StatusReason *string

	noSmithyDocumentSerde
}

// A time range for which the margin summary is effective. The time range can be
// up to 12 months.
type BillingPeriodRange struct {

	// The exclusive end billing period that defines a billing period range for the
	// margin summary. For example, if you choose a billing period that starts in
	// October 2023 and ends in December 2023, the margin summary will only include
	// data from October 2023 and November 2023.
	//
	// This member is required.
	ExclusiveEndBillingPeriod *string

	// The inclusive start billing period that defines a billing period range for the
	// margin summary.
	//
	// This member is required.
	InclusiveStartBillingPeriod *string

	noSmithyDocumentSerde
}

// The preferences and settings that will be used to compute the Amazon Web
// Services charges for a billing group.
type ComputationPreference struct {

	//  The Amazon Resource Name (ARN) of the pricing plan that's used to compute the
	// Amazon Web Services charges for a billing group.
	//
	// This member is required.
	PricingPlanArn *string

	noSmithyDocumentSerde
}

// The possible Amazon Web Services Free Tier configurations.
type CreateFreeTierConfig struct {

	//  Activate or deactivate Amazon Web Services Free Tier.
	//
	// This member is required.
	Activated *bool

	noSmithyDocumentSerde
}

// The set of tiering configurations for the pricing rule.
type CreateTieringInput struct {

	//  The possible Amazon Web Services Free Tier configurations.
	//
	// This member is required.
	FreeTier *CreateFreeTierConfig

	noSmithyDocumentSerde
}

// The billing period range in which the custom line item request will be applied.
type CustomLineItemBillingPeriodRange struct {

	// The inclusive start billing period that defines a billing period range where a
	// custom line is applied.
	//
	// This member is required.
	InclusiveStartBillingPeriod *string

	// The inclusive end billing period that defines a billing period range where a
	// custom line is applied.
	ExclusiveEndBillingPeriod *string

	noSmithyDocumentSerde
}

// The charge details of a custom line item. It should contain only one of Flat or
// Percentage .
type CustomLineItemChargeDetails struct {

	// The type of the custom line item that indicates whether the charge is a fee or
	// credit.
	//
	// This member is required.
	Type CustomLineItemType

	// A CustomLineItemFlatChargeDetails that describes the charge details of a flat
	// custom line item.
	Flat *CustomLineItemFlatChargeDetails

	// A representation of the line item filter.
	LineItemFilters []LineItemFilter

	// A CustomLineItemPercentageChargeDetails that describes the charge details of a
	// percentage custom line item.
	Percentage *CustomLineItemPercentageChargeDetails

	noSmithyDocumentSerde
}

// A representation of the charge details that are associated with a flat custom
// line item.
type CustomLineItemFlatChargeDetails struct {

	// The custom line item's fixed charge value in USD.
	//
	// This member is required.
	ChargeValue *float64

	noSmithyDocumentSerde
}

// A representation of a custom line item.
type CustomLineItemListElement struct {

	// The Amazon Web Services account in which this custom line item will be applied
	// to.
	AccountId *string

	// The Amazon Resource Names (ARNs) for custom line items.
	Arn *string

	// The number of resources that are associated to the custom line item.
	AssociationSize int64

	// The Amazon Resource Name (ARN) that references the billing group where the
	// custom line item applies to.
	BillingGroupArn *string

	// A ListCustomLineItemChargeDetails that describes the charge details of a custom
	// line item.
	ChargeDetails *ListCustomLineItemChargeDetails

	// The time created.
	CreationTime int64

	// The custom line item's charge value currency. Only one of the valid values can
	// be used.
	CurrencyCode CurrencyCode

	// The custom line item's description. This is shown on the Bills page in
	// association with the charge value.
	Description *string

	// The most recent time when the custom line item was modified.
	LastModifiedTime int64

	// The custom line item's name.
	Name *string

	// The product code that's associated with the custom line item.
	ProductCode *string

	noSmithyDocumentSerde
}

// A representation of the charge details that are associated with a percentage
// custom line item.
type CustomLineItemPercentageChargeDetails struct {

	// The custom line item's percentage value. This will be multiplied against the
	// combined value of its associated resources to determine its charge value.
	//
	// This member is required.
	PercentageValue *float64

	// A list of resource ARNs to associate to the percentage custom line item.
	AssociatedValues []string

	noSmithyDocumentSerde
}

// A representation of a custom line item version.
type CustomLineItemVersionListElement struct {

	// The Amazon Web Services account in which this custom line item will be applied
	// to.
	AccountId *string

	//  A list of custom line item Amazon Resource Names (ARNs) to retrieve
	// information.
	Arn *string

	// The number of resources that are associated with the custom line item.
	AssociationSize int64

	// The Amazon Resource Name (ARN) of the billing group that the custom line item
	// applies to.
	BillingGroupArn *string

	//  A representation of the charge details of a custom line item.
	ChargeDetails *ListCustomLineItemChargeDetails

	// The time when the custom line item version was created.
	CreationTime int64

	// The charge value currency of the custom line item.
	CurrencyCode CurrencyCode

	// The description of the custom line item.
	Description *string

	// The end billing period of the custom line item version.
	EndBillingPeriod *string

	// The most recent time that the custom line item version was modified.
	LastModifiedTime int64

	// The name of the custom line item.
	Name *string

	// The product code that’s associated with the custom line item.
	ProductCode *string

	// The start billing period of the custom line item version.
	StartBillingPeriod *string

	//  The inclusive start time.
	StartTime int64

	noSmithyDocumentSerde
}

// A resource disassociation result for a percentage custom line item.
type DisassociateResourceResponseElement struct {

	// The resource ARN that was disassociated from the custom line item.
	Arn *string

	//  An AssociateResourceError that's shown if the resource disassociation fails.
	Error *AssociateResourceError

	noSmithyDocumentSerde
}

// The possible Amazon Web Services Free Tier configurations.
type FreeTierConfig struct {

	//  Activate or deactivate Amazon Web Services Free Tier application.
	//
	// This member is required.
	Activated *bool

	noSmithyDocumentSerde
}

// A representation of the line item filter for your custom line item. You can use
// line item filters to include or exclude specific resource values from the
// billing group's total cost. For example, if you create a custom line item and
// you want to filter out a value, such as Savings Plan discounts, you can update
// LineItemFilter to exclude it.
type LineItemFilter struct {

	// The attribute of the line item filter. This specifies what attribute that you
	// can filter on.
	//
	// This member is required.
	Attribute LineItemFilterAttributeName

	// The match criteria of the line item filter. This parameter specifies whether
	// not to include the resource value from the billing group total cost.
	//
	// This member is required.
	MatchOption MatchOption

	// The values of the line item filter. This specifies the values to filter on.
	// Currently, you can only exclude Savings Plan discounts.
	//
	// This member is required.
	Values []LineItemFilterValue

	noSmithyDocumentSerde
}

// The filter on the account ID of the linked account, or any of the following:
//
// MONITORED : linked accounts that are associated to billing groups.
//
// UNMONITORED : linked accounts that are not associated to billing groups.
//
// Billing Group Arn : linked accounts that are associated to the provided Billing
// Group Arn.
type ListAccountAssociationsFilter struct {

	// The Amazon Web Services account ID to filter on.
	AccountId *string

	//  The list of Amazon Web Services IDs to retrieve their associated billing group
	// for a given time range.
	AccountIds []string

	// MONITORED : linked accounts that are associated to billing groups.
	//
	// UNMONITORED : linked accounts that are not associated to billing groups.
	//
	// Billing Group Arn : linked accounts that are associated to the provided Billing
	// Group Arn.
	Association *string

	noSmithyDocumentSerde
}

// Specifies if the billing group has the following features enabled.
type ListBillingGroupAccountGrouping struct {

	// Specifies if this billing group will automatically associate newly added Amazon
	// Web Services accounts that join your consolidated billing family.
	AutoAssociate *bool

	noSmithyDocumentSerde
}

// The filter used to retrieve specific BillingGroupCostReportElements .
type ListBillingGroupCostReportsFilter struct {

	// The list of Amazon Resource Names (ARNs) used to filter billing groups to
	// retrieve reports.
	BillingGroupArns []string

	noSmithyDocumentSerde
}

// The filter that specifies the billing groups and pricing plans to retrieve
// billing group information.
type ListBillingGroupsFilter struct {

	// The list of billing group Amazon Resource Names (ARNs) to retrieve information.
	Arns []string

	// Specifies if this billing group will automatically associate newly added Amazon
	// Web Services accounts that join your consolidated billing family.
	AutoAssociate *bool

	// The pricing plan Amazon Resource Names (ARNs) to retrieve information.
	PricingPlan *string

	//  A list of billing groups to retrieve their current status for a specific time
	// range
	Statuses []BillingGroupStatus

	noSmithyDocumentSerde
}

// A representation of the charge details of a custom line item.
type ListCustomLineItemChargeDetails struct {

	//  The type of the custom line item that indicates whether the charge is a fee or
	// credit .
	//
	// This member is required.
	Type CustomLineItemType

	//  A ListCustomLineItemFlatChargeDetails that describes the charge details of a
	// flat custom line item.
	Flat *ListCustomLineItemFlatChargeDetails

	// A representation of the line item filter.
	LineItemFilters []LineItemFilter

	//  A ListCustomLineItemPercentageChargeDetails that describes the charge details
	// of a percentage custom line item.
	Percentage *ListCustomLineItemPercentageChargeDetails

	noSmithyDocumentSerde
}

//	A representation of the charge details that are associated with a flat custom
//
// line item.
type ListCustomLineItemFlatChargeDetails struct {

	//  The custom line item's fixed charge value in USD.
	//
	// This member is required.
	ChargeValue *float64

	noSmithyDocumentSerde
}

//	A representation of the charge details that are associated with a percentage
//
// custom line item.
type ListCustomLineItemPercentageChargeDetails struct {

	//  The custom line item's percentage value. This will be multiplied against the
	// combined value of its associated resources to determine its charge value.
	//
	// This member is required.
	PercentageValue *float64

	noSmithyDocumentSerde
}

// A filter that specifies the custom line items and billing groups to retrieve
// FFLI information.
type ListCustomLineItemsFilter struct {

	// The Amazon Web Services accounts in which this custom line item will be applied
	// to.
	AccountIds []string

	// A list of custom line item ARNs to retrieve information.
	Arns []string

	// The billing group Amazon Resource Names (ARNs) to retrieve information.
	BillingGroups []string

	// A list of custom line items to retrieve information.
	Names []string

	noSmithyDocumentSerde
}

// A billing period filter that specifies the custom line item versions to
// retrieve.
type ListCustomLineItemVersionsBillingPeriodRangeFilter struct {

	// The exclusive end billing period that defines a billing period range where a
	// custom line item version is applied.
	EndBillingPeriod *string

	// The inclusive start billing period that defines a billing period range where a
	// custom line item version is applied.
	StartBillingPeriod *string

	noSmithyDocumentSerde
}

// A filter that specifies the billing period range where the custom line item
// versions reside.
type ListCustomLineItemVersionsFilter struct {

	// The billing period range in which the custom line item version is applied.
	BillingPeriodRange *ListCustomLineItemVersionsBillingPeriodRangeFilter

	noSmithyDocumentSerde
}

// The filter that specifies the Amazon Resource Names (ARNs) of pricing plans, to
// retrieve pricing plan information.
type ListPricingPlansFilter struct {

	// A list of pricing plan Amazon Resource Names (ARNs) to retrieve information.
	Arns []string

	noSmithyDocumentSerde
}

//	The filter that specifies criteria that the pricing rules returned by the
//
// ListPricingRules API will adhere to.
type ListPricingRulesFilter struct {

	// A list containing the pricing rule Amazon Resource Names (ARNs) to include in
	// the API response.
	Arns []string

	noSmithyDocumentSerde
}

//	A filter that specifies the type of resource associations that should be
//
// retrieved for a custom line item.
type ListResourcesAssociatedToCustomLineItemFilter struct {

	//  The type of relationship between the custom line item and the associated
	// resource.
	Relationship CustomLineItemRelationship

	noSmithyDocumentSerde
}

// A representation of a resource association for a custom line item.
type ListResourcesAssociatedToCustomLineItemResponseElement struct {

	//  The ARN of the associated resource.
	Arn *string

	// The end billing period of the associated resource.
	EndBillingPeriod *string

	//  The type of relationship between the custom line item and the associated
	// resource.
	Relationship CustomLineItemRelationship

	noSmithyDocumentSerde
}

// A representation of a pricing plan.
type PricingPlanListElement struct {

	// The pricing plan Amazon Resource Names (ARN). This can be used to uniquely
	// identify a pricing plan.
	Arn *string

	// The time when the pricing plan was created.
	CreationTime int64

	// The pricing plan description.
	Description *string

	// The most recent time when the pricing plan was modified.
	LastModifiedTime int64

	// The name of a pricing plan.
	Name *string

	// The pricing rules count that's currently associated with this pricing plan list
	// element.
	Size int64

	noSmithyDocumentSerde
}

// A representation of a pricing rule.
type PricingRuleListElement struct {

	// The Amazon Resource Name (ARN) used to uniquely identify a pricing rule.
	Arn *string

	// The pricing plans count that this pricing rule is associated with.
	AssociatedPricingPlanCount int64

	//  The seller of services provided by Amazon Web Services, their affiliates, or
	// third-party providers selling services via Amazon Web Services Marketplace.
	BillingEntity *string

	// The time when the pricing rule was created.
	CreationTime int64

	// The pricing rule description.
	Description *string

	//  The most recent time when the pricing rule was modified.
	LastModifiedTime int64

	// A percentage modifier applied on the public pricing rates.
	ModifierPercentage *float64

	// The name of a pricing rule.
	Name *string

	//  Operation is the specific Amazon Web Services action covered by this line
	// item. This describes the specific usage of the line item.
	//
	// If the Scope attribute is set to SKU , this attribute indicates which operation
	// the PricingRule is modifying. For example, a value of RunInstances:0202
	// indicates the operation of running an Amazon EC2 instance.
	Operation *string

	// The scope of pricing rule that indicates if it is globally applicable, or if it
	// is service-specific.
	Scope PricingRuleScope

	// If the Scope attribute is SERVICE , this attribute indicates which service the
	// PricingRule is applicable for.
	Service *string

	//  The set of tiering configurations for the pricing rule.
	Tiering *Tiering

	// The type of pricing rule.
	Type PricingRuleType

	//  Usage type is the unit that each service uses to measure the usage of a
	// specific type of resource.
	//
	// If the Scope attribute is set to SKU , this attribute indicates which usage type
	// the PricingRule is modifying. For example, USW2-BoxUsage:m2.2xlarge describes an
	// M2 High Memory Double Extra Large instance in the US West (Oregon) Region.
	UsageType *string

	noSmithyDocumentSerde
}

// The set of tiering configurations for the pricing rule.
type Tiering struct {

	//  The possible Amazon Web Services Free Tier configurations.
	//
	// This member is required.
	FreeTier *FreeTierConfig

	noSmithyDocumentSerde
}

// Specifies if the billing group has the following features enabled.
type UpdateBillingGroupAccountGrouping struct {

	// Specifies if this billing group will automatically associate newly added Amazon
	// Web Services accounts that join your consolidated billing family.
	AutoAssociate *bool

	noSmithyDocumentSerde
}

//	A representation of the new charge details of a custom line item. This should
//
// contain only one of Flat or Percentage .
type UpdateCustomLineItemChargeDetails struct {

	//  An UpdateCustomLineItemFlatChargeDetails that describes the new charge details
	// of a flat custom line item.
	Flat *UpdateCustomLineItemFlatChargeDetails

	// A representation of the line item filter.
	LineItemFilters []LineItemFilter

	//  An UpdateCustomLineItemPercentageChargeDetails that describes the new charge
	// details of a percentage custom line item.
	Percentage *UpdateCustomLineItemPercentageChargeDetails

	noSmithyDocumentSerde
}

//	A representation of the new charge details that are associated with a flat
//
// custom line item.
type UpdateCustomLineItemFlatChargeDetails struct {

	//  The custom line item's new fixed charge value in USD.
	//
	// This member is required.
	ChargeValue *float64

	noSmithyDocumentSerde
}

//	A representation of the new charge details that are associated with a
//
// percentage custom line item.
type UpdateCustomLineItemPercentageChargeDetails struct {

	//  The custom line item's new percentage value. This will be multiplied against
	// the combined value of its associated resources to determine its charge value.
	//
	// This member is required.
	PercentageValue *float64

	noSmithyDocumentSerde
}

// The possible Amazon Web Services Free Tier configurations.
type UpdateFreeTierConfig struct {

	//  Activate or deactivate application of Amazon Web Services Free Tier.
	//
	// This member is required.
	Activated *bool

	noSmithyDocumentSerde
}

// The set of tiering configurations for the pricing rule.
type UpdateTieringInput struct {

	//  The possible Amazon Web Services Free Tier configurations.
	//
	// This member is required.
	FreeTier *UpdateFreeTierConfig

	noSmithyDocumentSerde
}

// The field's information of a request that resulted in an exception.
type ValidationExceptionField struct {

	// The message describing why the field failed validation.
	//
	// This member is required.
	Message *string

	// The field name.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde