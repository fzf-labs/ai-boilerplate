package service

import "testing"

func TestDeletedUserReference(t *testing.T) {
	t.Parallel()

	const userID = "0190e7f4-1c17-70ee-85e5-d6a963df4821"
	got := deletedUserReference(userID)
	want := deletedUserReferencePrefix + userID
	if got != want {
		t.Fatalf("deletedUserReference() = %q, want %q", got, want)
	}
}
