// Package jules provides access to the Jules session activities API.
//
// Activities represent events that occur during a session, such as plan
// generation, messages, progress updates, and completion events.
package jules

import (
	"context"
	"fmt"
	"net/http"
)

// Get retrieves a single activity by its resource name.
//
// The name should be in the format "sessions/{sessionId}/activities/{activityId}".
func (s *ActivitiesService) Get(ctx context.Context, name string) (*Activity, error) {
	activity := &Activity{}
	path := "/" + name
	if err := s.client.Do(ctx, http.MethodGet, path, nil, activity); err != nil {
		return nil, fmt.Errorf("activities: get: %w", err)
	}
	return activity, nil
}

// List lists all activities for a session.
//
// The sessionName should be in the format "sessions/{sessionId}".
// Use pageSize to control the number of results per page and pageToken to
// retrieve subsequent pages.
func (s *ActivitiesService) List(ctx context.Context, sessionName string, pageSize int, pageToken string) (*ListActivitiesResponse, error) {
	path := "/" + sessionName + "/activities"
	sep := "?"
	if pageSize > 0 {
		path += fmt.Sprintf("%spageSize=%d", sep, pageSize)
		sep = "&"
	}
	if pageToken != "" {
		path += fmt.Sprintf("%spageToken=%s", sep, pageToken)
	}

	resp := &ListActivitiesResponse{}
	if err := s.client.Do(ctx, http.MethodGet, path, nil, resp); err != nil {
		return nil, fmt.Errorf("activities: list: %w", err)
	}
	return resp, nil
}
