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

// HandleTask processa i PESI COMPLETI del FL
func (tw *TaskWorker) HandleTask(req *performerV1.TaskRequest) (*performerV1.TaskResponse, error) {
	tw.logger.Sugar().Infow("Processing FL weights task",
		"taskId", hex.EncodeToString(req.TaskId),
	)

	payload := req.Payload
	
	if len(payload) < 96 {
		return nil, fmt.Errorf("invalid payload length: %d", len(payload))
	}

	// Decode ABI-encoded payload: (bytes weightsData, uint256 clientId, uint256 claimedResult)
	// Struttura ABI:
	// - Offset 0-32: offset ai bytes dei pesi (tipicamente 96 = 0x60)
	// - Offset 32-64: clientID (uint256)
	// - Offset 64-96: claimedResult (uint256)
	// - Offset 96+: lunghezza e dati effettivi dei pesi
	
	weightsOffset := new(big.Int).SetBytes(payload[0:32]).Uint64()
	clientID := new(big.Int).SetBytes(payload[32:64])
	claimedResult := new(big.Int).SetBytes(payload[64:96])
	
	// Leggi i pesi effettivi
	var weightsBytes []byte
	var weightsSize uint64
	
	if weightsOffset < uint64(len(payload)) {
		// Lunghezza dei pesi (32 bytes all'offset)
		weightsSize = new(big.Int).SetBytes(payload[weightsOffset:weightsOffset+32]).Uint64()
		
		// Dati effettivi dei pesi
		dataStart := weightsOffset + 32
		dataEnd := dataStart + weightsSize
		
		if dataEnd <= uint64(len(payload)) {
			weightsBytes = payload[dataStart:dataEnd]
		}
	}

	tw.logger.Sugar().Infow("Received FL weights",
		"clientID", clientID.String(),
		"claimedResult", claimedResult.String(),
		"weightsSize", weightsSize,
		"weightsSizeKB", float64(weightsSize)/1024.0,
	)

	// VALIDAZIONE SEMPLICE
	// Controlla che:
	// 1. I pesi non siano vuoti
	// 2. I pesi non superino 10MB (limite ragionevole)
	isValid := true
	validationNote := "Weights accepted"
	
	if len(weightsBytes) == 0 {
		isValid = false
		validationNote = "Empty weights rejected"
	} else if len(weightsBytes) > 10*1024*1024 {
		isValid = false
		validationNote = "Weights too large (>10MB) rejected"
	}

	tw.logger.Sugar().Infow("FL weights validated",
		"isValid", isValid,
		"note", validationNote,
	)

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

	logger.Info("Starting FL Weights AVS Performer")

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