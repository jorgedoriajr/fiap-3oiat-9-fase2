package valueobjects

type Status string

const (
	Created  Status = "created"
	Approved Status = "approved"
	Rejected Status = "rejected"
)
