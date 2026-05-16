package mongoex

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// 全局变量
var (
	client     *mongo.Client
	clientOnce sync.Once
	connectErr error

	// 默认配置（可通过环境变量覆盖）
	defaultURI         = "mongodb://localhost:27017"
	defaultTimeout     = 15 * time.Second
	defaultMaxPoolSize = uint64(150)
	defaultMinPoolSize = uint64(10)
)

// init 注册程序退出钩子，自动关闭连接
func init() {
	// 监听程序退出信号（Ctrl+C、kill等）
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigChan
		// 收到退出信号时关闭连接
		if client != nil {
			if err := client.Disconnect(context.TODO()); err != nil {
				fmt.Printf("[WARN] 关闭MongoDB连接失败: %v\n", err)
			} else {
				fmt.Println("[INFO] MongoDB连接已正常关闭")
			}
		}
		os.Exit(0)
	}()
}

// 内部初始化函数，首次调用GetCollection时触发
func initClient() {
	// 从环境变量读取配置（优先级高于默认值）
	uri := getEnv("MONGO_URI", defaultURI)
	timeoutStr := getEnv("MONGO_TIMEOUT", "15")
	maxPoolSizeStr := getEnv("MONGO_MAX_POOL_SIZE", "150")
	minPoolSizeStr := getEnv("MONGO_MIN_POOL_SIZE", "10")

	// 解析配置
	timeout, err := time.ParseDuration(timeoutStr + "s")
	if err != nil {
		connectErr = fmt.Errorf("解析超时配置失败: %v，使用默认值%v", err, defaultTimeout)
		timeout = defaultTimeout
	}

	maxPoolSize, err := strconv.ParseUint(maxPoolSizeStr, 10, 64)
	//大

	if err != nil {
		connectErr = fmt.Errorf("解析最大连接池配置失败: %v，使用默认值%v", err, defaultMaxPoolSize)
		maxPoolSize = defaultMaxPoolSize
	}

	minPoolSize, err := strconv.ParseUint(minPoolSizeStr, 10, 64)
	if err != nil {
		connectErr = fmt.Errorf("解析最小连接池配置失败: %v，使用默认值%v", err, defaultMinPoolSize)
		minPoolSize = defaultMinPoolSize
	}

	// 构建客户端选项
	clientOpts := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(maxPoolSize).
		SetMinPoolSize(minPoolSize).
		SetConnectTimeout(timeout)

	// 创建Client
	client, connectErr = mongo.Connect(context.TODO(), clientOpts)
	if connectErr != nil {
		connectErr = fmt.Errorf("创建MongoDB Client失败: %v", connectErr)
		return
	}

	// 验证连接
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		connectErr = fmt.Errorf("MongoDB连接验证失败: %v", err)
		_ = client.Disconnect(context.TODO())
		client = nil
	}
}

// 辅助函数：读取环境变量，无值则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetCollection 对外暴露的核心函数：获取集合实例
// 首次调用时自动初始化连接，后续调用直接复用
func GetCollection(dbName, collName string) (*mongo.Collection, error) {
	// 首次调用时初始化Client
	clientOnce.Do(initClient)

	// 检查初始化是否失败
	if connectErr != nil {
		return nil, fmt.Errorf("MongoDB初始化失败: %v", connectErr)
	}
	if client == nil {
		return nil, fmt.Errorf("MongoDB Client未初始化成功")
	}

	// 返回集合实例
	return client.Database(dbName).Collection(collName), nil
}

// 可选：暴露GetClient，方便特殊场景使用
func GetClient() (*mongo.Client, error) {
	clientOnce.Do(initClient)
	if connectErr != nil {
		return nil, connectErr
	}
	return client, nil
}
