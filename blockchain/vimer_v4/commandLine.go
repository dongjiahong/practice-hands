package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
	createChain --address ADDRESS "create a blockchain"
	addBlock --data DATA "add a block to blockchain"
	send --from FROM --to TO --amount AMOUNT "send coin from FROM to TO"
	getBalance --address ADDRESS "get balance of the address"
	pritChain            "print all blocks"
`

//const AddBlockCmdString = "addBlock"
const PrintChainCmdString = "printChain"
const CreateChainCmdString = "createChain"
const GetBalanceCmdString = "getBalance"
const SendCmdString = "send"

type CLI struct {
	//bc *BlockChain
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

	//addBlockCmd := flag.NewFlagSet(AddBlockCmdString, flag.ExitOnError)
	createChainCmd := flag.NewFlagSet(CreateChainCmdString, flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet(GetBalanceCmdString, flag.ExitOnError)
	sendCmd := flag.NewFlagSet(SendCmdString, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(PrintChainCmdString, flag.ExitOnError)

	//addBlockCmdPara := addBlockCmd.String("data", "", "block transaction info!")
	createChainCmdPara := createChainCmd.String("address", "", "address info!")
	getBalanceCmdPara := getBalanceCmd.String("address", "", "balance info!")

	// send参数
	fromPara := sendCmd.String("from", "", "sender address info!")
	toPara := sendCmd.String("to", "", "to address info!")
	amountPara := sendCmd.Float64("amount", 0, "amount info!")

	switch os.Args[1] {
	case CreateChainCmdString:
		err := createChainCmd.Parse(os.Args[2:])
		CheckErr("Run()1 ", err)
		if createChainCmd.Parsed() {
			if *createChainCmdPara == "" {
				cli.printUsage()
				os.Exit(1)
			}
			cli.CreateChain(*createChainCmdPara)
		}
		/*
			case AddBlockCmdString:
				err := addBlockCmd.Parse(os.Args[2:])
				CheckErr("Run()2 ", err)
				if addBlockCmd.Parsed() {
					if *addBlockCmdPara == "" {
						cli.printUsage()
						os.Exit(1)
					}
					cli.AddBlock(*addBlockCmdPara)
				}
		*/
	case SendCmdString:
		// 发送交易
		err := sendCmd.Parse(os.Args[2:])
		CheckErr("Run2()", err)
		if sendCmd.Parsed() {
			if *fromPara == "" || *toPara == "" || *amountPara == 0 {
				fmt.Println("send cmd parameters invalid!")
				cli.printUsage()
			}
			cli.Send(*fromPara, *toPara, *amountPara)
		}
	case GetBalanceCmdString:
		// 获取余额
		err := getBalanceCmd.Parse(os.Args[2:])
		CheckErr("Run()3 ", err)
		if getBalanceCmd.Parsed() {
			if *getBalanceCmdPara == "" {
				fmt.Println("address should not be empty!")
				cli.printUsage()
			}
			cli.GetBalance(*getBalanceCmdPara)
		}
	case PrintChainCmdString:
		err := printChainCmd.Parse(os.Args[2:])
		CheckErr("Run()4 ", err)
		if printChainCmd.Parsed() {
			cli.PrintChain()
		}
	default:
		cli.printUsage()
	}
}
