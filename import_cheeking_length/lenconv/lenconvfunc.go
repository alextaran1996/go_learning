package lenconv

func MetrtoArsh(m Metr) Arsh {
	return Arsh(m / 0.71)
}

func MetrtoSaz(m Metr) Saz {
	return Saz(m / 2.133)
}

func MetrtoVerst(m Metr) Verst {
	return Verst(m / 1068.8)
}
