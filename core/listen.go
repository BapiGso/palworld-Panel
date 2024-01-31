package core

func (p *PalWorld) Listen(addr string) {
	p.e.Start(addr)
}
