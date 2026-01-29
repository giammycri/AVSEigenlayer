#!/usr/bin/env python3
import grpc
import time

# Crea payload semplice
payload = b'\x00' * 96  # 96 bytes di zeri

for i, port in enumerate([9090, 9091, 9092], 1):
    print(f"\n📤 Test Executor {i} (porta {port})...")
    
    try:
        channel = grpc.insecure_channel(f'localhost:{port}')
        
        # Verifica connessione
        grpc.channel_ready_future(channel).result(timeout=5)
        print(f"   ✅ Connesso a executor{i}")
        
        # Qui dovresti chiamare il metodo gRPC corretto
        # Ma serve il proto file compilato
        
    except Exception as e:
        print(f"   ❌ Errore: {e}")
    
    time.sleep(2)
