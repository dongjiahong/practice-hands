package main

func main() {
	bc := NewBlockChain()
	cli := CLI{bc: bc}
	cli.Run()
}
