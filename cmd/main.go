package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"go.uber.org/zap"
)

type TaskWorker struct {
	logger *zap.Logger
}

func NewTaskWorker(logger *zap.Logger) *TaskWorker {
	return &TaskWorker{
		logger: logger,
	}
}

// ValidateTask valida il task prima dell'esecuzione
func (tw *TaskWorker) ValidateTask(req *performerV1.TaskRequest) error {
	tw.logger.Sugar().Infow("Validating task",
		"taskId", hex.EncodeToString(req.TaskId),
		"payloadLength", len(req.Payload),
	)

	if len(req.Payload) == 0 {
		return fmt.Errorf("empty payload")
	}

	if len(req.Payload) < 96 {
		return fmt.Errorf("invalid payload length: expected at least 96 bytes, got %d", len(req.Payload))
	}

	tw.logger.Sugar().Info("Task validation successful")
	return nil
}

// HandleTask processa il task e verifica la somma
func (tw *TaskWorker) HandleTask(req *performerV1.TaskRequest) (*performerV1.TaskResponse, error) {
	tw.logger.Sugar().Infow("Processing sum verification task",
		"taskId", hex.EncodeToString(req.TaskId),
	)

	payload := req.Payload
	
	if len(payload) < 96 {
		return nil, fmt.Errorf("invalid payload length: %d", len(payload))
	}

	// Parse i valori dal payload (3 uint256)
	// Offset 0-32: a
	// Offset 32-64: b
	// Offset 64-96: claimedResult
	
	a := new(big.Int).SetBytes(payload[0:32])
	b := new(big.Int).SetBytes(payload[32:64])
	claimedResult := new(big.Int).SetBytes(payload[64:96])

	tw.logger.Sugar().Infow("Parsed task data",
		"a", a.String(),
		"b", b.String(),
		"claimedResult", claimedResult.String(),
	)

	// Calcola il risultato corretto
	correctSum := new(big.Int).Add(a, b)
	isCorrect := correctSum.Cmp(claimedResult) == 0

	tw.logger.Sugar().Infow("Sum verification result",
		"correctSum", correctSum.String(),
		"isCorrect", isCorrect,
	)

	// Codifica il risultato come ABI-encoded bool (32 bytes)
	// Un bool in ABI encoding è sempre 32 bytes con il valore nell'ultimo byte
	result := make([]byte, 32)
	if isCorrect {
		result[31] = 1 // true - ultimo byte = 1
	}
	// false è già rappresentato da tutti 0

	tw.logger.Sugar().Infow("Task processed successfully",
		"result", isCorrect,
		"resultBytes", hex.EncodeToString(result),
	)

	// Ritorna la risposta con il risultato a 32 bytes
	return &performerV1.TaskResponse{
		Result: result,
	}, nil
}

// main avvia il server gRPC Hourglass Performer
func main() {
	ctx := context.Background()
	
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}
	defer logger.Sync()

	logger.Info("Starting Sum Verification AVS Performer")

	worker := NewTaskWorker(logger)

	performer, err := server.NewPonosPerformerWithRpcServer(&server.PonosPerformerConfig{
		Port:    8080,
		Timeout: 5 * time.Second,
	}, worker, logger)
	if err != nil {
		panic(fmt.Errorf("failed to create performer: %w", err))
	}

	logger.Info("Performer server created, starting on port 8080")

	if err := performer.Start(ctx); err != nil {
		panic(fmt.Errorf("failed to start performer: %w", err))
	}
}