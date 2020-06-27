package units

import "fmt"

type C float32
type F float32

func (t C) ToF() F {
	return F(float32(t)*9.0/5.0 + 32)
}

func (t F) ToC() C {
	return C((float32(t) - 32) * 5.0 / 9.0)
}

func (t C) String() string {
	return fmt.Sprintf("%.2f °C", t)
}

func (t F) String() string {
	return fmt.Sprintf("%.2f °F", t)
}

type M float32
type Ft float32

func (d M) ToFt() Ft {
	return Ft(float32(d) * 3.28084)
}

func (d Ft) ToM() M {
	return M(float32(d) / 3.28084)
}

func (d M) String() string {
	return fmt.Sprintf("%.2f m", d)
}

func (d Ft) String() string {
	return fmt.Sprintf("%.2f ft", d)
}

type Pa float32
type Atm float32

func (p Pa) ToAtm() Atm {
	return Atm(float32(p) / 101325)
}

func (p Atm) ToPa() Pa {
	return Pa(float32(p) * 101325)
}

func (p Pa) String() string {
	return fmt.Sprintf("%.2f Pa", p)
}

func (p Atm) String() string {
	return fmt.Sprintf("%.2f atm", p)
}

type Percent float32

func (p Percent) String() string {
	return fmt.Sprintf("%.2f%%", p)
}
