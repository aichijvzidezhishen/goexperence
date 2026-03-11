package mongoex

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func Test_readDataFromFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []Order
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", args{"./test_orders.json"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readDataFromFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDataFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDataFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersInsert(t *testing.T) {
	orderList, err := readDataFromFile("./test_orders.json")
	if err != nil {
		t.Errorf("ordersInsert() readData error = %v", err)
	}
	type args struct {
		orders []Order
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", args{orderList}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ordersInsert(tt.args.orders); (err != nil) != tt.wantErr {
				t.Errorf("ordersInsert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindOrders(t *testing.T) {
	type args struct {
		filter bson.M
	}
	tests := []struct {
		name    string
		args    args
		want    []Order
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", args{nil}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindOrders(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOrders() = %v,length %d want %v", got, len(got), tt.want)
			}
		})
	}
}
