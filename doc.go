// Package jules provides a client for accessing the Google Jules API (v1alpha).
//
// Construct a new Jules client, then use the various services on the client to
// access different parts of the Jules API. For example:
//
//	client := jules.NewClient(os.Getenv("JULES_API_KEY"))
//
//	// Create a new session
//	session, err := client.Sessions.Create(ctx, &jules.CreateSessionRequest{
//		Prompt: "Format the codebase",
//	})
package jules
