package autogen

import (
	"time"
)

/* structure
Article
	Comments
		Content
	Content
*/
type Article struct {
	Author   string    `json:"author"`
	Comments Comments  `json:"comments"`
	Content  Content   `json:"content"`
	Ctime    time.Time `json:"ctime"`
	Title    string    `json:"title"`
}

type Comments []struct {
	Author  string    `json:"author"`
	Content Content   `json:"content"`
	Ctime   time.Time `json:"ctime"`
}

type Content struct {
	Abbrev string `json:"abbrev"`
	Body   string `json:"body"`
}

type Content struct {
	Body string `json:"body"`
}
