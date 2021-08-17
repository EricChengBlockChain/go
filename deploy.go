package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
  
    "inbox"//这个文件夹是在goroot下面新建一个box,在这个box文件夹下面有一个inbox.sol合约生成的inbox.go文件
)

func main() {
    client, err := ethclient.Dial("https://kovan.infura.io/v3/9aa3d95b3...你的节点")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA("d48c01744fcae61c9ffeb272895...你的私钥")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    input := "1.0"
    address, tx, instance, err := inbox.DeployInbox(auth, client, input)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("\n合约地址:\t")
    fmt.Println(address.Hex())   // 0x147B8eb97fD247D06....
    fmt.Print("交易哈希:\t")
    fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc9....
    fmt.Println("\n\n")
    _ = instance
}

