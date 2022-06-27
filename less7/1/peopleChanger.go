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
				toType := reflect.TypeOf(mapValue)

				if inFieldType.Name() != toType.Name() {
					return errors.New(fmt.Sprintf("in value %v mismatched types, expected: %v - received %v\n",
						inField.Name,
						inFieldType,
						toType))
				}

				switch c := mapValue.(type) {
				case int:
					if _, ok := mapValue.(int); !ok {
						return errors.New(fmt.Sprintf("type case matching, expected: int - received %v", reflect.TypeOf(c)))
					}
					inFieldValue.SetInt(int64(mapValue.(int)))

				case string:
					if _, ok := mapValue.(string); !ok {
						return errors.New(fmt.Sprintf("type case matching, expected: string - received %v", reflect.TypeOf(c)))
					}
					inFieldValue.SetString(mapValue.(string))

				case float64:
					if _, ok := mapValue.(float64); !ok {
						return errors.New(fmt.Sprintf("type case matching, expected: float64 - received %v", reflect.TypeOf(c)))
					}
					inFieldValue.SetFloat(mapValue.(float64))

				case bool:
					if _, ok := mapValue.(bool); !ok {
						return errors.New(fmt.Sprintf("type case matching, expected: bool - received %v", reflect.TypeOf(c)))
					}
					inFieldValue.SetBool(mapValue.(bool))

				default:
					log.Fatal(fmt.Sprintf("type case matching, received unknown format: %v\n", reflect.TypeOf(c)))
				}
			}
		}
	}
	return nil
}