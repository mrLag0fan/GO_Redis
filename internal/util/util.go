package util

import (
	"GO_Redis/internal/model"
	"fmt"
	"log"
	"reflect"
)

func StructTOMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		result[fieldName] = field.Interface()
	}
	return result
}

func MapToUser(m map[string]string) *model.User {
	u := &model.User{
		Name:    m["Name"],
		Surname: m["Surname"],
		Email:   m["Email"],
	}

	age, ok := m["Age"]
	if ok {
		_, err := fmt.Sscanf(age, "%d", &u.Age)
		if err != nil {
			log.Fatalf("Can't parse string to int: %s", err)
		}
	}

	return u
}
