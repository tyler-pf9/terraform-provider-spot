// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_cloudspace

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func CloudspaceResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bids": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bid_name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"won_count": schema.Int64Attribute{
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
					},
					CustomType: BidsType{
						ObjectType: types.ObjectType{
							AttrTypes: BidsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"cloudspace_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The name of the cloudspace.",
				MarkdownDescription: "The name of the cloudspace.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), "Must be a valid kubernetes name"),
				},
			},
			"cni": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Container Network Interface (CNI) to use. Supported values: calico, cilium, byocni",
				MarkdownDescription: "Container Network Interface (CNI) to use. Supported values: calico, cilium, byocni",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("calico", "cilium", "byocni"),
				},
				Default: stringdefault.StaticString("calico"),
			},
			"deployment_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Specifies the deployment type for the cloudspace (Only gen2 is allowed value).",
				MarkdownDescription: "Specifies the deployment type for the cloudspace (Only gen2 is allowed value).",
				DeprecationMessage:  "Support will be limited to Gen2 cloudspaces, rendering this field obsolete.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("gen2"),
				},
				Default: stringdefault.StaticString("gen2"),
			},
			"first_ready_timestamp": schema.StringAttribute{
				Computed:            true,
				Description:         "The time when the cloudspace was first ready.",
				MarkdownDescription: "The time when the cloudspace was first ready.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"hacontrol_plane": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "High Availability Kubernetes (replicated control plane for redundancy). This is a critical feature for production workloads.",
				MarkdownDescription: "High Availability Kubernetes (replicated control plane for redundancy). This is a critical feature for production workloads.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Default: booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The id of the cloudspace",
				MarkdownDescription: "The id of the cloudspace",
				DeprecationMessage:  "Use the cloudspace_name attribute instead",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"kubernetes_version": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Kubernetes version to deploy in the cloudspace. Supported values: 1.29.6, 1.30.10, 1.31.1",
				MarkdownDescription: "Kubernetes version to deploy in the cloudspace. Supported values: 1.29.6, 1.30.10, 1.31.1",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Default: stringdefault.StaticString("1.31.1"),
			},
			"last_updated": schema.StringAttribute{
				Computed:            true,
				Description:         "The last time the cloudspace was updated.",
				MarkdownDescription: "The last time the cloudspace was updated.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The name of the cloudspace.",
				MarkdownDescription: "The name of the cloudspace.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), "Must be a valid kubernetes name"),
					stringvalidator.AtLeastOneOf(path.Expressions{path.MatchRoot("name"), path.MatchRoot("cloudspace_name")}...),
				},
			},
			"pending_allocations": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bid_name": schema.StringAttribute{
							Computed: true,
						},
						"count": schema.Int64Attribute{
							Computed: true,
						},
						"server_class": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: PendingAllocationsType{
						ObjectType: types.ObjectType{
							AttrTypes: PendingAllocationsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
			"preemption_webhook": schema.StringAttribute{
				Optional:            true,
				Description:         "Webhook URL for preemption notifications.",
				MarkdownDescription: "Webhook URL for preemption notifications.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
					stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+`), "Must be a valid URL"),
				},
			},
			"region": schema.StringAttribute{
				Required:            true,
				Description:         "The region where the cloudspace will be created.",
				MarkdownDescription: "The region where the cloudspace will be created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"spotnodepool_ids": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				Description:         "IDs of the spotnodepools associated with the cloudspace.",
				MarkdownDescription: "IDs of the spotnodepools associated with the cloudspace.",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
				Create: true,
			}),
			"wait_until_ready": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "If true, waits until the cloudspace control plane is ready",
				MarkdownDescription: "If true, waits until the cloudspace control plane is ready",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Default: booldefault.StaticBool(false),
			},
		},
	}
}

type CloudspaceModel struct {
	Bids                types.Set    `tfsdk:"bids"`
	CloudspaceName      types.String `tfsdk:"cloudspace_name"`
	Cni                 types.String `tfsdk:"cni"`
	DeploymentType      types.String `tfsdk:"deployment_type"`
	FirstReadyTimestamp types.String `tfsdk:"first_ready_timestamp"`
	HacontrolPlane      types.Bool   `tfsdk:"hacontrol_plane"`
	Id                  types.String `tfsdk:"id"`
	KubernetesVersion   types.String `tfsdk:"kubernetes_version"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	Name                types.String `tfsdk:"name"`
	PendingAllocations  types.Set    `tfsdk:"pending_allocations"`
	PreemptionWebhook   types.String `tfsdk:"preemption_webhook"`
	Region              types.String `tfsdk:"region"`
	SpotnodepoolIds     types.List   `tfsdk:"spotnodepool_ids"`
	Timeouts timeouts.Value `tfsdk:"timeouts"`
	WaitUntilReady      types.Bool   `tfsdk:"wait_until_ready"`
}

var _ basetypes.ObjectTypable = BidsType{}

type BidsType struct {
	basetypes.ObjectType
}

func (t BidsType) Equal(o attr.Type) bool {
	other, ok := o.(BidsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t BidsType) String() string {
	return "BidsType"
}

func (t BidsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	bidNameAttribute, ok := attributes["bid_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`bid_name is missing from object`)

		return nil, diags
	}

	bidNameVal, ok := bidNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`bid_name expected to be basetypes.StringValue, was: %T`, bidNameAttribute))
	}

	wonCountAttribute, ok := attributes["won_count"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`won_count is missing from object`)

		return nil, diags
	}

	wonCountVal, ok := wonCountAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`won_count expected to be basetypes.Int64Value, was: %T`, wonCountAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return BidsValue{
		BidName:  bidNameVal,
		WonCount: wonCountVal,
		state:    attr.ValueStateKnown,
	}, diags
}

func NewBidsValueNull() BidsValue {
	return BidsValue{
		state: attr.ValueStateNull,
	}
}

func NewBidsValueUnknown() BidsValue {
	return BidsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewBidsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (BidsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing BidsValue Attribute Value",
				"While creating a BidsValue value, a missing attribute value was detected. "+
					"A BidsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("BidsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid BidsValue Attribute Type",
				"While creating a BidsValue value, an invalid attribute value was detected. "+
					"A BidsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("BidsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("BidsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra BidsValue Attribute Value",
				"While creating a BidsValue value, an extra attribute value was detected. "+
					"A BidsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra BidsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewBidsValueUnknown(), diags
	}

	bidNameAttribute, ok := attributes["bid_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`bid_name is missing from object`)

		return NewBidsValueUnknown(), diags
	}

	bidNameVal, ok := bidNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`bid_name expected to be basetypes.StringValue, was: %T`, bidNameAttribute))
	}

	wonCountAttribute, ok := attributes["won_count"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`won_count is missing from object`)

		return NewBidsValueUnknown(), diags
	}

	wonCountVal, ok := wonCountAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`won_count expected to be basetypes.Int64Value, was: %T`, wonCountAttribute))
	}

	if diags.HasError() {
		return NewBidsValueUnknown(), diags
	}

	return BidsValue{
		BidName:  bidNameVal,
		WonCount: wonCountVal,
		state:    attr.ValueStateKnown,
	}, diags
}

func NewBidsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) BidsValue {
	object, diags := NewBidsValue(attributeTypes, attributes)

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

		panic("NewBidsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t BidsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewBidsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewBidsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewBidsValueNull(), nil
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

	return NewBidsValueMust(BidsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t BidsType) ValueType(ctx context.Context) attr.Value {
	return BidsValue{}
}

var _ basetypes.ObjectValuable = BidsValue{}

type BidsValue struct {
	BidName  basetypes.StringValue `tfsdk:"bid_name"`
	WonCount basetypes.Int64Value  `tfsdk:"won_count"`
	state    attr.ValueState
}

func (v BidsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["bid_name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["won_count"] = basetypes.Int64Type{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.BidName.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["bid_name"] = val

		val, err = v.WonCount.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["won_count"] = val

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

func (v BidsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v BidsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v BidsValue) String() string {
	return "BidsValue"
}

func (v BidsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"bid_name":  basetypes.StringType{},
			"won_count": basetypes.Int64Type{},
		},
		map[string]attr.Value{
			"bid_name":  v.BidName,
			"won_count": v.WonCount,
		})

	return objVal, diags
}

func (v BidsValue) Equal(o attr.Value) bool {
	other, ok := o.(BidsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.BidName.Equal(other.BidName) {
		return false
	}

	if !v.WonCount.Equal(other.WonCount) {
		return false
	}

	return true
}

func (v BidsValue) Type(ctx context.Context) attr.Type {
	return BidsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v BidsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bid_name":  basetypes.StringType{},
		"won_count": basetypes.Int64Type{},
	}
}

var _ basetypes.ObjectTypable = PendingAllocationsType{}

type PendingAllocationsType struct {
	basetypes.ObjectType
}

func (t PendingAllocationsType) Equal(o attr.Type) bool {
	other, ok := o.(PendingAllocationsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t PendingAllocationsType) String() string {
	return "PendingAllocationsType"
}

func (t PendingAllocationsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	bidNameAttribute, ok := attributes["bid_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`bid_name is missing from object`)

		return nil, diags
	}

	bidNameVal, ok := bidNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`bid_name expected to be basetypes.StringValue, was: %T`, bidNameAttribute))
	}

	countAttribute, ok := attributes["count"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`count is missing from object`)

		return nil, diags
	}

	countVal, ok := countAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`count expected to be basetypes.Int64Value, was: %T`, countAttribute))
	}

	serverClassAttribute, ok := attributes["server_class"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`server_class is missing from object`)

		return nil, diags
	}

	serverClassVal, ok := serverClassAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`server_class expected to be basetypes.StringValue, was: %T`, serverClassAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return PendingAllocationsValue{
		BidName:     bidNameVal,
		Count:       countVal,
		ServerClass: serverClassVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewPendingAllocationsValueNull() PendingAllocationsValue {
	return PendingAllocationsValue{
		state: attr.ValueStateNull,
	}
}

func NewPendingAllocationsValueUnknown() PendingAllocationsValue {
	return PendingAllocationsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewPendingAllocationsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (PendingAllocationsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing PendingAllocationsValue Attribute Value",
				"While creating a PendingAllocationsValue value, a missing attribute value was detected. "+
					"A PendingAllocationsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("PendingAllocationsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid PendingAllocationsValue Attribute Type",
				"While creating a PendingAllocationsValue value, an invalid attribute value was detected. "+
					"A PendingAllocationsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("PendingAllocationsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("PendingAllocationsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra PendingAllocationsValue Attribute Value",
				"While creating a PendingAllocationsValue value, an extra attribute value was detected. "+
					"A PendingAllocationsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra PendingAllocationsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewPendingAllocationsValueUnknown(), diags
	}

	bidNameAttribute, ok := attributes["bid_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`bid_name is missing from object`)

		return NewPendingAllocationsValueUnknown(), diags
	}

	bidNameVal, ok := bidNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`bid_name expected to be basetypes.StringValue, was: %T`, bidNameAttribute))
	}

	countAttribute, ok := attributes["count"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`count is missing from object`)

		return NewPendingAllocationsValueUnknown(), diags
	}

	countVal, ok := countAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`count expected to be basetypes.Int64Value, was: %T`, countAttribute))
	}

	serverClassAttribute, ok := attributes["server_class"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`server_class is missing from object`)

		return NewPendingAllocationsValueUnknown(), diags
	}

	serverClassVal, ok := serverClassAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`server_class expected to be basetypes.StringValue, was: %T`, serverClassAttribute))
	}

	if diags.HasError() {
		return NewPendingAllocationsValueUnknown(), diags
	}

	return PendingAllocationsValue{
		BidName:     bidNameVal,
		Count:       countVal,
		ServerClass: serverClassVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewPendingAllocationsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) PendingAllocationsValue {
	object, diags := NewPendingAllocationsValue(attributeTypes, attributes)

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

		panic("NewPendingAllocationsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t PendingAllocationsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewPendingAllocationsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewPendingAllocationsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewPendingAllocationsValueNull(), nil
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

	return NewPendingAllocationsValueMust(PendingAllocationsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t PendingAllocationsType) ValueType(ctx context.Context) attr.Value {
	return PendingAllocationsValue{}
}

var _ basetypes.ObjectValuable = PendingAllocationsValue{}

type PendingAllocationsValue struct {
	BidName     basetypes.StringValue `tfsdk:"bid_name"`
	Count       basetypes.Int64Value  `tfsdk:"count"`
	ServerClass basetypes.StringValue `tfsdk:"server_class"`
	state       attr.ValueState
}

func (v PendingAllocationsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["bid_name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["count"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["server_class"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.BidName.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["bid_name"] = val

		val, err = v.Count.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["count"] = val

		val, err = v.ServerClass.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["server_class"] = val

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

func (v PendingAllocationsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v PendingAllocationsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v PendingAllocationsValue) String() string {
	return "PendingAllocationsValue"
}

func (v PendingAllocationsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"bid_name":     basetypes.StringType{},
			"count":        basetypes.Int64Type{},
			"server_class": basetypes.StringType{},
		},
		map[string]attr.Value{
			"bid_name":     v.BidName,
			"count":        v.Count,
			"server_class": v.ServerClass,
		})

	return objVal, diags
}

func (v PendingAllocationsValue) Equal(o attr.Value) bool {
	other, ok := o.(PendingAllocationsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.BidName.Equal(other.BidName) {
		return false
	}

	if !v.Count.Equal(other.Count) {
		return false
	}

	if !v.ServerClass.Equal(other.ServerClass) {
		return false
	}

	return true
}

func (v PendingAllocationsValue) Type(ctx context.Context) attr.Type {
	return PendingAllocationsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v PendingAllocationsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bid_name":     basetypes.StringType{},
		"count":        basetypes.Int64Type{},
		"server_class": basetypes.StringType{},
	}
}
