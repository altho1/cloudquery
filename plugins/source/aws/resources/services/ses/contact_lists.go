// Code generated by codegen; DO NOT EDIT.

package ses

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ContactLists() *schema.Table {
	return &schema.Table{
		Name:                "aws_ses_contact_lists",
		Description:         "https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetContactList.html",
		Resolver:            fetchSesContactLists,
		PreResourceResolver: getContactList,
		Multiplex:           client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContactListName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "last_updated_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTimestamp"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "topics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Topics"),
			},
		},
	}
}