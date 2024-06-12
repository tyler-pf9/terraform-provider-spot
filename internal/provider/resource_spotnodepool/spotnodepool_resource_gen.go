// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_spotnodepool

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/rackerlabs/terraform-provider-spot/internal/spotvalidator"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func SpotnodepoolResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"autoscaling": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"max_nodes": schema.Int64Attribute{
						Optional:            true,
						Description:         "The maximum number of nodes in the node pool.",
						MarkdownDescription: "The maximum number of nodes in the node pool.",
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtLeastSumOf(path.MatchRelative().AtParent().AtName("min_nodes")),
						},
					},
					"min_nodes": schema.Int64Attribute{
						Optional:            true,
						Description:         "The minimum number of nodes in the node pool.",
						MarkdownDescription: "The minimum number of nodes in the node pool.",
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
				CustomType: AutoscalingType{
					ObjectType: types.ObjectType{
						AttrTypes: AutoscalingValue{}.AttributeTypes(ctx),
					},
				},
				Optional:            true,
				Description:         "Scales the nodes in a cluster based on usage. This block should be omitted to disable autoscaling.",
				MarkdownDescription: "Scales the nodes in a cluster based on usage. This block should be omitted to disable autoscaling.",
				Validators: []validator.Object{
					objectvalidator.AlsoRequires(path.MatchRelative().AtName("max_nodes"), path.MatchRelative().AtName("min_nodes")),
				},
			},
			"bid_price": schema.Float64Attribute{
				Required:            true,
				Description:         "The bid price for the server in USD, rounded to three decimal places.",
				MarkdownDescription: "The bid price for the server in USD, rounded to three decimal places.",
				Validators: []validator.Float64{
					float64validator.AtLeast(0.001),
					spotvalidator.DecimalDigitsAtMost(3),
				},
			},
			"bid_status": schema.StringAttribute{
				Computed:            true,
				Description:         "Status of the bid associated with this spotnodepool.",
				MarkdownDescription: "Status of the bid associated with this spotnodepool.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cloudspace_name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the cloudspace.",
				MarkdownDescription: "The name of the cloudspace.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), "Must be valid kubernetes name"),
				},
			},
			"desired_server_count": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "The desired number of servers in the node pool. Should be removed if autoscaling is enabled.",
				MarkdownDescription: "The desired number of servers in the node pool. Should be removed if autoscaling is enabled.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
					int64validator.ConflictsWith(path.MatchRelative().AtParent().AtName("autoscaling")),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The id of the spotnodepool.",
				MarkdownDescription: "The id of the spotnodepool.",
				DeprecationMessage:  "Use the name attribute instead",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed:            true,
				Description:         "The last time the spotnodepool was updated.",
				MarkdownDescription: "The last time the spotnodepool was updated.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Computed:            true,
				Description:         "The name of the spotnodepool.",
				MarkdownDescription: "The name of the spotnodepool.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"server_class": schema.StringAttribute{
				Required:            true,
				Description:         "The class of servers to use for the node pool, can be can be retrieved using the serverclasses data source.",
				MarkdownDescription: "The class of servers to use for the node pool, can be can be retrieved using the serverclasses data source.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"won_count": schema.Int64Attribute{
				Computed:            true,
				Description:         "Number of won bids.",
				MarkdownDescription: "Number of won bids.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

type SpotnodepoolModel struct {
	Autoscaling        AutoscalingValue `tfsdk:"autoscaling"`
	BidPrice           types.Float64    `tfsdk:"bid_price"`
	BidStatus          types.String     `tfsdk:"bid_status"`
	CloudspaceName     types.String     `tfsdk:"cloudspace_name"`
	DesiredServerCount types.Int64      `tfsdk:"desired_server_count"`
	Id                 types.String     `tfsdk:"id"`
	LastUpdated        types.String     `tfsdk:"last_updated"`
	Name               types.String     `tfsdk:"name"`
	ServerClass        types.String     `tfsdk:"server_class"`
	WonCount           types.Int64      `tfsdk:"won_count"`
}

var _ basetypes.ObjectTypable = AutoscalingType{}

type AutoscalingType struct {
	basetypes.ObjectType
}

func (t AutoscalingType) Equal(o attr.Type) bool {
	other, ok := o.(AutoscalingType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t AutoscalingType) String() string {
	return "AutoscalingType"
}

func (t AutoscalingType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	maxNodesAttribute, ok := attributes["max_nodes"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`max_nodes is missing from object`)

		return nil, diags
	}

	maxNodesVal, ok := maxNodesAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`max_nodes expected to be basetypes.Int64Value, was: %T`, maxNodesAttribute))
	}

	minNodesAttribute, ok := attributes["min_nodes"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`min_nodes is missing from object`)

		return nil, diags
	}

	minNodesVal, ok := minNodesAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`min_nodes expected to be basetypes.Int64Value, was: %T`, minNodesAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return AutoscalingValue{
		MaxNodes: maxNodesVal,
		MinNodes: minNodesVal,
		state:    attr.ValueStateKnown,
	}, diags
}

func NewAutoscalingValueNull() AutoscalingValue {
	return AutoscalingValue{
		state: attr.ValueStateNull,
	}
}

func NewAutoscalingValueUnknown() AutoscalingValue {
	return AutoscalingValue{
		state: attr.ValueStateUnknown,
	}
}

func NewAutoscalingValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (AutoscalingValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing AutoscalingValue Attribute Value",
				"While creating a AutoscalingValue value, a missing attribute value was detected. "+
					"A AutoscalingValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("AutoscalingValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid AutoscalingValue Attribute Type",
				"While creating a AutoscalingValue value, an invalid attribute value was detected. "+
					"A AutoscalingValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("AutoscalingValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("AutoscalingValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra AutoscalingValue Attribute Value",
				"While creating a AutoscalingValue value, an extra attribute value was detected. "+
					"A AutoscalingValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra AutoscalingValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewAutoscalingValueUnknown(), diags
	}

	maxNodesAttribute, ok := attributes["max_nodes"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`max_nodes is missing from object`)

		return NewAutoscalingValueUnknown(), diags
	}

	maxNodesVal, ok := maxNodesAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`max_nodes expected to be basetypes.Int64Value, was: %T`, maxNodesAttribute))
	}

	minNodesAttribute, ok := attributes["min_nodes"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`min_nodes is missing from object`)

		return NewAutoscalingValueUnknown(), diags
	}

	minNodesVal, ok := minNodesAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`min_nodes expected to be basetypes.Int64Value, was: %T`, minNodesAttribute))
	}

	if diags.HasError() {
		return NewAutoscalingValueUnknown(), diags
	}

	return AutoscalingValue{
		MaxNodes: maxNodesVal,
		MinNodes: minNodesVal,
		state:    attr.ValueStateKnown,
	}, diags
}

func NewAutoscalingValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) AutoscalingValue {
	object, diags := NewAutoscalingValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewAutoscalingValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t AutoscalingType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewAutoscalingValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewAutoscalingValueUnknown(), nil
	}

	if in.IsNull() {
		return NewAutoscalingValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewAutoscalingValueMust(AutoscalingValue{}.AttributeTypes(ctx), attributes), nil
}

func (t AutoscalingType) ValueType(ctx context.Context) attr.Value {
	return AutoscalingValue{}
}

var _ basetypes.ObjectValuable = AutoscalingValue{}

type AutoscalingValue struct {
	MaxNodes basetypes.Int64Value `tfsdk:"max_nodes"`
	MinNodes basetypes.Int64Value `tfsdk:"min_nodes"`
	state    attr.ValueState
}

func (v AutoscalingValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["max_nodes"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["min_nodes"] = basetypes.Int64Type{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.MaxNodes.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["max_nodes"] = val

		val, err = v.MinNodes.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["min_nodes"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v AutoscalingValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v AutoscalingValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v AutoscalingValue) String() string {
	return "AutoscalingValue"
}

func (v AutoscalingValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"max_nodes": basetypes.Int64Type{},
			"min_nodes": basetypes.Int64Type{},
		},
		map[string]attr.Value{
			"max_nodes": v.MaxNodes,
			"min_nodes": v.MinNodes,
		})

	return objVal, diags
}

func (v AutoscalingValue) Equal(o attr.Value) bool {
	other, ok := o.(AutoscalingValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.MaxNodes.Equal(other.MaxNodes) {
		return false
	}

	if !v.MinNodes.Equal(other.MinNodes) {
		return false
	}

	return true
}

func (v AutoscalingValue) Type(ctx context.Context) attr.Type {
	return AutoscalingType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v AutoscalingValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"max_nodes": basetypes.Int64Type{},
		"min_nodes": basetypes.Int64Type{},
	}
}
