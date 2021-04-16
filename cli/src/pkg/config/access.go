package config

import (
	"errors"
	"reflect"
)

func (c *Config) getField(name string) *reflect.Value {
	t := reflect.ValueOf(*c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Type().Field(i)
		valueField := t.Field(i)
		tag := field.Tag.Get("config")

		if tag == name {
			return &valueField
		}
	}

	return nil
}

func (c *Config) HasField(name string) bool {
	field := c.getField(name)
	return field != nil
}

func (c *Config) GetField(name string) interface{} {
	field := c.getField(name)
	if field == nil {
		panic(errors.New("field cannot be nil"))
	}

	return field.Interface()
}
