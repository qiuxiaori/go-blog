package model

type Auth struct {
	// *Model
	AppKey    string
	AppSecret string
}

func (a Auth) Get() (Auth, error) {
	return Auth{"111", "122"}, nil
}
