package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    
    helloworldl1 "github.com/Layr-Labs/hourglass-avs-template/contracts/bindings/l1/helloworldl1"
)

const (
    DefaultRPCURL = "http://localhost:8545"
    PollInterval  = 12 * time.Second
)

type Config struct {
    RPCURL              string
    ServiceManagerAddr  common.Address
    OperatorPrivateKey  string
}

type Operator struct {
    client         *ethclient.Client
    serviceManager *helloworldl1.HelloWorldL1
    config         *Config
    auth           *bind.TransactOpts
}

func main() {
    log.Println("🚀 Starting Sum Verifier AVS Operator...")
    
    config := loadConfig()
    
    operator, err := NewOperator(config)
    if err != nil {
        log.Fatalf("❌ Failed to create operator: %v", err)
    }
    defer operator.client.Close()
    
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    go func() {
        <-sigChan
        log.Println("🛑 Shutdown signal received, stopping operator...")
        cancel()
    }()
    
    log.Println("✅ Operator initialized successfully")
    
    if err := operator.Run(ctx); err != nil {
        log.Fatalf("❌ Operator error: %v", err)
    }
    
    log.Println("👋 Operator stopped gracefully")
}

func loadConfig() *Config {
    rpcURL := os.Getenv("RPC_URL")
    if rpcURL == "" {
        rpcURL = DefaultRPCURL
    }
    
    serviceManagerAddrStr := os.Getenv("SERVICE_MANAGER_ADDRESS")
    if serviceManagerAddrStr == "" {
        serviceManagerAddrStr = "0x0000000000000000000000000000000000000000"
        log.Println("⚠️  Warning: SERVICE_MANAGER_ADDRESS not set, using zero address")
    }
    
    operatorKey := os.Getenv("OPERATOR_PRIVATE_KEY")
    if operatorKey == "" {
        operatorKey = "0x0000000000000000000000000000000000000000000000000000000000000000"
        log.Println("⚠️  Warning: OPERATOR_PRIVATE_KEY not set, using default")
    }
    
    return &Config{
        RPCURL:             rpcURL,
        ServiceManagerAddr: common.HexToAddress(serviceManagerAddrStr),
        OperatorPrivateKey: operatorKey,
    }
}

func NewOperator(config *Config) (*Operator, error) {
    log.Printf("🔌 Connecting to RPC: %s", config.RPCURL)
    
    client, err := ethclient.Dial(config.RPCURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to RPC: %w", err)
    }
    
    log.Printf("📝 Creating HelloWorldL1 instance at: %s", config.ServiceManagerAddr.Hex())
    
    serviceManager, err := helloworldl1.NewHelloWorldL1(config.ServiceManagerAddr, client)
    if err != nil {
        client.Close()
        return nil, fmt.Errorf("failed to create service manager instance: %w", err)
    }
    
    return &Operator{
        client:         client,
        serviceManager: serviceManager,
        config:         config,
    }, nil
}

func (o *Operator) Run(ctx context.Context) error {
    log.Println("👂 Operator started, listening for tasks...")
    
    ticker := time.NewTicker(PollInterval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return nil
        case <-ticker.C:
            if err := o.processTasks(ctx); err != nil {
                log.Printf("❌ Error processing tasks: %v", err)
            }
        }
    }
}

func (o *Operator) processTasks(ctx context.Context) error {
    // Get current task counter
    taskCounter, err := o.serviceManager.TaskCounter(&bind.CallOpts{Context: ctx})
    if err != nil {
        return fmt.Errorf("failed to get task counter: %w", err)
    }
    
    if taskCounter.Uint64() == 0 {
        // No tasks yet, skip logging
        return nil
    }
    
    log.Printf("📊 Current task counter: %s", taskCounter.String())
    
    // Process each unverified task
    maxTasks := taskCounter.Uint64()
    for i := uint64(0); i < maxTasks; i++ {
        if err := o.processTask(ctx, new(big.Int).SetUint64(i)); err != nil {
            log.Printf("❌ Error processing task %d: %v", i, err)
            continue
        }
    }
    
    return nil
}

func (o *Operator) processTask(ctx context.Context, taskID *big.Int) error {
    // Get task details
    task, err := o.serviceManager.GetTask(&bind.CallOpts{Context: ctx}, taskID)
    if err != nil {
        return fmt.Errorf("failed to get task: %w", err)
    }
    
    // Skip if already verified
    if task.Verified {
        return nil
    }
    
    log.Printf("🔍 Processing task %s: %s + %s = %s (claimed)", 
        taskID.String(), task.A.String(), task.B.String(), task.ClaimedResult.String())
    
    // Calculate correct sum
    correctSum := new(big.Int).Add(task.A, task.B)
    isCorrect := correctSum.Cmp(task.ClaimedResult) == 0
    
    if isCorrect {
        log.Printf("✅ Task %s verification: CORRECT (sum = %s)", taskID.String(), correctSum.String())
    } else {
        log.Printf("❌ Task %s verification: INCORRECT (expected %s, got %s)", 
            taskID.String(), correctSum.String(), task.ClaimedResult.String())
    }
    
    // TODO: Submit verification result to L1
    // This requires setting up the TransactOpts with the operator's private key
    
    return nil
}