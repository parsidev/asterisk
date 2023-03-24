package events

type AuthDetailEvent struct {
	ObjectType    string
	ObjectName    string
	UserName      string
	Password      string
	Md5Cred       string
	Realm         string
	NonceLifetime string
	AuthType      string
	EndpointName  string
}

func (AuthDetailEvent) EventTypeName() string {
	return "AuthDetail"
}
