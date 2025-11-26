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

	// Verifica che il payload non sia vuoto
	if len(req.Payload) == 0 {
		return fmt.Errorf("empty payload")
	}

	// Il payload dovrebbe essere lungo almeno 128 bytes (4 uint256)
	// uint256 taskId + uint256 a + uint256 b + uint256 claimedResult
	if len(req.Payload) < 128 {
		return fmt.Errorf("invalid payload length: expected at least 128 bytes, got %d", len(req.Payload))
	}

	tw.logger.Sugar().Info("Task validation successful")
	return nil
}

// HandleTask processa il task e verifica la somma
func (tw *TaskWorker) HandleTask(req *performerV1.TaskRequest) (*performerV1.TaskResponse, error) {
	tw.logger.Sugar().Infow("Processing sum verification task",
		"taskId", hex.EncodeToString(req.TaskId),
	)

	// Decodifica il payload: (taskId, a, b, claimedResult)
	payload := req.Payload
	
	// Il payload è ABI-encoded, quindi ogni uint256 occupa 32 bytes
	if len(payload) < 128 {
		return nil, fmt.Errorf("invalid payload length: %d", len(payload))
	}

	// Parse i valori dal payload
	// Offset 0-32: taskId (non ci serve qui)
	// Offset 32-64: a
	// Offset 64-96: b
	// Offset 96-128: claimedResult
	
	a := new(big.Int).SetBytes(payload[32:64])
	b := new(big.Int).SetBytes(payload[64:96])
	claimedResult := new(big.Int).SetBytes(payload[96:128])

	tw.logger.Sugar().Infow("Parsed task data",
		"a", a.String(),
		"b", b.String(),
		"claimedResult", claimedResult.String(),
	)

	// Calcola il risultato corretto
	correctSum := new(big.Int).Add(a, b)
	
	// Verifica se il risultato è corretto
	isCorrect := correctSum.Cmp(claimedResult) == 0

	tw.logger.Sugar().Infow("Sum verification result",
		"correctSum", correctSum.String(),
		"isCorrect", isCorrect,
	)

	// Codifica il risultato (bool) in un byte array
	var result []byte
	if isCorrect {
		result = []byte{1} // true
	} else {
		result = []byte{0} // false
	}

	tw.logger.Sugar().Infow("Task processed successfully",
		"result", isCorrect,
	)

	// Ritorna la risposta con il risultato della verifica
	return &performerV1.TaskResponse{
		Result: result,
	}, nil
}

// main avvia il server gRPC Hourglass Performer
func main() {
	ctx := context.Background()
	
	// Inizializza logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}
	defer logger.Sync()

	logger.Info("Starting Sum Verification AVS Performer")

	// Crea il TaskWorker con la logica custom
	worker := NewTaskWorker(logger)

	// Crea il server Hourglass Performer sulla porta 8080
	performer, err := server.NewPonosPerformerWithRpcServer(&server.PonosPerformerConfig{
		Port:    8080,
		Timeout: 5 * time.Second,
	}, worker, logger)
	if err != nil {
		panic(fmt.Errorf("failed to create performer: %w", err))
	}

	logger.Info("Performer server created, starting on port 8080")

	// Avvia il server (blocking call)
	if err := performer.Start(ctx); err != nil {
		panic(fmt.Errorf("failed to start performer: %w", err))
	}
}