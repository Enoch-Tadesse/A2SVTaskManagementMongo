package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	MISSED    Status = "missed"
	PENDING   Status = "pending"
	COMPLETED Status = "completed"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     Date               `bson:"due_date" json:"due_date"`
	Status      string             `bson:"status" json:"status"`
}
