package wordpress

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-wordpress",
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"wordpress_author":   tableWordPressAuthor(ctx),
			"wordpress_category": tableWordPressCategory(ctx),
			"wordpress_comment":  tableWordPressComment(ctx),
			"wordpress_core":     tableWordPressCore(ctx),
			"wordpress_post":     tableWordPressPost(ctx),
			"wordpress_tag":      tableWordPressTag(ctx),
		},
	}
	return p
}
