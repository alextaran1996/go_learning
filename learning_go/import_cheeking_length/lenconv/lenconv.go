package lenconv

import "fmt"

type Metr float64
type Arsh float64
type Saz float64
type Verst float64

func (m Metr) String() string {
	return fmt.Sprintf("%g m", m)
}

func (a Arsh) String() string {
	return fmt.Sprintf("%g arsh", a)
}

func (s Saz) String() string {
	return fmt.Sprintf("%g saz", s)
}

func (v Verst) String() string {
	return fmt.Sprintf("%g verst", v)
}
