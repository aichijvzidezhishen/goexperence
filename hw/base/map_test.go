package base

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestSyncMap(t *testing.T) {
	SyncMap()
}

func A() {

}

func TestSMap_Rd(t *testing.T) {
	type fields struct {
		rmx sync.RWMutex
		c   map[string]*Entery
	}
	type args struct {
		key string
		tt  time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				key: "k1",
				tt:  1 * time.Second,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				key: "k1",
				tt:  1 * time.Second,
			},
			want: nil,
		}, {
			name: "k2",
			args: args{
				key: "test",
				tt:  1 * time.Second,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SMap{
				rmx: tt.fields.rmx,
				c:   tt.fields.c,
			}
			m.c = make(map[string]*Entery)
			if got := m.Rd(tt.args.key, tt.args.tt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SMap.Rd() = %v, want %v", got, tt.want)
			}
		})
	}
}
