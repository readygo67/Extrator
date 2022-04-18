package runner

import (
	leveldb "github.com/syndtr/goleveldb/leveldb"
)

const HashLength = 32

// Hash to identify uniqueness
type Hash [HashLength]byte

var (
	KeyLastHandledIndex = []byte("last_handled_index")
	TokenPrefix         = []byte("token")
	SafeMoonPrefix      = []byte("safemoond")
	VictimPrefix        = []byte("victim")
)

func LastHandledIndexStoreKey() []byte {
	return KeyLastHandledIndex
}

func TokenStoreKey(address []byte) []byte {
	return append(TokenPrefix, address...)
}

func SafeMoonStoreKey(account []byte) []byte {
	return append(SafeMoonPrefix, account...)
}

func VictimStoreKey(account []byte) []byte {
	return append(VictimPrefix, account...)
}

func NewDB(path string) (*leveldb.DB, error) {
	return leveldb.OpenFile(path, nil)
}
