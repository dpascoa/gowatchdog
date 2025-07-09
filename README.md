# 🐶 GoWatchdog

GoWatchdog is a lightweight uptime monitoring tool built in Go.  
It checks websites periodically, logs their status, and serves a clean web dashboard with auto-refresh.

![image](https://github.com/user-attachments/assets/99aaf719-9603-494a-b30d-bce0dfef5975)

## Features
- ✅ Website status checker (200 OK / DOWN)
- ⚡ Concurrent checks using goroutines
- 📝 Log results to a file (`gowatchdog.log`)
- 🌐 Simple HTML dashboard with live countdown refresh
- 🔧 Built entirely in Go (no dependencies)

## Usage

1. Clone the repo  
2. Add your sites in `config.json`  
3. Run it:

```bash
go run main.go
```

Then open http://localhost:8080 to view the dashboard.
___

## Example config.json
```bash
{
  "interval_seconds": 30,
  "urls": [
    "https://www.google.com",
    "https://developer.mozilla.org/en-US/"
  ]
}
```
___
## Future Features
- Email/Slack alerts
- Historical timeline
- Docker support

___
Made with 💙 in Go
 
