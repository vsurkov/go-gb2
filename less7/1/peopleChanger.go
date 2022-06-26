package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func peopleChanger(in *Person, values map[string]interface{}) error {
	if in == nil {
		return errors.New("impute is nil")
	}

	if values == nil {
		return nil
	}

	inValue := reflect.ValueOf(in)
	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
	}

	for i := 0; i < inValue.NumField(); i++ {
		inField := inValue.Type().Field(i)

		if inField.Type.Kind() != reflect.Struct {

			//nested fields not support yet
			mapValue := values[inField.Name]
			if mapValue != nil {
				inFieldValue := inValue.Field(i)
				inFieldType := inField.Type

				err := updateField(inFieldType, &inFieldValue, mapValue)
				if err != nil {
					return errors.New(fmt.Sprintf("in value: %v - %v", inField.Name, err))
				}
			}
		}
	}

	return nil
}

func updateField(fromType reflect.Type, fromValue *reflect.Value, to interface{}) error {
	toType := reflect.TypeOf(to)

	if fromType != toType {
		return errors.New(fmt.Sprintf("mismatched types, expected: %v - received %v\n", fromType, toType))
	}

	switch to.(type) {
	case int:
		log.Println(fromValue.Int())
		//from = to.(int64)
	case float64:
		log.Println(fromValue.Float())
		//from = to.(float64)
	case string:
		log.Println(fromValue.String())
		//from = to.(string)
	case bool:
		log.Println(fromValue.Bool())
		//toVal := to.(bool)
	default:
		return errors.New("unknown type")
	}
	return nil
}
