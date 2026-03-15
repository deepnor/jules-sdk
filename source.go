package jules

// Source represents a connected repository.
type Source struct {
	// Name is the resource name (e.g., "sources/github-myorg-myrepo").
	Name string `json:"name,omitempty"`
	// ID is the unique identifier of the source.
	ID string `json:"id,omitempty"`
	// GitHubRepo contains GitHub repository details.
	GitHubRepo *GitHubRepo `json:"githubRepo,omitempty"`
}

// GitHubRepo represents a GitHub repository.
type GitHubRepo struct {
	// Owner is the GitHub organization or user that owns the repository.
	Owner string `json:"owner,omitempty"`
	// Repo is the repository name.
	Repo string `json:"repo,omitempty"`
	// IsPrivate indicates whether the repository is private.
	IsPrivate bool `json:"isPrivate,omitempty"`
	// DefaultBranch is the default branch of the repository.
	DefaultBranch *GitHubBranch `json:"defaultBranch,omitempty"`
	// Branches is the list of branches in the repository.
	Branches []GitHubBranch `json:"branches,omitempty"`
}

// GitHubBranch represents a GitHub branch.
type GitHubBranch struct {
	// DisplayName is the name of the branch.
	DisplayName string `json:"displayName,omitempty"`
}

// ListSourcesResponse is the response from listing sources.
type ListSourcesResponse struct {
	// Sources is the list of sources.
	Sources []Source `json:"sources,omitempty"`
	// NextPageToken is the token to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}
