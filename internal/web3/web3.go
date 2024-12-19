package web3

import (
	"context"
	"fmt"
	"go-api/internal/reader"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Connect() {
    client, err := ethclient.Dial("http://localhost:8545")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")

	//address
	address := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	fmt.Println("we have an address", address)


	//balance
	balance, err := client.BalanceAt(context.Background(), address, nil)

	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	float, succes := fbalance.SetString(balance.String())
	fmt.Println(float, succes)
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(balance, ethValue)


	//blocks
	// blockNumber := big.NewInt(0)
	blockNumber, err := client.BlockNumber(context.Background())

	fmt.Println(blockNumber, big.NewInt(int64(1)))

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))

	if err != nil {
		log.Fatal(err)
	  }
	  fmt.Printf("\n ---------- block stuff ----------- \n \n")
	  fmt.Println(block.Number().Uint64())     
	  fmt.Println(block.Time())       
	  fmt.Println(block.Difficulty().Uint64()) 
	  fmt.Println(block.Hash().Hex())          
	  fmt.Println(block.Transactions())   

	  //txns

	  fmt.Printf("\n ----------- txns!! ----------- \n \n")
	  for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        
		fmt.Println(tx.Value().String())    
		// fmt.Println(tx.Gas())               
		// fmt.Println(tx.GasPrice().Uint64()) 
		// fmt.Println(tx.Nonce())             
		// fmt.Println(tx.Data())              
		fmt.Println("to:", tx.To().Hex())      
		chainId, err := client.NetworkID(context.Background())
		if err != nil {
		  log.Fatal(err)
		}

		sender, err := types.Sender(types.NewLondonSigner(chainId), tx)
		if err != nil {
			log.Fatal("Not able to retrieve sender:", err)
		}
	
		fmt.Println(sender.Hex())
		
		// if from, err := types.Sender(types.NewEIP155Signer(chainID), tx); err != nil {
		//   fmt.Println("from: ",chainID,  from.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		// }    
	  }


	  // contracts
	//   contractAddress := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
	//   instance, err := store
	contractReader, err := reader.New(reader.Config{
		Client: client,
		ContractAddress: "0x5fbdb2315678afecb367f032d93f642f64180aa3",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer contractReader.Close()
	


	name, err := contractReader.ReadTokenName(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name of token is: ", name)



}