package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/net/context"

	"time"

	"github.com/dispatchlabs/commons/types"
	"github.com/dispatchlabs/commons/utils"
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
}

func loadConfig(file string) Config {

	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func newConnection() (*grpc.ClientConn, error) {

	cfg := loadConfig("key.json")

	add := fmt.Sprintf("%s:%d", cfg.DelegateIP, 1975)
	con, err := grpc.Dial(add, grpc.WithInsecure())

	if err != nil {

		return nil, err
	}

	return con, nil
}

func buildGRPCConnectionPool(connections int) *grpcpool.Pool {

	var f grpcpool.Factory

	f = newConnection

	p, err := grpcpool.New(f, connections, connections, 0)

	if err != nil {
		fmt.Println("The pool returned an error: %s", err.Error())
	}

	return p
}

func createSampleTransaction(cfg *Config) *types.Transaction {

	tx := types.NewTransaction(cfg.PrivateKey, 1,
		cfg.From,
		cfg.To, 1, time.Now())

	return tx
}

func sendTransaction(tx *types.Transaction, cfg *Config, mtr *Meter, pool *grpcpool.Pool) *types.Action {

	byt := []byte(tx.String())
	buffer := new(bytes.Buffer)
	buffer.Write(byt)

	client, err := pool.Get(context.Background())
	if err != nil {
		fmt.Println(err)
	} else {
		st := client.ClientConn.GetState()

		fmt.Println(st.String())

		if client.ClientConn == nil {
			add := fmt.Sprintf("%s:%d", cfg.DelegateIP, 1975)
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

		//utils.Warn(fmt.Sprintf("unable to execute remote delegate [host=%s, port=%d]", contact.Endpoint.Host, contact.Endpoint.Port), err)

	} else {
		action, err := types.ToActionFromJson([]byte(response.Payload))
		if err != nil {
			return types.NewActionWithStatus(actionType, types.StatusInternalError, err.Error())
		}
		return action
	}

	/*url := "http://" + cfg.DelegateIP + ":1975/v1/transactions"
	resp, err := http.Post(url, "application/json", buffer)
	if err != nil {
		fmt.Println(err.Error())
	} else {
	}
	*/
	return nil
}

func run(cfg *Config, mtr *Meter, pool *grpcpool.Pool) {

	ret := createSampleTransaction(cfg)

	//fmt.Println(ret)

	resp := sendTransaction(ret, cfg, mtr, pool)

	//	var d time.Duration
	//	d.Nanoseconds = sleep

	for {

		if resp.Status == types.StatusPending {
			//time.Sleep(d)
		}
		if resp.Status == types.StatusOk {
			break
		}
	}

	if resp.Status == types.StatusOk {
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

func runLoad(cfg *Config, tx int, mtr *Meter, pool *grpcpool.Pool) {

	for i := 0; i <= tx; i++ {
		run(cfg, mtr, pool)
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

	runLoad(&cfg, mtr.Total, &mtr, pool)

	wait()
}
