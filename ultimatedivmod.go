package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	modulo := *a % *b
	*a = div
	*b = modulo
}
