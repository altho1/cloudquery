// Code generated by codegen; DO NOT EDIT.

package quicksight

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataSets() *schema.Table {
	return &schema.Table{
		Name:      "aws_quicksight_data_sets",
		Resolver:  fetchQuicksightDataSets,
		Multiplex: client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveTags(),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "column_level_permission_rules_applied",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ColumnLevelPermissionRulesApplied"),
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "data_set_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataSetId"),
			},
			{
				Name:     "import_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImportMode"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "row_level_permission_data_set",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RowLevelPermissionDataSet"),
			},
			{
				Name:     "row_level_permission_tag_configuration_applied",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RowLevelPermissionTagConfigurationApplied"),
			},
		},

		Relations: []*schema.Table{
			Ingestions(),
		},
	}
}
