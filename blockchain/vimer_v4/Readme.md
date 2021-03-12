# v4思路

1. 创建区块链的操作放到命令
    * NewBlockChain

2. 定义交易结构
    * 交易ID
    * 交易输入: TXInput
    * 交易输出: TXOutput 

3. 根据交易结构，改写代码
    * 创建区块链的时候生成奖励
    * 通过指定地址检索到他相关的UTXO
    * 实现UTXO的转移（创建交易函数: NewTransaction(from, to string, amount float64)）

4. 实现命令
	* send --from FROM --to TO --amount AMOUNT "send coin from FROM to TO"
	* getbalance --address ADDRESS "get balance of the address"
