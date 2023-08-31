package exec

type EntityInterface interface {
	GetIdentity() string
	GetVersion() int64
	AddVersion()
}
