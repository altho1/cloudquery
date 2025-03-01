package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbEventSubscriptions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Docdb

	input := &docdb.DescribeEventSubscriptionsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
	}

	p := docdb.NewDescribeEventSubscriptionsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.EventSubscriptionsList
	}
	return nil
}
