// Package jules provides access to the Jules sessions API.
//
// A session represents a unit of work where Jules executes a coding task
// on a repository. Use this package to create, list, get, and interact with
// sessions.
package jules

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Create creates a new session to start a coding task.
func (s *SessionsService) Create(ctx context.Context, req *CreateSessionRequest) (*Session, error) {
	session := &Session{}
	if err := s.client.Do(ctx, http.MethodPost, "/sessions", req, session); err != nil {
		return nil, fmt.Errorf("sessions: create: %w", err)
	}
	return session, nil
}

// Get retrieves a single session by its resource name.
//
// The name should be in the format "sessions/{sessionId}".
func (s *SessionsService) Get(ctx context.Context, name string) (*Session, error) {
	session := &Session{}
	path := "/" + url.PathEscape(name)
	if err := s.client.Do(ctx, http.MethodGet, path, nil, session); err != nil {
		return nil, fmt.Errorf("sessions: get: %w", err)
	}
	return session, nil
}

// List lists all sessions for the authenticated user.
//
// Use pageSize to control the number of results per page and pageToken to
// retrieve subsequent pages.
func (s *SessionsService) List(ctx context.Context, pageSize int, pageToken string) (*ListSessionsResponse, error) {
	path := "/sessions"
	sep := "?"
	if pageSize > 0 {
		path += fmt.Sprintf("%spageSize=%d", sep, pageSize)
		sep = "&"
	}
	if pageToken != "" {
		path += fmt.Sprintf("%spageToken=%s", sep, pageToken)
	}

	resp := &ListSessionsResponse{}
	if err := s.client.Do(ctx, http.MethodGet, path, nil, resp); err != nil {
		return nil, fmt.Errorf("sessions: list: %w", err)
	}
	return resp, nil
}

// ApprovePlan approves a pending plan in a session.
//
// The sessionName should be in the format "sessions/{sessionId}".
// This is only needed when requirePlanApproval was set to true when
// creating the session.
func (s *SessionsService) ApprovePlan(ctx context.Context, sessionName string) error {
	path := "/" + url.PathEscape(sessionName) + ":approvePlan"
	if err := s.client.Do(ctx, http.MethodPost, path, &ApprovePlanRequest{}, nil); err != nil {
		return fmt.Errorf("sessions: approve plan: %w", err)
	}
	return nil
}

// SendMessage sends a message from the user to an active session.
//
// The sessionName should be in the format "sessions/{sessionId}".
func (s *SessionsService) SendMessage(ctx context.Context, sessionName string, prompt string) error {
	path := "/" + url.PathEscape(sessionName) + ":sendMessage"
	req := &SendMessageRequest{Prompt: prompt}
	if err := s.client.Do(ctx, http.MethodPost, path, req, nil); err != nil {
		return fmt.Errorf("sessions: send message: %w", err)
	}
	return nil
}
