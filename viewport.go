/*
  To the extent possible under law, Ian Davis has waived all copyright
  and related or neighboring rights to this Amberfell Source Code file.
  This work is published from the United Kingdom. 
*/
package main

type Viewport struct {
	rotx float64
	roty float64
	rotz float64
	x float64
	y float64
	z float64
	scale float64
	matrix *Matrix4
}

func (self *Viewport) Recalc() {
	self.matrix = NewIdentity().Rotatex(self.rotx).Rotatey(self.roty).Rotatez(self.rotz).Translation(self.transx, self.transy, self.transz)
}


func (self *Viewport) Rotx(angle float64) {
	self.rotx += angle
	if self.rotx > 75 {
		self.rotx = 75
	} else if self.rotx < 15 {
		self.rotx = 15
	}
	self.Recalc()
}
func (self *Viewport) Roty(angle float64) {
	self.roty += angle
	self.Recalc()
}

func (self *Viewport) Transx(d float64) {
	self.transx += d
	self.Recalc()
}

func (self *Viewport) Rotz(angle float64) {
	self.rotz += angle
	self.Recalc()
}


func (self *Viewport) Transy(d float64) {
	self.transy += d
	self.Recalc()
}

func (self *Viewport) Transz(d float64) {
	self.transz += d
	self.Recalc()
}

func (self *Viewport) Zoomstd() {
	self.scale = 0.75
	self.Recalc()
}

func (self *Viewport) Zoomin() {
	self.scale += 0.2
	if self.scale > 3 {
		self.scale = 3
	}
	self.Recalc()
}

func (self *Viewport) Zoomout() {
	self.scale -= 0.2
	if self.scale < 0.2 {
		self.scale = 0.2
	}
	self.Recalc()
}