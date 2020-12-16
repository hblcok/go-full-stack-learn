package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apache/dubbo-go/common/logger"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/config_center/nacos"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/rest"
	_ "github.com/apache/dubbo-go/registry/nacos"
	_ "github.com/apache/dubbo-go/registry/protocol"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	"hb.study/clound-native/dubbo-go/code/basic/hello_world/client/consumer"
)

var (
	survivalTimeout int = 10e9
)

/**
先设置环境变量：
	export CONF_CONSUMER_FILE_PATH="./conf/client.yml"
	export APP_LOG_CONF_FILE="./conf/log.yml"
*/
func main() {
	userProvider := new(consumer.UserProvider)
	config.SetConsumerService(userProvider)
	logger.Infof("userProvider :%v", userProvider)
	test(userProvider)

	config.Load()
	initSignal()
}
func test(userProvider *consumer.UserProvider) {
	req := &consumer.UserRequest{
		ID: 1,
	}
	user, err := userProvider.GetUser(req)
	if err != nil {
		logger.Error("userProvider.GetUser(req) failed,", err)
		return
	}
	logger.Infof("success get user :%v", user)
}

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP,
		syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
		// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("app exit now...")
			return
		}
	}
}
