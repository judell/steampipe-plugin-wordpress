package wordpress

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableWordPressCore(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "wordpress_core",
		Description: "Provides basic information about the WordPress site.",
		List: &plugin.ListConfig{
			Hydrate: listWordPressCore,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the WordPress site."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A short description of the site."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The URL of the WordPress site."},
			//{Name: "home_url", Type: proto.ColumnType_STRING, Description: "The home URL of the WordPress site."},
			{Name: "gmt_offset", Type: proto.ColumnType_INT, Description: "The GMT offset of the site."},
			{Name: "timezone_string", Type: proto.ColumnType_STRING, Description: "The timezone string for the site."},
			{Name: "permalink_structure", Type: proto.ColumnType_STRING, Description: "The permalink structure used by the site."},
			{Name: "namespaces", Type: proto.ColumnType_JSON, Description: "The namespaces supported by the WordPress installation."},
			{Name: "authentication", Type: proto.ColumnType_JSON, Description: "Authentication information for the WordPress site."},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "The timezone location of the site.", Transform: transform.FromField("Location").Transform(locationToString)},
			{Name: "raw", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "The raw data of the WordPress site information."},
		},
	}
}

func listWordPressCore(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	rootInfo, resp, err := conn.BasicInfo(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("wordpress_core.listWordPressCore", "api_error", err, "response", resp)
		return nil, err
	}

	d.StreamListItem(ctx, rootInfo)

	return nil, nil
}

func locationToString(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if loc, ok := d.Value.(*time.Location); ok && loc != nil {
		return loc.String(), nil
	}
	return nil, nil
}