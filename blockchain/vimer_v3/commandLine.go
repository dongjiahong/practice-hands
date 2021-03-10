package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
	addBlock --data DATA "add a block to blockchain"
	pritChain            "print all blocks"
`

const AddBlockCmdString = "addBlock"
const PrintChainCmdString = "printChain"

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println(usage)
	os.Exit(1)
}

func (cli *CLI) parameterCheck() {
	if len(os.Args) < 2 {
		fmt.Println("invalid input!")
		cli.printUsage()
	}
}

func (cli *CLI) Run() {
	cli.parameterCheck()

	addBlockCmd := flag.NewFlagSet(AddBlockCmdString, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(PrintChainCmdString, flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data", "", "block transaction info!")

	switch os.Args[1] {
	case AddBlockCmdString:
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("Run() ", err)
		if addBlockCmd.Parsed() {
			if *addBlockCmdPara == "" {
				cli.printUsage()
				os.Exit(1)
			}
			cli.AddBlock(*addBlockCmdPara)
		}
	case PrintChainCmdString:
		err := printChainCmd.Parse(os.Args[2:])
		CheckErr("Run() ", err)
		if printChainCmd.Parsed() {
			cli.PrintChain()
		}
	default:
		cli.printUsage()
	}
}
