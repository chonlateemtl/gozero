package atm

type Payer interface {
	Deposit(int)
	WithDraw() int
}

type atm struct {
	payer Payer
}

func New() {

}
