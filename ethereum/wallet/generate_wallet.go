package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// GenerateKey用来生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// FromECDSA方法将privateKey转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// hexutil.Encode方法将字节转换为十六进制字符串, 然后我们去除十六进制前的0x
	fmt.Println("====> privateKey: ", hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	// Public可以用来生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("can't assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 转换为十六进制，并去除前面的0x和04前缀, 04前缀是EC的前缀，始终有，但不是必须的
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("=====> publicKey: ", hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	// PubkeyToAddress公钥生成公告地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("=====> public address: ", address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	// 公共地址其实就是Keccak-256哈希, 然后我们取40个字节(20个字节)并用0x作前缀
	// 一下是使用golang.org/x/crypto/sha3的Keccak256函数手动完成的方法, 结果跟上面的PubKeyToAddress一样
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
