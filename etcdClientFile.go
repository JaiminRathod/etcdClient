package main
import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)
func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	add_to_db(context,client,"name","jaimin rathod")
	value,error:=get_from_db(context,client,"name")
	if error==nil {
		fmt.Println(value)
	}
	cancel()
	defer client.Close()
}
func add_to_db(context context.Context,client *clientv3.Client,key string,value string) {
	client.Put(context,key,value)
}
func get_from_db(context context.Context,client *clientv3.Client,key string) (string,error) {
	resp,err:=client.Get(context,key)
	return string(resp.Kvs[0].Value), err
}