package mongoex

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// 定义与JSON数据匹配的结构体（兼容MongoDB的bson）
type Goods struct {
	GoodsID string  `json:"goodsId" bson:"goodsId"`
	Name    string  `json:"name" bson:"name"`
	Price   float64 `json:"price" bson:"price"`
	Num     int     `json:"num" bson:"num"`
}

type Address struct {
	Province string `json:"province" bson:"province"`
	City     string `json:"city" bson:"city"`
	Detail   string `json:"detail" bson:"detail"`
}

type Order struct {
	UserID       string    `json:"userId" bson:"userId"`
	Username     string    `json:"username" bson:"username"`
	OrderID      string    `json:"orderId" bson:"orderId"`
	OrderTime    time.Time `json:"orderTime" bson:"orderTime"`
	TotalAmount  float64   `json:"totalAmount" bson:"totalAmount"`
	PayStatus    string    `json:"payStatus" bson:"payStatus"`
	PayType      string    `json:"payType" bson:"payType"`
	Goods        []Goods   `json:"goods" bson:"goods"`
	Address      Address   `json:"address" bson:"address"`
	IsRefund     bool      `json:"isRefund" bson:"isRefund"`
	RefundAmount float64   `json:"refundAmount" bson:"refundAmount"`
}

// 方式1：从JSON文件读取数据
func readDataFromFile(filePath string) ([]Order, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败：%v", err)
	}
	defer file.Close()

	var orders []Order
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&orders); err != nil {
		return nil, fmt.Errorf("解析JSON失败：%v", err)
	}
	return orders, nil
}

func ordersInsert(orders []Order) error {
	// 4. 转换为MongoDB可插入的格式
	var insertData []interface{}
	for _, order := range orders {
		insertData = append(insertData, order)
	}

	// // 5. 插入数据（先清空集合，避免重复导入）
	// _, err := collection.DeleteMany(context.TODO(), bson.D{})
	// if err != nil {
	// 	panic(fmt.Sprintf("清空集合失败：%v", err))
	// }
	coll, err := GetCollection("account", "orders")
	if err != nil {
		return err
	}
	coll.InsertMany(context.TODO(), insertData)
	return nil
}

// // FindOrders 根据条件查询订单列表
// filter: 查询条件，例如 bson.M{"status": "paid"}，传 nil 表示查询所有
func FindOrders(filter bson.M) ([]Order, error) {
	// 1. 获取集合
	coll, err := GetCollection("account", "orders")
	if err != nil {
		return nil, err
	}

	// 2. 设置查询选项（可选），例如按创建时间倒序排列
	// opts := options.Find().SetSort(bson.D{{"created_at", -1}})

	// 3. 执行查询
	// 如果 filter 为 nil，MongoDB 会匹配所有文档
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// 4. 解析结果
	var results []Order
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	// 即使没有数据，也返回空切片而不是 nil，保持一致性
	if results == nil {
		results = []Order{}
	}

	return results, nil
}
