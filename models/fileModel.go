package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	FileID        primitive.ObjectID `bson:"_id"`
	FileName      string             `bson:"fileName"`
	FileExtension string             `bson:"fileExtension"`
	FileLocation  string             `bson:"fileLocation"`
	FilePassword  string             `bson:"filePassword"`
	UploadedAt    time.Time          `bson:"uploadedAt"`
}
