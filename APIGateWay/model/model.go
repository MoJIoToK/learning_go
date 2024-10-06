package model

import "time"

type NewsFullDetailed struct {
	News     NewsShortDetailed
	Comments []Comment
}

type NewsShortDetailed struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	PubTime int64  `json:"PubTime"`
	Link    string `json:"Link"`
}

type Comment struct {
	ID       int    `json:"ID"`
	ParentID int    `json:"ParentID"`
	NewsID   int    `json:"NewsID"`
	Content  string `json:"Content"`
}

var HardCode = []NewsShortDetailed{
	{ID: 1, Title: "Title 1", Content: "Content 1", PubTime: time.Now().Unix(), Link: "Link 1"},
	{ID: 2, Title: "Title 2", Content: "Content 2", PubTime: time.Now().Unix(), Link: "Link 2"},
	{ID: 3, Title: "Title 3", Content: "Content 3", PubTime: time.Now().Unix(), Link: "Link 3"},
}

var CommentNews1 = []Comment{
	{ID: 1, ParentID: 0, NewsID: 1, Content: "Content 1"},
}

var CommentNews2 = []Comment{
	{ID: 2, ParentID: 0, NewsID: 1, Content: "Content 2"},
	{ID: 3, ParentID: 0, NewsID: 1, Content: "Content 3"},
}
