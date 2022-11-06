// Code generated by codegen; DO NOT EDIT.

package apprunner

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ObservabilityConfigurations() *schema.Table {
	return &schema.Table{
		Name:                "aws_apprunner_observability_configurations",
		Description:         `https://docs.aws.amazon.com/apprunner/latest/api/API_ObservabilityConfiguration.html`,
		Resolver:            fetchApprunnerObservabilityConfigurations,
		PreResourceResolver: getObservabilityConfiguration,
		Multiplex:           client.ServiceAccountRegionMultiplexer("apprunner"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObservabilityConfigurationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "deleted_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletedAt"),
			},
			{
				Name:     "latest",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Latest"),
			},
			{
				Name:     "observability_configuration_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObservabilityConfigurationName"),
			},
			{
				Name:     "observability_configuration_revision",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ObservabilityConfigurationRevision"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "trace_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TraceConfiguration"),
			},
		},
	}
}