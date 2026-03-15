package jules

// Activity represents a single event within a session.
type Activity struct {
	// Name is the resource name (e.g., "sessions/123/activities/act1").
	Name string `json:"name,omitempty"`
	// ID is the unique identifier of the activity.
	ID string `json:"id,omitempty"`
	// Originator indicates who created the activity ("system", "agent", "user").
	Originator string `json:"originator,omitempty"`
	// Description is a human-readable description of the activity.
	Description string `json:"description,omitempty"`
	// CreateTime is the RFC 3339 timestamp when the activity was created.
	CreateTime string `json:"createTime,omitempty"`

	// Artifacts contains outputs produced during execution.
	Artifacts []Artifact `json:"artifacts,omitempty"`

	// Event fields — exactly one of these will be populated per activity.

	// PlanGenerated is set when Jules has created a plan for the task.
	PlanGenerated *PlanGenerated `json:"planGenerated,omitempty"`
	// PlanApproved is set when a plan was approved.
	PlanApproved *PlanApproved `json:"planApproved,omitempty"`
	// UserMessaged is set when the user sent a message.
	UserMessaged *UserMessaged `json:"userMessaged,omitempty"`
	// AgentMessaged is set when Jules sent a message.
	AgentMessaged *AgentMessaged `json:"agentMessaged,omitempty"`
	// ProgressUpdated is set when a progress update occurred.
	ProgressUpdated *ProgressUpdated `json:"progressUpdated,omitempty"`
	// SessionCompleted is set when the session finished successfully.
	SessionCompleted *SessionCompleted `json:"sessionCompleted,omitempty"`
	// SessionFailed is set when the session encountered an error.
	SessionFailed *SessionFailed `json:"sessionFailed,omitempty"`
}

// PlanGenerated indicates that Jules has created a plan for the task.
type PlanGenerated struct {
	// Plan is the generated plan.
	Plan *Plan `json:"plan,omitempty"`
}

// PlanApproved indicates that a plan was approved.
type PlanApproved struct {
	// PlanID is the identifier of the approved plan.
	PlanID string `json:"planId,omitempty"`
}

// UserMessaged represents a message from the user.
type UserMessaged struct {
	// UserMessage is the message text.
	UserMessage string `json:"userMessage,omitempty"`
}

// AgentMessaged represents a message from Jules.
type AgentMessaged struct {
	// AgentMessage is the message text.
	AgentMessage string `json:"agentMessage,omitempty"`
}

// ProgressUpdated represents a status update during execution.
type ProgressUpdated struct {
	// Title is a short label for the progress update.
	Title string `json:"title,omitempty"`
	// Description is the detail of the progress update.
	Description string `json:"description,omitempty"`
}

// SessionCompleted indicates the session finished successfully.
type SessionCompleted struct{}

// SessionFailed indicates the session encountered an error.
type SessionFailed struct {
	// Reason describes why the session failed.
	Reason string `json:"reason,omitempty"`
}

// Plan is a sequence of steps that Jules will take to complete a task.
type Plan struct {
	// ID is the unique identifier of the plan.
	ID string `json:"id,omitempty"`
	// Steps is the ordered list of plan steps.
	Steps []PlanStep `json:"steps,omitempty"`
	// CreateTime is the RFC 3339 timestamp when the plan was created.
	CreateTime string `json:"createTime,omitempty"`
}

// PlanStep is a single step in a plan.
type PlanStep struct {
	// ID is the unique identifier of the step.
	ID string `json:"id,omitempty"`
	// Index is the zero-based position of the step in the plan.
	Index int `json:"index,omitempty"`
	// Title is a short label for the step.
	Title string `json:"title,omitempty"`
	// Description is the detail of the step.
	Description string `json:"description,omitempty"`
}

// Artifact represents a single unit of data produced by an activity.
type Artifact struct {
	// ChangeSet is set when the artifact is a set of code changes.
	ChangeSet *ChangeSet `json:"changeSet,omitempty"`
	// BashOutput is set when the artifact is bash command output.
	BashOutput *BashOutput `json:"bashOutput,omitempty"`
	// Media is set when the artifact is a media file.
	Media *Media `json:"media,omitempty"`
}

// ChangeSet represents a set of changes to be applied to a source.
type ChangeSet struct {
	// Source is the resource name of the source.
	Source string `json:"source,omitempty"`
	// GitPatch contains the patch in Git format.
	GitPatch *GitPatch `json:"gitPatch,omitempty"`
}

// GitPatch represents a patch in Git format.
type GitPatch struct {
	// BaseCommitID is the commit the patch is based on.
	BaseCommitID string `json:"baseCommitId,omitempty"`
	// UnidiffPatch is the unified diff patch content.
	UnidiffPatch string `json:"unidiffPatch,omitempty"`
	// SuggestedCommitMessage is the suggested commit message for the patch.
	SuggestedCommitMessage string `json:"suggestedCommitMessage,omitempty"`
}

// BashOutput represents output from a bash command.
type BashOutput struct {
	// Command is the command that was executed.
	Command string `json:"command,omitempty"`
	// Output is the standard output from the command.
	Output string `json:"output,omitempty"`
	// ExitCode is the exit code of the command.
	ExitCode int `json:"exitCode,omitempty"`
}

// Media represents a media file output.
type Media struct {
	// MimeType is the MIME type of the media (e.g., "image/png").
	MimeType string `json:"mimeType,omitempty"`
	// Data is the base64-encoded content of the media.
	Data string `json:"data,omitempty"`
}

// ListActivitiesResponse is the response from listing activities.
type ListActivitiesResponse struct {
	// Activities is the list of activities.
	Activities []Activity `json:"activities,omitempty"`
	// NextPageToken is the token to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}
