#!/usr/bin/env bash

# Function to log to stderr
log() {
  if [ $# -eq 0 ]; then
    echo "" >&2
  else
    echo -e "$@" >&2
  fi
}

function ensureJq() {
    if ! command -v jq &> /dev/null; then
        log "Error: jq not found. Please run 'avs create' first."
        exit 1
    fi
}


function ensureYq() {
    if ! command -v yq &> /dev/null; then
        log "Error: yq not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureMake() {
    if ! command -v make &> /dev/null; then
        log "Error: make not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureDocker() {
    if ! command -v docker &> /dev/null; then
        log "Error: docker not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureRealpath() {
    if ! command -v realpath &> /dev/null; then
        log "Error: realpath not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureForge() {
    if ! command -v forge &> /dev/null; then
        log "Error: forge not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureGomplate() {
    if ! command -v gomplate &> /dev/null; then
        log "Error: gomplate not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureCast() {
    if ! command -v cast &> /dev/null; then
        log "Error: cast not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureAbigen() {
    if ! command -v abigen &> /dev/null; then
        log "Error: abigen not found. Please install with: go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
        exit 1
    fi
}

# Pass in RPC_URL ($1)
function ensureDockerHost() {
    local url="$1"
    
    # Replace localhost with appropriate container names for Docker networking
    if [[ "$url" == *"localhost:8545"* ]]; then
        echo "$url" | sed 's|localhost:8545|devkit-devnet-l1-my-avs-project:8545|g'
    elif [[ "$url" == *"localhost:9545"* ]]; then
        echo "$url" | sed 's|localhost:9545|devkit-devnet-l2-my-avs-project:9545|g'
    elif [[ "$url" == *"127.0.0.1:8545"* ]]; then
        echo "$url" | sed 's|127.0.0.1:8545|devkit-devnet-l1-my-avs-project:8545|g'
    elif [[ "$url" == *"127.0.0.1:9545"* ]]; then
        echo "$url" | sed 's|127.0.0.1:9545|devkit-devnet-l2-my-avs-project:9545|g'
    else
        echo "$url"
    fi
}

# Function to get current nonce from provider for an address
get_current_nonce() {
    local address="$1"
    local rpc_url="$2"
    
    if [ -z "$address" ] || [ -z "$rpc_url" ]; then
        log "Error: get_current_nonce requires address and RPC URL"
        return 1
    fi
    
    # Get nonce from provider
    local nonce=$(cast nonce "$address" --rpc-url "$rpc_url" 2>/dev/null)
    if [ $? -ne 0 ] || [ -z "$nonce" ]; then
        log "Error: Failed to get nonce for address $address"
        return 1
    fi
    
    echo "$nonce"
}

# Function to sync Anvil nonce for an address
sync_anvil_nonce() {
    local private_key="$1"
    local rpc_url="$2"
    local description="${3:-address}"
    
    if [ -z "$private_key" ] || [ -z "$rpc_url" ]; then
        log "Error: sync_anvil_nonce requires private_key and RPC URL"
        return 1
    fi
    
    # Derive address from private key
    local address=$(cast wallet address --private-key "$private_key" 2>/dev/null)
    if [ $? -ne 0 ] || [ -z "$address" ]; then
        log "Error: Failed to derive address from private key"
        return 1
    fi
    
    # Get current nonce from provider
    local current_nonce=$(get_current_nonce "$address" "$rpc_url")
    if [ $? -ne 0 ]; then
        return 1
    fi
    
    log "Syncing nonce for $description ($address): setting to $current_nonce"
    
    # Set nonce using anvil_setNonce RPC call
    local response=$(curl -s -X POST "$rpc_url" \
        -H "Content-Type: application/json" \
        -d "{\"jsonrpc\":\"2.0\",\"method\":\"anvil_setNonce\",\"params\":[\"$address\",\"$current_nonce\"],\"id\":1}")
    
    if [ $? -ne 0 ]; then
        log "Error: Failed to sync nonce via RPC call"
        return 1
    fi
    
    # Check if the response contains an error
    local error=$(echo "$response" | jq -r '.error // empty' 2>/dev/null)
    if [ -n "$error" ]; then
        log "Error: RPC call failed: $error"
        return 1
    fi
    
    log "Successfully synced nonce for $description to $current_nonce"
    return 0
}

# Function to sync nonces and sleep 
sync_nonce_and_sleep() {
    local private_key="$1"
    local rpc_url="$2"
    local description="${3:-deployer}"
    local sleep_duration="${4:-1}"
    
    sync_anvil_nonce "$private_key" "$rpc_url" "$description"
    local sync_result=$?
    
    log "Sleeping for ${sleep_duration} second(s) to ensure nonce propagation..."
    sleep "$sleep_duration"
    
    return $sync_result
}