package dao

type TokenDAO struct {

}

func (t TokenDAO) GenToken(phone string) (token string, err error) {
	// TODO gen token

	// TODO cache token
	return
}

func (t TokenDAO) AuthToken(token string) error {
	// TODO get cache token
	return nil
}

