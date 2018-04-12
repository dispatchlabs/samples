package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/net/context"

	"time"

	"github.com/dispatchlabs/commons/types"
	"github.com/dispatchlabs/commons/utils"
	"github.com/dispatchlabs/dapos"
	"github.com/dispatchlabs/dapos/proto"
	"github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
)

type Meter struct {
	ResultCount int
	Start       time.Time
	End         time.Time
	Total       int
}

// Private key storage (genisis key) for alpha testing
type Config struct {
	PrivateKey string
	From       string
	To         string
	DelegateIP string
	GrpcPort   int
	HTTPPort   string
}

func loadConfig(file string) Config {

	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		utils.Warn(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func newConnection() (*grpc.ClientConn, error) {

	cfg := loadConfig("key.json")

	add := fmt.Sprintf("%s:%d", cfg.DelegateIP, 1973)
	con, err := grpc.Dial(add, grpc.WithInsecure())

	if err != nil {
		utils.Error(err.Error())
		return nil, err
	}

	return con, nil
}

func buildGRPCConnectionPool(connections int) *grpcpool.Pool {

	var f grpcpool.Factory
	f = newConnection

	p, err := grpcpool.New(f, connections, connections, -1)

	if err != nil {
		utils.Error(err.Error())
	}

	return p
}

func createSampleTransaction(cfg *Config) *types.Transaction {

	// TODO:  Make a wallet bucket ... many wallets and add coins
	tx := types.NewTransaction(cfg.PrivateKey, 1,
		cfg.From,
		cfg.To, 1, time.Now())

	return tx
}

func sendGprcTransaction(tx *types.Transaction, cfg *Config, mtr *Meter, pool *grpcpool.Pool) *types.Action {

	byt := []byte(tx.String())
	buffer := new(bytes.Buffer)
	buffer.Write(byt)

	client, err := pool.Get(context.Background())
	defer client.Close()

	if err != nil {
		utils.Warn(err.Error())
	} else {
		client.ClientConn.GetState()

		if client.ClientConn == nil {
			add := fmt.Sprintf("%s:%d", cfg.DelegateIP, 1973)
			con, err := grpc.Dial(add, grpc.WithInsecure())

			if err != nil {

				client.ClientConn = con
			} else {
				return nil
			}
		}
	}

	actionType := types.ActionNewTransaction
	payLoad := tx.String()

	//defer client.ClientConn.Close()

	p := proto.NewDAPoSGrpcClient(client.ClientConn)
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()
	response, err := p.Execute(contextWithTimeout, &proto.Request{Action: actionType, Payload: payLoad})
	if err != nil {
		utils.Warn(err)
	} else {
		action, err := types.ToActionFromJson([]byte(response.Payload))
		if err != nil {
			return types.NewActionWithStatus(actionType, types.StatusInternalError, err.Error())
		}
		return action
	}

	return nil
}

func sendHttpTransaction(tx *types.Transaction, cfg *Config, mtr *Meter) (http.Response, string) {

	var bodyString string

	byt := []byte(tx.String())
	buffer := new(bytes.Buffer)
	buffer.Write(byt)

	url := "http://" + cfg.DelegateIP + ":" + cfg.HTTPPort + "/v1/transactions"
	resp, err := http.Post(url, "application/json", buffer)
	if err != nil {
		fmt.Println(err.Error())
	} else {
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
	}

	return *resp, bodyString
}

func run(cfg *Config, mtr *Meter, pool *grpcpool.Pool) {

	ret := createSampleTransaction(cfg)
	resp := sendGprcTransaction(ret, cfg, mtr, pool)

	if resp.Status == types.StatusPending {
		if mtr.ResultCount < mtr.Total {
			mtr.ResultCount++

			if mtr.ResultCount%100 == 0 {
				fmt.Println(mtr.ResultCount)
			}
		} else {
			if mtr.End == mtr.Start {
				mtr.End = time.Now()

				fmt.Println("Calculating")
				fmt.Println("Total TX %d, Time Diff %d", mtr.Total, time.Since(mtr.Start))
				fmt.Println("DONE")
			}
		}
	}
}

func getHttpReceipt(id string) string {

	var waiting bool
	var action *types.Action
	waiting = true

	service := dapos.GetDAPoSService()

	for waiting {
		action = service.GetAction(id)

		if action.Status == types.StatusPending {
			time.Sleep(500 * time.Millisecond)
		} else {
			return action.Status
		}
	}
	return types.StatusInternalError
}

func runHttp(cfg *Config, mtr *Meter) {
	ret := createSampleTransaction(cfg)

	//fmt.Println(ret)

	resp, body := sendHttpTransaction(ret, cfg, mtr)
	fmt.Println(body)

	if resp.StatusCode == 200 {

		action, err := types.ToActionFromJson([]byte(body))
		if err != nil {
			utils.Warn(err)
		}
		status := getHttpReceipt(action.Id)

		if status != types.StatusOk {
			utils.Warn("Transaction Failed or rejected")
		} else {

		}
		if mtr.ResultCount < mtr.Total {
			mtr.ResultCount++
		} else {
			if mtr.End == mtr.Start {
				mtr.End = time.Now()

				fmt.Println("Calculating")
				fmt.Println("Total TX %d, Time Diff %d", mtr.Total, time.Since(mtr.Start))
				fmt.Println("DONE")
			}
		}
	}
}

func runLoad(cfg *Config, tx int, mtr *Meter, pool *grpcpool.Pool, runType string) {

	for i := 0; i <= tx; i++ {
		if runType == "GRPC" {
			go run(cfg, mtr, pool)
		}
		if runType == "HTTP" {
			runHttp(cfg, mtr)
		}
	}

}

func wait() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("\nReceived an interrupt, stopping services...\n")
			//cleanup(services, c)
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func main() {

	var mtr Meter
	mtr.ResultCount = 0
	mtr.Start = time.Now()
	mtr.End = mtr.Start
	mtr.Total = 1000

	pool := buildGRPCConnectionPool(10)

	fmt.Println("Strating load test")
	cfg := loadConfig("./key.json")

	runLoad(&cfg, mtr.Total, &mtr, pool, "HTTP")

	mtr.ResultCount = 0
	mtr.Start = time.Now()
	mtr.End = mtr.Start
	mtr.Total = 1000

	runLoad(&cfg, mtr.Total, &mtr, pool, "GRPC")

	wait()
}
