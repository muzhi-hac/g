package main

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	shell "github.com/ipfs/go-ipfs-api"
//	"io/ioutil"
//)
//
//var sh *shell.Shell
//
//// 交易结构体(未来的通道)
//type Transaction struct {
//	Person1      string `json:"person1,omitempty" xml:"person1"`
//	Person2      string `json:"person2,omitempty" xml:"person2"`
//	Person1money string `json:"person1Money,omitempty" xml:"person1Money"`
//	Person2money string `json:"person2Money,omitempty" xml:"person2Money"`
//}
//
//// 数据上传到ipfs
//func UploadIPFS(str string) string {
//	sh = shell.NewShell("localhost:5001") //连接客户端
//	hash, err := sh.Add(bytes.NewBufferString(str))
//	if err != nil {
//		fmt.Println("上传ipfs时错误：", err)
//	}
//	return hash
//}
//
//// 从ipfs获取数据   只读
//func CatIPFS(hash string) string {
//	sh = shell.NewShell("localhost:5001")
//	read, err := sh.Cat(hash) //cat命令用于显示ipfs网络中的一个文件内容，注意显示的是字节形式。
//	if err != nil {
//		fmt.Println(err)
//	}
//	body, err := ioutil.ReadAll(read) //ReadAll 从 r 读取直到出现错误或 EOF 并返回它读取的数据。
//	return string(body)
//}
//func UploadFileIpfs(str string) string {
//	sh = shell.NewShell("localhost:5001")
//	hash, err := sh.AddDir(str)
//	if err != nil {
//		fmt.Println("上传ipfs时错误：", err)
//	}
//	return hash
//}
//
//// 通道序列化
//func marshalStruct(transaction Transaction) []byte {
//	data, err := json.Marshal(&transaction)
//	if err != nil {
//		fmt.Println("序列化err=", err)
//	}
//	return data
//}
//
//// 数据反序列化为通道
//func unmarshalStruct(str []byte) Transaction {
//	var transaction Transaction
//	err := json.Unmarshal(str, &transaction)
//	if err != nil {
//		fmt.Println("unmarshal err=%v", err)
//	}
//	return transaction
//}
//func main() {
//	//生成一个交易结构体(未来的通道)
//	transaction := Transaction{
//		Person1:      "Aaron",
//		Person2:      "Bob",
//		Person1money: "100",
//		Person2money: "200",
//	}
//	//结构体序列化
//	data := marshalStruct(transaction)
//	//上传到ipfs
//	hash := UploadIPFS(string(data))
//	hash2 := UploadFileIpfs("/home/wang/GolandProjects/Gee/Gee/Gee.go")
//	fmt.Println("文件hash是", hash)
//	fmt.Println("文件夹hash是", hash2)
//	//从ipfs下载数据
//	str2 := CatIPFS(hash2)
//	//数据反序列化
//	fmt.Println(str2)
//	//transaction2 := unmarshalStruct([]byte(str2))
//	//验证下数据
//	//fmt.Println(transaction2)
//}
