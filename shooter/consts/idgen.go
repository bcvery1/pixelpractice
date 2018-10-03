package consts

var (
	// IDs is a channel of constantly generated ID numbers
	IDs = make(chan float64, 5)
)

func init() {
	// Constantly create IDs to make removing from map easy and quick
	go func() {
		i := 0.
		for {
			IDs <- i
			i++
		}
	}()
}
