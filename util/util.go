package util

import (
	errors "GO_Redis/error"
	"GO_Redis/user"
	"fmt"
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

func MapToUser(m map[string]string) *user.User {
	u := &user.User{
		Name:    m["Name"],
		Surname: m["Surname"],
		Email:   m["Email"],
	}

	age, ok := m["Age"]
	if ok {
		_, err := fmt.Sscanf(age, "%d", &u.Age)
		errors.CheckError(err)
	}

	return u
}
