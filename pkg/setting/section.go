package setting

type ServerSettingS struct {
	HttpPort string
}

func (s *Setting) ReadSetting(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}