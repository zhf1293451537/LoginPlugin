package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	ArticleID primitive.ObjectID  `bson:"article_id,omitempty" json:"article_id,omitempty"`
	UserID    primitive.ObjectID  `bson:"user_id,omitempty" json:"user_id,omitempty"`
	ParentID  *primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
	Content   string              `bson:"content,omitempty" json:"content,omitempty"`
	CreatedAt time.Time           `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
