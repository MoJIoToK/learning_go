package model

import "time"

type Comment struct {
	ID       string    `json:"ID" bson:"_id"`
	ParentID string    `json:"ParentID" bson:"ParentID"`
	NewsID   string    `json:"NewsID" bson:"NewsID"`
	Content  string    `json:"Content" bson:"Content"`
	PubTime  time.Time `json:"PubTime" bson:"PubTime"`
	Childs   []Comment `json:"Childs" bson:"Childs"`
}
