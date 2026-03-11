package redisrankzsetex

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestPrintPlayerPowerRange(t *testing.T) {
	tests := []struct {
		name     string
		players  []redis.Z
		minPower string
		want     string
	}{
		{
			name: "正常情况-单个玩家",
			players: []redis.Z{
				{
					Member: "player1",
					Score:  1000,
				},
			},
			minPower: "500",
			want:     "\n🎯 战力 ≥ 500 的玩家：\n玩家：player1，战力：1000\n",
		},
		{
			name: "正常情况-多个玩家",
			players: []redis.Z{
				{
					Member: "player1",
					Score:  1000,
				},
				{
					Member: "player2",
					Score:  1500,
				},
			},
			minPower: "500",
			want:     "\n🎯 战力 ≥ 500 的玩家：\n玩家：player1，战力：1000\n玩家：player2，战力：1500\n",
		},
		{
			name:     "边界情况-空玩家列表",
			players:  []redis.Z{},
			minPower: "500",
			want:     "\n🎯 战力 ≥ 500 的玩家：\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 捕获标准输出
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintPlayerPowerRange(tt.players, tt.minPower)

			// 恢复标准输出并读取捕获的内容
			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)
			os.Stdout = oldStdout

			got := buf.String()
			if got != tt.want {
				t.Errorf("PrintPlayerPowerRange() = %q, want %q", got, tt.want)
			}
		})
	}
}
