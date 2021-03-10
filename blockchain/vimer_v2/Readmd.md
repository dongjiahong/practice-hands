# v2版本思路

1. 定一个工作量证明的结构ProofOfWork
    * block
    * 目标值

2. 提供一个创建POW的方法
    * NewProofOfWork(参数)

3. 提供一个计算哈希值的方法
    * Run()

4. 提供一个校验函数
    * IsValid()
