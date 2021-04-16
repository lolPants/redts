package config

import (
	"errors"
	"reflect"
)

func (c *Config) getField(name string) *reflect.Value {
	t := reflect.ValueOf(c).Elem()

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

func (c *Config) GetField(name string) string {
	field := c.getField(name)
	if field == nil {
		panic(errors.New("field cannot be nil"))
	}

	return field.Interface().(string)
}

func (c *Config) SetField(name string, value string) {
	field := c.getField(name)
	if field == nil {
		panic(errors.New("field cannot be nil"))
	}

	field.SetString(value)
}
