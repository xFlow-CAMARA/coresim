#!/bin/bash
# Script to keep CoreSim simulation running
# Run this in the background to automatically restart the simulation when it stops

CORESIM_API="http://localhost:8081/core-simulator/v1"
CHECK_INTERVAL=10  # seconds

echo "Starting CoreSim auto-restart monitor..."
echo "Checking every ${CHECK_INTERVAL} seconds"

while true; do
    # Check simulation status
    STATUS=$(curl -s "${CORESIM_API}/status" | jq -r '.Status' 2>/dev/null)
    
    if [ "$STATUS" == "STOPPED" ] || [ "$STATUS" == "ERROR" ]; then
        echo "[$(date '+%Y-%m-%d %H:%M:%S')] Simulation is $STATUS, restarting..."
        
        # Start the simulation
        RESULT=$(curl -s -X POST "${CORESIM_API}/start" | jq -r '.Status' 2>/dev/null)
        
        if [ "$RESULT" == "STARTED" ]; then
            echo "[$(date '+%Y-%m-%d %H:%M:%S')] ✓ Simulation restarted successfully"
        else
            echo "[$(date '+%Y-%m-%d %H:%M:%S')] ✗ Failed to restart simulation: $RESULT"
        fi
    elif [ "$STATUS" == "STARTED" ]; then
        echo "[$(date '+%Y-%m-%d %H:%M:%S')] Simulation is running"
    else
        echo "[$(date '+%Y-%m-%d %H:%M:%S')] Unknown status: $STATUS"
    fi
    
    sleep $CHECK_INTERVAL
done
