package structs

import (
	"reflect"

	"github.com/pkg/errors"
)

// GetValueByAttributeName ...
func GetValueByAttributeName(s interface{}, attribute string) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Struct {
		return nil, errors.New("The model have different type than struct")
	}
	_, ok := reflect.TypeOf(s).FieldByName(attribute)
	if !ok {
		return nil, errors.New("Attribute name doesn't exist in the model")
	}

	r := reflect.ValueOf(s)
	f := reflect.Indirect(r).FieldByName(attribute)
	return f.Interface(), nil
}

// GetFieldNameByAttributeNameAndTag ...
func GetFieldNameByAttributeNameAndTag(s interface{}, attribute, tag string) (string, error) {
	field, ok := reflect.TypeOf(s).FieldByName(attribute)
	if !ok {
		return "", errors.New("Attribute name doesn't exist in the model")
	}
	dbFieldName := field.Tag.Get(tag)
	if len(dbFieldName) < 1 {
		return "", errors.Errorf("Attribute doesn't have '%s' tag", tag)
	}
	return dbFieldName, nil
}
