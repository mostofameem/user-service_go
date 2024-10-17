package mongodb

import (
	"context"
	"log"
)

type UserTypeRepo struct {
	schema string
}

var userTypeRepo *UserTypeRepo

func initUserTypeRepo() {
	userTypeRepo = &UserTypeRepo{
		schema: "user",
	}
}

func GetUserTypeRepo() *UserTypeRepo {
	return userTypeRepo
}

func (r *UserTypeRepo) Save(data interface{}) error {
	collection := GetDB().Collection(r.schema)

	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	log.Println("Data inserted successfully!")
	return nil
}
