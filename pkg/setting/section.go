package setting

import (
	"reflect"
	"strings"
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Settings struct {
	Server ServerSettingS
}

func (s *Setting) ReadSection(key string, value any) error {
	err := s.vp.UnmarshalKey(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *Setting) ReadAllSection(value any) error {
	err := s.vp.Unmarshal(value)
	if err != nil {
		return err
	}
	return nil
}

func (s *Setting) BindEnvs(iface any, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapsstructure")
		if !ok {
			tv = strings.ToUpper(t.Name)
		}
		switch v.Kind() {
		case reflect.Struct:
			s.BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			s.vp.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
