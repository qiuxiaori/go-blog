package setting

import "time"

type ServerSettingS struct {
	HttpPort string
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSetting(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
