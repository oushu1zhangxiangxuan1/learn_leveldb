package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

// type Transaction struct{
// 	  Contract int
// 	  ContractType type = 1;
// 	  google.protobuf.Any parameter = 2;
// 	  bytes provider = 3;
// 	  bytes ContractName = 4;
// 	  int32 Permission_id = 5;
// 	}

func main(){
	if len(os.Args)<2{
		fmt.Println("请填写数据目录")
		return
	}
    fmt.Println("hello")

    // 创建或打开一个数据库
    db,err := leveldb.OpenFile(os.Args[1], nil)
    if err != nil {
        panic(err)
    }

    defer db.Close()
    // 遍历数据库内容
    iter := db.NewIterator(nil, nil)
	i:=0
    for iter.Next() {
		if(i>5){
			break
		}
		i++
        fmt.Printf("raw:\n[%s] ===> %s\n", string(iter.Key()), string(iter.Value()))
        fmt.Printf("hex:\n[%s]n", hex.EncodeToString(iter.Key()))
    }
    iter.Release()
    err = iter.Error()
    if err != nil {
        panic(err)
    }

    fmt.Println("***************************************************")

    // // 读取某条数据
    // data, err := db.Get([]byte("2"), nil)
    // fmt.Printf("[2]:%s:%s\n",data,err)

    // // 根据前缀遍历数据库内容
    // fmt.Println("***************************************************")
    // iter = db.NewIterator(util.BytesPrefix([]byte("foo-")), nil)
    // for iter.Next() {
    //     fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
    // }
    // iter.Release()
    // err = iter.Error()

    // // 遍历从指定 key
    // fmt.Println("***************************************************")
    // iter = db.NewIterator(nil, nil)
    // for ok := iter.Seek([]byte("5")); ok; ok = iter.Next() {
    //     fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
    // }
    // iter.Release()
    // err = iter.Error()

    // // 遍历子集范围
    // fmt.Println("***************************************************")
    // iter = db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("loo")}, nil)
    // for iter.Next() {
    //     fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
    // }
    // iter.Release()
    // err = iter.Error()
 
    // // 批量操作
    // fmt.Println("***************************************************")
    // batch := new(leveldb.Batch)
    // batch.Put([]byte("foo"), []byte("value"))
    // batch.Put([]byte("bar"), []byte("another value"))
    // batch.Delete([]byte("baz"))
    // err = db.Write(batch, nil)
 
    // // 遍历数据库内容
    // iter = db.NewIterator(nil, nil)
    // for iter.Next() {
    //     fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
    // }
    // iter.Release()
    // err = iter.Error()
 
}