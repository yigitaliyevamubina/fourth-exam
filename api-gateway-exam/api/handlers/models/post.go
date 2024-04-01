package models

type Post struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	Title     string `json:"title"`
	Likes     int64  `json:"likes"`
	Dislikes  int64  `json:"dislikes"`
	Views     int64  `json:"views"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListPosts struct {
	Count int64   `json:"count"`
	Posts []*Post `json:"posts"`
}

type PostWithComments struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Content   string     `json:"content"`
	Title     string     `json:"title"`
	Likes     int64      `json:"likes"`
	Dislikes  int64      `json:"dislikes"`
	Views     int64      `json:"views"`
	Category  string     `json:"category"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	Comments  []*Comment `json:"comments"`
}

type PostReq struct {
	PostId string `json:"post_id"`
}
