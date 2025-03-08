package resources

import "time"

type PostStruct struct {
	PostId      *int64     `db:"PostID"`
	PostContent []string   `db:"PostContent"`
	CreatedAt   *time.Time `db:"CreatedAt"`
}
