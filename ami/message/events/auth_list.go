package events

type AuthListEvent struct {
	ObjectType    string
	ObjectName    string
	UserName      string
	Md5Cred       string
	Realm         string
	AuthType      string
	Password      string
	NonceLifetime string
}

func (AuthListEvent) EventTypeName() string {
	return "AuthList"
}
