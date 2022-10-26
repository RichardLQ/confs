package confs

type Starter struct {
	confs *Container
	log string
}

//启动服务
func NewStart() *Starter {
	return &Starter{
		confs: NewContainer().Start(),
		log: "",
	}
}

func (s *Starter) BinComb(u ...interface{})  {
	Default.StartConf(u ...)
}