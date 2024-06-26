package connections

type ProccesConnection interface {
	Listen() error
	SwitchTransmiFlag()
	GetName() string
	GetSendFlag() bool
}
