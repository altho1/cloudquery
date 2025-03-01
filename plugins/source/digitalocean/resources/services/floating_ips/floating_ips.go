// Code generated by codegen; DO NOT EDIT.

package floating_ips

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func FloatingIps() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_floating_ips",
		Resolver: fetchFloatingIpsFloatingIps,
		Columns: []schema.Column{
			{
				Name:     "ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IP"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "droplet",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Droplet"),
			},
		},
	}
}
