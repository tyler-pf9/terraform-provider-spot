// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_region

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RegionDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"country": schema.StringAttribute{
				Computed:            true,
				Description:         "Country of the region",
				MarkdownDescription: "Country of the region",
			},
			"description": schema.StringAttribute{
				Computed:            true,
				Description:         "Description of the region",
				MarkdownDescription: "Description of the region",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the region",
				MarkdownDescription: "Name of the region",
			},
			"provider": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"region_name": schema.StringAttribute{
						Computed:            true,
						Description:         "Name of the region",
						MarkdownDescription: "Name of the region",
					},
					"type": schema.StringAttribute{
						Computed:            true,
						Description:         "Actual infrastructure backing the region",
						MarkdownDescription: "Actual infrastructure backing the region",
					},
				},
				CustomType: ProviderType{
					ObjectType: types.ObjectType{
						AttrTypes: ProviderValue{}.AttributeTypes(ctx),
					},
				},
				Computed: true,
			},
		},
	}
}

type RegionModel struct {
	Country     types.String  `tfsdk:"country"`
	Description types.String  `tfsdk:"description"`
	Name        types.String  `tfsdk:"name"`
	Provider    ProviderValue `tfsdk:"provider"`
}

var _ basetypes.ObjectTypable = ProviderType{}

type ProviderType struct {
	basetypes.ObjectType
}

func (t ProviderType) Equal(o attr.Type) bool {
	other, ok := o.(ProviderType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ProviderType) String() string {
	return "ProviderType"
}

func (t ProviderType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	regionNameAttribute, ok := attributes["region_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region_name is missing from object`)

		return nil, diags
	}

	regionNameVal, ok := regionNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region_name expected to be basetypes.StringValue, was: %T`, regionNameAttribute))
	}

	typeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return nil, diags
	}

	typeVal, ok := typeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be basetypes.StringValue, was: %T`, typeAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ProviderValue{
		RegionName:   regionNameVal,
		ProviderType: typeVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewProviderValueNull() ProviderValue {
	return ProviderValue{
		state: attr.ValueStateNull,
	}
}

func NewProviderValueUnknown() ProviderValue {
	return ProviderValue{
		state: attr.ValueStateUnknown,
	}
}

func NewProviderValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ProviderValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ProviderValue Attribute Value",
				"While creating a ProviderValue value, a missing attribute value was detected. "+
					"A ProviderValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ProviderValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ProviderValue Attribute Type",
				"While creating a ProviderValue value, an invalid attribute value was detected. "+
					"A ProviderValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ProviderValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ProviderValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ProviderValue Attribute Value",
				"While creating a ProviderValue value, an extra attribute value was detected. "+
					"A ProviderValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ProviderValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewProviderValueUnknown(), diags
	}

	regionNameAttribute, ok := attributes["region_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region_name is missing from object`)

		return NewProviderValueUnknown(), diags
	}

	regionNameVal, ok := regionNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region_name expected to be basetypes.StringValue, was: %T`, regionNameAttribute))
	}

	typeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return NewProviderValueUnknown(), diags
	}

	typeVal, ok := typeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be basetypes.StringValue, was: %T`, typeAttribute))
	}

	if diags.HasError() {
		return NewProviderValueUnknown(), diags
	}

	return ProviderValue{
		RegionName:   regionNameVal,
		ProviderType: typeVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewProviderValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ProviderValue {
	object, diags := NewProviderValue(attributeTypes, attributes)

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

		panic("NewProviderValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ProviderType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewProviderValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewProviderValueUnknown(), nil
	}

	if in.IsNull() {
		return NewProviderValueNull(), nil
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

	return NewProviderValueMust(ProviderValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ProviderType) ValueType(ctx context.Context) attr.Value {
	return ProviderValue{}
}

var _ basetypes.ObjectValuable = ProviderValue{}

type ProviderValue struct {
	RegionName   basetypes.StringValue `tfsdk:"region_name"`
	ProviderType basetypes.StringValue `tfsdk:"type"`
	state        attr.ValueState
}

func (v ProviderValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["region_name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["type"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.RegionName.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["region_name"] = val

		val, err = v.ProviderType.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["type"] = val

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

func (v ProviderValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ProviderValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ProviderValue) String() string {
	return "ProviderValue"
}

func (v ProviderValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"region_name": basetypes.StringType{},
			"type":        basetypes.StringType{},
		},
		map[string]attr.Value{
			"region_name": v.RegionName,
			"type":        v.ProviderType,
		})

	return objVal, diags
}

func (v ProviderValue) Equal(o attr.Value) bool {
	other, ok := o.(ProviderValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.RegionName.Equal(other.RegionName) {
		return false
	}

	if !v.ProviderType.Equal(other.ProviderType) {
		return false
	}

	return true
}

func (v ProviderValue) Type(ctx context.Context) attr.Type {
	return ProviderType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ProviderValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"region_name": basetypes.StringType{},
		"type":        basetypes.StringType{},
	}
}
