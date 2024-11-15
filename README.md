# Bank Transaction Deadlock Example in Go

## Overview
This project demonstrates a deadlock scenario in a banking system where transfers are made between accounts concurrently. It includes an example that leads to a deadlock and a solution that prevents it by maintaining a consistent locking order.

## Files
- **main.go**: Contains the code for a bank transfer that can cause a deadlock and a modified version to avoid it.

## Deadlock Scenario
In the `Transfer` function:
- Goroutine 1 locks **Account 1** and tries to lock **Account 2**.
- Goroutine 2 locks **Account 2** and tries to lock **Account 1**.
  
This results in a deadlock since each goroutine waits indefinitely for the account the other holds.

### Output Example
```plaintext
Initiating transfer of 100 from Account 1 to Account 2
Locked Account 1
Initiating transfer of 50 from Account 2 to Account 1
Locked Account 2

