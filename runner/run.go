package runner

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/ExtractSafeMoon/config"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"strings"
)

var (
	safeMoonSignatures = []string{
		"4549b039", // reflectionFromToken(uint256 tAmount, bool deductTransferFee) view returns(uint256)
		"2d838119", // tokenFromReflection(uint256 rAmount) view returns(uint256)
	}

	withdrawSignatures = []string{
		"3ccfd60b", // withdraw() returns()
		"2e1a7d4d", // withdraw(uint256 amount) returns()
		"f3fef3a3", // withdraw(address to, uint256 amount) returns()
		"4782f779", // withdrawETH(address to, uint256 amount) returns()
		"e086e5ec", // withdrawETH() returns()
		"f14210a6", // withdrawETH(uint256 amount)
		"205c2878", // withdrawTo(address to, uint256 amount) returns()
		"c3b35a7e", // withdrawTo(address token, address to, uint256 amount) returns()
		"01e33667", // withdrawToken(address token, address to, uint256 amount) returns()
		"223c217b", // withdrawTokenTo(address token, address to, uint256 amount) returns()
	}

	IMPLEMENTATION_SLOT = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
)

func Run(config *config.Config) {
	tokenMap := make(map[common.Address]bool)
	tokenCount := int(0)
	safeMoonList := []common.Address{}
	victimList := []common.Address{}

	db, err := NewDB(config.DB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	storedIndex := int64(0)
	startIndex := int64(0)

	exist, err := db.Has(LastHandledIndexStoreKey(), nil)
	if exist {
		bz, err := db.Get(LastHandledIndexStoreKey(), nil)
		if err != nil {
			panic(err)
		}
		storedIndex = big.NewInt(0).SetBytes(bz).Int64()
		startIndex = storedIndex + 1
	}

	iter := db.NewIterator(util.BytesPrefix(TokenPrefix), nil)
	defer iter.Release()
	for iter.Next() {
		addr := common.BytesToAddress(iter.Value())
		tokenMap[addr] = true
		tokenCount++
	}

	iter = db.NewIterator(util.BytesPrefix(SafeMoonPrefix), nil)
	for iter.Next() {
		addr := common.BytesToAddress(iter.Value())
		safeMoonList = append(safeMoonList, addr)
	}

	iter = db.NewIterator(util.BytesPrefix(VictimPrefix), nil)
	for iter.Next() {
		addr := common.BytesToAddress(iter.Value())
		victimList = append(victimList, addr)
	}

	fmt.Printf("intial, startIndex:%v, tokenCount:%v\n", startIndex, tokenCount)

	rpcURL := config.RPCURL
	c, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	factoryAddr := common.HexToAddress("0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73")

	factory, err := NewPancakeFactory(factoryAddr, c)
	if err != nil {
		panic(err)
	}
	allLength, err := factory.AllPairsLength(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("all length:%v\n", allLength)

	for i := startIndex; i < allLength.Int64(); i++ {
		addrs := []common.Address{}
		pairAddr, _ := factory.AllPairs(nil, big.NewInt(i))
		pair, _ := NewPancakePair(pairAddr, c)

		fmt.Printf("%v, pair:%v\n", i, pairAddr)
		token0Addr, _ := pair.Token0(nil)
		token1Addr, _ := pair.Token1(nil)

		//process proxy token
		bz, err := c.StorageAt(context.Background(), token0Addr, IMPLEMENTATION_SLOT, nil)
		if err == nil {
			addr := common.BytesToAddress(bz)
			codeByte, err := c.CodeAt(context.Background(), addr, nil)
			if err == nil && len(codeByte) != 0 {
				token0Addr = addr
			}
		}

		bz, err = c.StorageAt(context.Background(), token1Addr, IMPLEMENTATION_SLOT, nil)
		if err == nil {
			addr := common.BytesToAddress(bz)
			codeByte, err := c.CodeAt(context.Background(), addr, nil)
			if err == nil && len(codeByte) != 0 {
				token1Addr = addr
			}
		}

		addrs = append(addrs, token0Addr, token1Addr)

		for _, addr := range addrs {
			if tokenMap[addr] == false {
				tokenMap[addr] = true

				db.Put(TokenStoreKey(addr.Bytes()), addr.Bytes(), nil)
				fmt.Printf("tokenCount:%v, addr:%v\n", tokenCount, addr)

				tokenCount++
				byteCode, err := c.CodeAt(context.Background(), addr, nil)
				if err != nil {
					fmt.Printf("fail to getCode:%v\n", addr)
					continue
				}
				code := hex.EncodeToString(byteCode)

				if isSafeMoon(code) {
					db.Put(SafeMoonStoreKey(addr.Bytes()), addr.Bytes(), nil)
					if isVictim(code) {
						db.Put(VictimStoreKey(addr.Bytes()), addr.Bytes(), nil)
						fmt.Printf("found: %v\n", addr)
					}
				}
			}
		}

		db.Put(LastHandledIndexStoreKey(), big.NewInt(i).Bytes(), nil)
	}
}

func isSafeMoon(code string) bool {
	found := false
	for _, signature := range safeMoonSignatures {
		if strings.Contains(code, signature) {
			found = true
			break
		}
	}
	return found
}

func isVictim(code string) bool {
	found := false
	for _, signature := range withdrawSignatures {
		if strings.Contains(code, signature) {
			found = true
			break
		}
	}
	return found
}
