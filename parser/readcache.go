package main

func read_cache_main(args *ProgramArgs) {
	StateLookup = CreateStateLookup()
	c := NewConsumer(*args.name)
	c.Read()
	c.Report()
	//c.Report()
}
