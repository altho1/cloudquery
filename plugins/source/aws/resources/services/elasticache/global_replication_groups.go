// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GlobalReplicationGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_global_replication_groups",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_GlobalReplicationGroup.html`,
		Resolver:    fetchElasticacheGlobalReplicationGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
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
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "at_rest_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AtRestEncryptionEnabled"),
			},
			{
				Name:     "auth_token_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AuthTokenEnabled"),
			},
			{
				Name:     "cache_node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheNodeType"),
			},
			{
				Name:     "cluster_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ClusterEnabled"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "global_node_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GlobalNodeGroups"),
			},
			{
				Name:     "global_replication_group_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalReplicationGroupDescription"),
			},
			{
				Name:     "global_replication_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalReplicationGroupId"),
			},
			{
				Name:     "members",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Members"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "transit_encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TransitEncryptionEnabled"),
			},
		},
	}
}
