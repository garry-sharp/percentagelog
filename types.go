package percentagelog

type Printable interface {
	String() string
	Percentage() float32
}
