package global

var Initok chan struct{}

func init() {
	//Initok = make(chan struct{})
	InitLog()
}
