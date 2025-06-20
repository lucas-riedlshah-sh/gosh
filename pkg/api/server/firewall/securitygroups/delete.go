package securitygroups

import (
	"context"
	"net/url"
	"path"

	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/sitehostnz/gosh/pkg/net"
)

// Delete removes a security group.
func (s *Client) Delete(ctx context.Context, request DeleteRequest) (response models.APIResponse, err error) {
	uri := path.Join(apiPrefix, "delete.json")

	keys := []string{
		"client_id",
		"name",
	}

	values := url.Values{}
	values.Add("client_id", s.client.ClientID)
	values.Add("name", request.Name)

	req, err := s.client.NewRequest("POST", uri, net.Encode(values, keys))
	if err != nil {
		return response, err
	}

	if err := s.client.Do(ctx, req, &response); err != nil {
		return response, err
	}

	return response, nil
}
