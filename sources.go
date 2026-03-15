// Package jules provides access to the Jules sources API.
//
// Sources represent repositories connected to Jules. Currently, Jules supports
// GitHub repositories. Use this package to list and retrieve source details.
package jules

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Get retrieves a single source by its resource name.
//
// The name should be in the format "sources/{sourceId}".
func (s *SourcesService) Get(ctx context.Context, name string) (*Source, error) {
	source := &Source{}
	path := "/" + url.PathEscape(name)
	if err := s.client.Do(ctx, http.MethodGet, path, nil, source); err != nil {
		return nil, fmt.Errorf("sources: get: %w", err)
	}
	return source, nil
}

// List lists all sources (repositories) connected to the authenticated user.
//
// Use filter to retrieve specific sources, pageSize to control the number of
// results per page, and pageToken to retrieve subsequent pages.
func (s *SourcesService) List(ctx context.Context, filter string, pageSize int, pageToken string) (*ListSourcesResponse, error) {
	path := "/sources"
	sep := "?"
	if filter != "" {
		path += fmt.Sprintf("%sfilter=%s", sep, url.QueryEscape(filter))
		sep = "&"
	}
	if pageSize > 0 {
		path += fmt.Sprintf("%spageSize=%d", sep, pageSize)
		sep = "&"
	}
	if pageToken != "" {
		path += fmt.Sprintf("%spageToken=%s", sep, pageToken)
		sep = "&"
	}

	resp := &ListSourcesResponse{}
	if err := s.client.Do(ctx, http.MethodGet, path, nil, resp); err != nil {
		return nil, fmt.Errorf("sources: list: %w", err)
	}
	return resp, nil
}
