# 🖥️ Go Server Health Monitor

A lightweight, CLI-based system health monitor written in Go (Golang). It fetches live CPU, RAM, and Disk usage statistics directly from the host machine.

## 🛠️ Tech Stack
- **Language:** Go (Golang)
- **Library:** `gopsutil` (for system metrics)
- **OS:** Windows (tested) / Linux / macOS

## 🚀 Features
- Fetches real-time CPU usage percentage.
- Fetches real-time RAM usage and calculates used memory in MB.
- Fetches Disk usage for the C: drive.
- (Optional) Alerts via Telegram if usage goes above a certain threshold.

## 🏃‍♂️ How to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/nirnayabikramshah/server-health-monitor.git
