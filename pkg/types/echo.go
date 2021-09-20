package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Echo struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Echo string             `bson:"echo,omitempty"`
}
