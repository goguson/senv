package senv

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Load(cfg any) error {
	t := reflect.TypeOf(cfg)
	fieldsCount := t.NumField()
	cfgValue := reflect.ValueOf(cfg).Elem()
	for i := 0; i < fieldsCount; i++ {
		f := t.Field(i)
		envName := f.Tag.Get("senv")
		if envName == "" {
			fmt.Println("missing")
		}
		env := os.Getenv(envName)
		if env == "" {
			return fmt.Errorf("missing expected env variable: %s", envName)
		}

		switch f.Type.Kind() {
		case reflect.String:
			cfgValue.Field(i).Set(reflect.ValueOf(env))
		case reflect.Int:
			v, err := strconv.Atoi(env)
			if err != nil {
				return fmt.Errorf("reflect.Int: %w", err)
			}
			cfgValue.Field(i).Set(reflect.ValueOf(v))
		case reflect.Bool:
		}
	}
	return nil
}
