package models

import (
	"labix.org/v2/mgo/bson"
)

type Comment struct {
	ID        bson.ObjectId  `bson:"_id,omitempty"`
	ParentID  *bson.ObjectId `bson:"parent_id,omitempty"`
	ArticleID bson.ObjectId  `bson:"article_id"`
	UserID    string         `bson:"userid"`
	Content   string         `bson:"content"`
	// Create_At time.Time
}
