"""
AVS Bridge - Connessione tra Federated Learning e EigenLayer AVS

Questo modulo gestisce la comunicazione tra il framework FL (Python)
e l'AVS EigenLayer (blockchain) per la validazione dei pesi dei modelli.
"""

from web3 import Web3
from web3.middleware import geth_poa_middleware
import hashlib
import json
import pickle
from typing import List, Dict, Any, Optional
import numpy as np
import os
from pathlib import Path


class AVSBridge:
    """
    Bridge per comunicare con l'AVS EigenLayer per validare i pesi del FL.
    """
    
    def __init__(
        self, 
        rpc_url: str = "http://localhost:8545",
        contract_address: Optional[str] = None,
        private_key: Optional[str] = None
    ):
        """
        Inizializza il bridge AVS.
        
        Args:
            rpc_url: URL del nodo RPC (default: localhost devnet)
            contract_address: Indirizzo del contratto TaskMailbox
            private_key: Chiave privata per firmare transazioni
        """
        # Connessione Web3
        self.w3 = Web3(Web3.HTTPProvider(rpc_url))
        
        # Necessario per reti PoA come devnet
        self.w3.middleware_onion.inject(geth_poa_middleware, layer=0)
        
        # Verifica connessione
        if not self.w3.is_connected():
            raise ConnectionError(f"Impossibile connettersi a {rpc_url}")
        
        print(f"✅ Connesso a blockchain (Chain ID: {self.w3.eth.chain_id})")
        
        # Account
        if private_key:
            self.account = self.w3.eth.account.from_key(private_key)
            print(f"✅ Account: {self.account.address}")
        else:
            # Usa primo account disponibile (devnet)
            self.account = None
            print(f"⚠️  Usando account default del nodo")
        
        # Contratto (da caricare dopo)
        self.contract_address = contract_address
        self.contract = None
        
    def load_contract(self, abi_path: Optional[str] = None):
        """
        Carica il contratto TaskMailbox.
        
        Args:
            abi_path: Path al file ABI JSON (opzionale, cerca automaticamente)
        """
        if not self.contract_address:
            raise ValueError("contract_address non specificato")
        
        # Cerca ABI automaticamente se non specificato
        if not abi_path:
            # Prova path relativo al progetto AVS
            possible_paths = [
                "../my-avs-project/contracts/out/TaskMailbox.sol/TaskMailbox.json",
                "../../my-avs-project/contracts/out/TaskMailbox.sol/TaskMailbox.json",
                "./contracts/out/TaskMailbox.sol/TaskMailbox.json"
            ]
            
            for path in possible_paths:
                if os.path.exists(path):
                    abi_path = path
                    break
        
        if not abi_path or not os.path.exists(abi_path):
            raise FileNotFoundError(
                f"File ABI non trovato. Specifica il path con abi_path"
            )
        
        # Carica ABI
        with open(abi_path) as f:
            contract_json = json.load(f)
            abi = contract_json.get('abi', contract_json)
        
        # Crea istanza contratto
        self.contract = self.w3.eth.contract(
            address=Web3.to_checksum_address(self.contract_address),
            abi=abi
        )
        
        print(f"✅ Contratto caricato: {self.contract_address}")
    
    def compute_weights_hash(self, weights: List[np.ndarray]) -> bytes:
        """
        Calcola hash SHA256 dei pesi del modello.
        
        Args:
            weights: Lista di array numpy (pesi del modello)
            
        Returns:
            Hash SHA256 (32 bytes)
        """
        # Serializza i pesi
        weights_bytes = pickle.dumps(weights)
        
        # Calcola hash
        weights_hash = hashlib.sha256(weights_bytes).digest()
        
        return weights_hash
    
    def submit_weights_for_validation(
        self, 
        weights: List[np.ndarray], 
        client_id: int,
        metadata: Optional[Dict[str, Any]] = None
    ) -> Dict[str, Any]:
        """
        Sottomette i PESI COMPLETI del modello all'AVS per validazione.
        
        Args:
            weights: Lista di array numpy (pesi del modello)
            client_id: ID del client (0-based)
            metadata: Metadati aggiuntivi opzionali
            
        Returns:
            Dict con risultati della validazione
        """
        if not self.contract:
            raise RuntimeError("Contratto non caricato. Chiama load_contract() prima.")
        
        # Serializza i pesi COMPLETI
        weights_bytes = pickle.dumps(weights)
        
        print(f"📤 Sottomissione pesi client {client_id}...")
        print(f"   Dimensione pesi: {len(weights_bytes):,} bytes ({len(weights_bytes)/1024:.2f} KB)")
        
        # Crea payload ABI-encoded: (bytes weightsData, uint256 clientId, uint256 claimedResult)
        from eth_abi import encode
        
        payload = encode(
            ['bytes', 'uint256', 'uint256'],
            [weights_bytes, client_id, 1]  # 1 = "pesi validi"
        )
        
        print(f"   Payload totale (con ABI): {len(payload):,} bytes ({len(payload)/1024:.2f} KB)")
        
        try:
            # Stima gas
            try:
                if self.account:
                    gas_estimate = self.contract.functions.submitTask(payload).estimate_gas({
                        'from': self.account.address
                    })
                else:
                    gas_estimate = self.contract.functions.submitTask(payload).estimate_gas({
                        'from': self.w3.eth.accounts[0]
                    })
                
                # Calcola costo stimato (assumendo 1 gwei = 1e-9 ETH)
                gas_price_gwei = self.w3.eth.gas_price / 1e9
                cost_eth = gas_estimate * gas_price_gwei / 1e9
                
                print(f"   💰 Gas stimato: {gas_estimate:,}")
                print(f"   💰 Costo stimato: {cost_eth:.6f} ETH @ {gas_price_gwei:.2f} gwei")
            except Exception as e:
                print(f"   ⚠️  Impossibile stimare gas: {e}")
                gas_estimate = 10000000  # 10M gas come fallback
            
            # Prepara transazione con gas maggiorato del 20%
            gas_limit = int(gas_estimate * 1.2)
            
            if self.account:
                # Usa account specifico
                tx = self.contract.functions.submitTask(payload).build_transaction({
                    'from': self.account.address,
                    'nonce': self.w3.eth.get_transaction_count(self.account.address),
                    'gas': gas_limit,
                    'gasPrice': self.w3.eth.gas_price
                })
                
                # Firma e invia
                signed_tx = self.account.sign_transaction(tx)
                tx_hash = self.w3.eth.send_raw_transaction(signed_tx.rawTransaction)
            else:
                # Usa account default del nodo
                tx_hash = self.contract.functions.submitTask(payload).transact({
                    'from': self.w3.eth.accounts[0],
                    'gas': gas_limit
                })
            
            print(f"   TX Hash: {tx_hash.hex()[:18]}...")
            
            # Attendi conferma
            receipt = self.w3.eth.wait_for_transaction_receipt(tx_hash, timeout=120)
            
            if receipt['status'] == 1:
                gas_used = receipt['gasUsed']
                cost_actual = gas_used * gas_price_gwei / 1e9
                
                print(f"   ✅ Transazione confermata")
                print(f"   💰 Gas usato: {gas_used:,} (Block: {receipt['blockNumber']})")
                print(f"   💰 Costo effettivo: {cost_actual:.6f} ETH")
                
                return {
                    'valid': True,
                    'tx_hash': tx_hash.hex(),
                    'block_number': receipt['blockNumber'],
                    'client_id': client_id,
                    'gas_used': gas_used,
                    'gas_cost_eth': cost_actual,
                    'weights_size_bytes': len(weights_bytes)
                }
            else:
                print(f"   ❌ Transazione fallita")
                return {
                    'valid': False,
                    'error': 'Transaction failed',
                    'tx_hash': tx_hash.hex()
                }
                
        except Exception as e:
            print(f"   ❌ Errore: {e}")
            return {
                'valid': False,
                'error': str(e),
                'client_id': client_id
            }
    
    def batch_validate_weights(
        self, 
        weights_list: List[List[np.ndarray]]
    ) -> List[Dict[str, Any]]:
        """
        Valida una lista di pesi in batch.
        
        Args:
            weights_list: Lista di liste di pesi (uno per client)
            
        Returns:
            Lista di risultati di validazione
        """
        results = []
        
        print(f"\n🔍 Validazione batch di {len(weights_list)} clients...\n")
        
        for i, weights in enumerate(weights_list):
            result = self.submit_weights_for_validation(weights, client_id=i)
            results.append(result)
            
            # Pausa tra sottomissioni per non sovraccaricare
            if i < len(weights_list) - 1:
                print("   ⏳ Attesa 2s...")
                import time
                time.sleep(2)
        
        # Statistiche
        valid_count = sum(1 for r in results if r.get('valid', False))
        print(f"\n📊 Risultati: {valid_count}/{len(results)} clients validati")
        
        return results


# Esempio di utilizzo
if __name__ == "__main__":
    # Inizializza bridge
    bridge = AVSBridge(
        rpc_url="http://localhost:8545",
        contract_address="0xB99CC53e8db7018f557606C2a5B066527bF96b26"  # Inserisci indirizzo del TaskMailbox
    )
    
    # Carica contratto
    bridge.load_contract()
    
    # Test con pesi finti
    test_weights = [np.random.rand(10, 10) for _ in range(3)]
    
    result = bridge.submit_weights_for_validation(test_weights, client_id=0)
    print(f"\nRisultato: {result}")