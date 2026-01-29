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

	if len(req.Payload) != 96 {
		return fmt.Errorf("invalid payload length: expected 96 bytes (hash format), got %d", len(req.Payload))
	}

	tw.logger.Sugar().Info("Task validation successful")
	return nil
}

// HandleTask processa l'HASH dei pesi del FL
func (tw *TaskWorker) HandleTask(req *performerV1.TaskRequest) (*performerV1.TaskResponse, error) {
	tw.logger.Sugar().Infow("Processing FL weights hash task",
		"taskId", hex.EncodeToString(req.TaskId),
	)

	payload := req.Payload
	
	if len(payload) != 96 {
		return nil, fmt.Errorf("invalid payload length: expected 96 bytes, got %d", len(payload))
	}

	// Decode ABI-encoded payload: (bytes32 weightsHash, uint256 clientId, uint256 claimedResult)
	// Struttura ABI:
	// - Offset 0-32: weightsHash (bytes32)
	// - Offset 32-64: clientID (uint256)
	// - Offset 64-96: claimedResult (uint256)
	
	weightsHash := payload[0:32]
	clientID := new(big.Int).SetBytes(payload[32:64])
	claimedResult := new(big.Int).SetBytes(payload[64:96])

	tw.logger.Sugar().Infow("Received FL weights hash",
		"clientID", clientID.String(),
		"claimedResult", claimedResult.String(),
		"weightsHash", hex.EncodeToString(weightsHash),
	)

	// VALIDAZIONE SEMPLICE
	// Controlla che:
	// 1. L'hash non sia vuoto (non tutto zero)
	// 2. Il clientID sia valido (< 1000)
	isValid := true
	validationNote := "Hash accepted"
	
	// Check 1: Hash non vuoto
	allZeros := true
	for _, b := range weightsHash {
		if b != 0 {
			allZeros = false
			break
		}
	}
	
	if allZeros {
		isValid = false
		validationNote = "Empty hash (all zeros) rejected"
		tw.logger.Sugar().Warn(validationNote)
	}
	
	// Check 2: ClientID ragionevole
	if clientID.Cmp(big.NewInt(1000)) > 0 {
		isValid = false
		validationNote = "ClientID too large (>1000) rejected"
		tw.logger.Sugar().Warn(validationNote)
	}

	if isValid {
		tw.logger.Sugar().Infow("✅ FL weights hash VALID",
			"clientID", clientID.String(),
			"weightsHash", hex.EncodeToString(weightsHash)[:16]+"...",
		)
	} else {
		tw.logger.Sugar().Warnw("❌ FL weights hash INVALID",
			"clientID", clientID.String(),
			"reason", validationNote,
		)
	}

	// Codifica risultato come bool (32 bytes)
	result := make([]byte, 32)
	if isValid {
		result[31] = 1 // true
	}

	tw.logger.Sugar().Infow("Task completed successfully",
		"result", isValid,
		"resultBytes", hex.EncodeToString(result),
	)

	return &performerV1.TaskResponse{
		TaskId: req.TaskId,
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

	logger.Info("Starting FL Weights Hash AVS Performer")

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