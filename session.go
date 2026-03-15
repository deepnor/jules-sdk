package jules

// SessionState represents the current state of a session.
type SessionState string

const (
	// SessionStateUnspecified indicates the state is unspecified.
	SessionStateUnspecified SessionState = "STATE_UNSPECIFIED"
	// SessionStateQueued indicates the session is waiting to be processed.
	SessionStateQueued SessionState = "QUEUED"
	// SessionStatePlanning indicates Jules is creating a plan.
	SessionStatePlanning SessionState = "PLANNING"
	// SessionStateAwaitingPlanApproval indicates the plan is ready for user approval.
	SessionStateAwaitingPlanApproval SessionState = "AWAITING_PLAN_APPROVAL"
	// SessionStateAwaitingUserFeedback indicates Jules needs user input.
	SessionStateAwaitingUserFeedback SessionState = "AWAITING_USER_FEEDBACK"
	// SessionStateInProgress indicates Jules is actively working.
	SessionStateInProgress SessionState = "IN_PROGRESS"
	// SessionStatePaused indicates the session is paused.
	SessionStatePaused SessionState = "PAUSED"
	// SessionStateFailed indicates the session failed.
	SessionStateFailed SessionState = "FAILED"
	// SessionStateCompleted indicates the session completed successfully.
	SessionStateCompleted SessionState = "COMPLETED"
)

// AutomationMode controls session automation settings.
type AutomationMode string

const (
	// AutomationModeUnspecified indicates no automation (default).
	AutomationModeUnspecified AutomationMode = "AUTOMATION_MODE_UNSPECIFIED"
	// AutomationModeAutoCreatePR automatically creates a pull request when
	// code changes are ready.
	AutomationModeAutoCreatePR AutomationMode = "AUTO_CREATE_PR"
)

// Session represents a unit of work where Jules executes a coding task.
type Session struct {
	// Name is the resource name of the session (e.g., "sessions/1234567").
	Name string `json:"name,omitempty"`
	// ID is the unique identifier of the session.
	ID string `json:"id,omitempty"`
	// Prompt is the task description provided when creating the session.
	Prompt string `json:"prompt,omitempty"`
	// Title is the optional human-readable title for the session.
	Title string `json:"title,omitempty"`
	// State is the current state of the session.
	State SessionState `json:"state,omitempty"`
	// URL is the web URL to view the session in the Jules UI.
	URL string `json:"url,omitempty"`
	// CreateTime is the RFC 3339 timestamp when the session was created.
	CreateTime string `json:"createTime,omitempty"`
	// UpdateTime is the RFC 3339 timestamp when the session was last updated.
	UpdateTime string `json:"updateTime,omitempty"`
	// SourceContext describes the source repository context for the session.
	SourceContext *SourceContext `json:"sourceContext,omitempty"`
	// RequirePlanApproval indicates whether the session requires explicit
	// plan approval before proceeding.
	RequirePlanApproval *bool `json:"requirePlanApproval,omitempty"`
	// AutomationMode controls the automation behavior for the session.
	AutomationMode AutomationMode `json:"automationMode,omitempty"`
	// Outputs contains the session outputs such as pull requests.
	Outputs []SessionOutput `json:"outputs,omitempty"`
}

// SourceContext provides the context for how to use a source in a session.
type SourceContext struct {
	// Source is the resource name of the source (e.g., "sources/github-myorg-myrepo").
	Source string `json:"source,omitempty"`
	// GitHubRepoContext provides GitHub-specific context.
	GitHubRepoContext *GitHubRepoContext `json:"githubRepoContext,omitempty"`
}

// GitHubRepoContext provides context for using a GitHub repo in a session.
type GitHubRepoContext struct {
	// StartingBranch is the branch to start the session from.
	StartingBranch string `json:"startingBranch,omitempty"`
}

// SessionOutput represents an output of a session.
type SessionOutput struct {
	// PullRequest contains pull request details if one was created.
	PullRequest *PullRequest `json:"pullRequest,omitempty"`
}

// PullRequest represents a pull request created by a session.
type PullRequest struct {
	// URL is the URL of the pull request.
	URL string `json:"url,omitempty"`
	// Title is the title of the pull request.
	Title string `json:"title,omitempty"`
	// Description is the description of the pull request.
	Description string `json:"description,omitempty"`
}

// CreateSessionRequest is the request body for creating a new session.
type CreateSessionRequest struct {
	// Prompt is the task description for the session.
	Prompt string `json:"prompt"`
	// Title is the optional human-readable title for the session.
	Title string `json:"title,omitempty"`
	// SourceContext describes the source repository context.
	SourceContext *SourceContext `json:"sourceContext,omitempty"`
	// RequirePlanApproval controls whether plan approval is required.
	RequirePlanApproval *bool `json:"requirePlanApproval,omitempty"`
	// AutomationMode controls the automation behavior.
	AutomationMode AutomationMode `json:"automationMode,omitempty"`
}

// ListSessionsResponse is the response from listing sessions.
type ListSessionsResponse struct {
	// Sessions is the list of sessions.
	Sessions []Session `json:"sessions,omitempty"`
	// NextPageToken is the token to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// SendMessageRequest is the request body for sending a message to a session.
type SendMessageRequest struct {
	// Prompt is the message text to send.
	Prompt string `json:"prompt"`
}

// SendMessageResponse is the response from sending a message.
// The response is empty on success.
type SendMessageResponse struct{}

// ApprovePlanRequest is the request body for approving a plan.
// The request body is empty.
type ApprovePlanRequest struct{}

// ApprovePlanResponse is the response from approving a plan.
// The response is empty on success.
type ApprovePlanResponse struct{}
