package display

// NOTE: we can't use !+..!- comments to excerpt these tests
// into the book because it defeats the Example mechanism,
// which requires the // Output comment to be at the end
// of the function.

func Example_cycle() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}

	Display("c", c)
	// Output:
	// Display c (display.Cycle):
	// c.Value = 42
	// (*c.Tail).Value = 42
	// (*(*c.Tail).Tail).Value = 42
	// ...ad infinitum...

}
