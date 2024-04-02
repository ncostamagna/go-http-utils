package request

import (
	"reflect"
	"strconv"
	"strings"
)

func DecodeMap(m map[string]string, s interface{}) error {
	for k, v := range m {
		if err := setField(s, k, v); err != nil {
			return ErrFormatValue
		}
	}
	return nil
}

func setField(obj interface{}, name string, value string) error {
	fv := getFieldValueByTag(name, "json", obj)
	if !fv.IsValid() || !fv.CanSet() {
		return nil
	}

	ft := fv.Type()
	val := reflect.ValueOf(value)

	if fv.Kind() == reflect.Ptr {
		ft = ft.Elem()
		if ft.Kind() == reflect.Int {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			val = reflect.ValueOf(&i)
			fv.Set(val)
			return nil
		}

		if ft.Kind() == reflect.Bool {
			i, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			val = reflect.ValueOf(&i)
			fv.Set(val)
			return nil
		}
		return nil
	}

	if fv.Kind() == reflect.Int {
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		val = reflect.ValueOf(i)
		fv.Set(val)
		return nil
	}

	if fv.Kind() == reflect.Bool {
		i, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		val = reflect.ValueOf(i)
		fv.Set(val)
		return nil
	}

	if fv.Kind() == reflect.Slice {
		s := strings.Split(value, ",")
		switch ft.Elem().Kind() {
		case reflect.String:
			val = reflect.ValueOf(s)
			fv.Set(val)
		case reflect.Int:
			var items []int
			for _, v := range s {
				i, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				items = append(items, i)
				val = reflect.ValueOf(items)
				fv.Set(val)
			}
		}
	}

	if ft != val.Type() {
		return nil
	}

	fv.Set(val)
	return nil
}

func getFieldValueByTag(tag, key string, s interface{}) reflect.Value {
	val := reflect.ValueOf(s).Elem()
	if val.Kind() != reflect.Struct {
		return reflect.Value{}
	}
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		v := strings.Split(typeField.Tag.Get(key), ",")[0]
		if v == tag {
			return val.Field(i)
		}
	}
	return reflect.Value{}
}
