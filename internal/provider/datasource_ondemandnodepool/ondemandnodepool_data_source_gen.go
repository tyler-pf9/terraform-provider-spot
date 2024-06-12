// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_ondemandnodepool

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OndemandnodepoolDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cloudspace_name": schema.StringAttribute{
				Computed:            true,
				Description:         "The name of the cloudspace",
				MarkdownDescription: "The name of the cloudspace",
			},
			"desired_server_count": schema.Int64Attribute{
				Computed:            true,
				Description:         "The desired number of servers in the node pool.",
				MarkdownDescription: "The desired number of servers in the node pool.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Name of the ondemandnodepool",
				MarkdownDescription: "Name of the ondemandnodepool",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), "Must be a valid kubernetes name"),
				},
			},
			"reserved_count": schema.Int64Attribute{
				Computed:            true,
				Description:         "Number of reserved on-demand nodes.",
				MarkdownDescription: "Number of reserved on-demand nodes.",
			},
			"reserved_status": schema.StringAttribute{
				Computed:            true,
				Description:         "Status of the ondemandnodepool.",
				MarkdownDescription: "Status of the ondemandnodepool.",
			},
			"server_class": schema.StringAttribute{
				Computed:            true,
				Description:         "The class of servers used for the node pool",
				MarkdownDescription: "The class of servers used for the node pool",
			},
		},
	}
}

type OndemandnodepoolModel struct {
	CloudspaceName     types.String `tfsdk:"cloudspace_name"`
	DesiredServerCount types.Int64  `tfsdk:"desired_server_count"`
	Name               types.String `tfsdk:"name"`
	ReservedCount      types.Int64  `tfsdk:"reserved_count"`
	ReservedStatus     types.String `tfsdk:"reserved_status"`
	ServerClass        types.String `tfsdk:"server_class"`
}
