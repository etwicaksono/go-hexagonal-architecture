package entity

const (
	Error    string = "error"
	Success  string = "success"
	Topic    string = "topic"
	Message  string = "message"
	Request  string = "request"
	Response string = "response"
	Body     string = "body"
)

type BulkWriteResult struct {
	// The number of documents inserted.
	InsertedCount int64

	// The number of documents matched by filters in update and replace operations.
	MatchedCount int64

	// The number of documents modified by update and replace operations.
	ModifiedCount int64

	// The number of documents deleted.
	DeletedCount int64

	// The number of documents upserted by update and replace operations.
	UpsertedCount int64

	// A map of operation index to the _id of each upserted document.
	UpsertedIDs map[int64]interface{}
}
