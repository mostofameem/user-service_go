package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r *UserTypeRepo) GetAll() ([]User, error) {
	collection := GetDB().Collection(r.schema)

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []User
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	log.Println("Data retrieved successfully!")
	return users, nil
}

func (r *UserTypeRepo) GetOne(name string) (User, error) {

	collection := GetDB().Collection(r.schema)

	filter := bson.M{"name": name}
	var user User

	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}, fmt.Errorf("no document found with the name %s", name)
		}
		return User{}, err
	}

	log.Println("Data retrieved successfully!")
	return user, nil
}
