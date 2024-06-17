package connections

type ProccesConnection interface {
	Listen()
	SwitchTransmiFlag()
	GetName() string
}
