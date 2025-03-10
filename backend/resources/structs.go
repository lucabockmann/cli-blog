package resources

import "time"

type PostStruct struct {
	PostId      *int64     `db:"PostID"`
	PostContent string     `db:"PostContent" json:"PostContent"`
	CreatedAt   *time.Time `db:"CreatedAt"`
}

type UserStruct struct {
	UserID     *int64     `db:"UserID"`
	UserName   string     `db:"UserName"`
	ProfilePic string     `db:"ImageUrl"`
	CreatedAt  *time.Time `db:"CreatedAt"`
}
