package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

func main() {
	indexerClient := instantiateClient()
	checkAccountBalance(indexerClient)
	getAssetInfo(indexerClient)
	interactWithContract(indexerClient)
}

func instantiateClient() *indexer.Client {
	fmt.Println(">>> instantiateClient")
	const indexerAddress = "https://testnet.algoexplorerapi.io/idx2"
	const indexerToken = ""
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		panic(err)
	}
	return indexerClient
}

func checkAccountBalance(indexerClient *indexer.Client) {
	fmt.Println(">>> checkAccountBalance")
	account := "ECASBGDTZBBXQL4BPAH64U7BR3TI7Y4YHOJVRNGOXVBMKUBLU4DKCZQ7JY"
	_, accountInfo, err := indexerClient.LookupAccountByID(account).Do(context.Background())
	// accountInfo, err = algodClient.AccountInformation(account).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("microAlgos:", accountInfo.Amount)
	fmt.Printf("assets: %+v\n", accountInfo.Assets)
}

func getAssetInfo(indexerClient *indexer.Client) {
	fmt.Println(">>> getAssetInfo")
	assetID := uint64(408947)
	_, assetInfo, err := indexerClient.LookupAssetByID(assetID).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("assetParams: %+v\n", assetInfo.Params)
}

func interactWithContract(indexerClient *indexer.Client) {
	fmt.Println(">>> interactWithContract")
	appID := uint64(43178587)
	appInfo, err := indexerClient.LookupApplicationByID(appID).Do(context.Background())
	if err != nil {
		panic(err)
	}

	globalState := appInfo.Application.Params.GlobalState
	globalStateMap := make(map[string]interface{})
	for i := range globalState {
		key, _ := base64.StdEncoding.DecodeString(globalState[i].Key)
		switch globalState[i].Value.Type {
		case 1:
			globalStateMap[string(key)] = globalState[i].Value.Bytes
		case 2:
			globalStateMap[string(key)] = globalState[i].Value.Uint
		}

	}
	fmt.Printf("globalState: %+v\n", globalStateMap)
}
