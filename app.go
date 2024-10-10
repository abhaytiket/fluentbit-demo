package main

import (
    "encoding/json"
    "log"
    "os"
    "time"
)

type LogEntry struct {
    Timestamp string `json:"timestamp"`
    SourceIp  string `json:"sourceIp"`
    SourceIdentity string `json:"sourceIdentity"`
    TargetIdentity string `json:"targetIdentity"`
    EventType string `json:"eventType"`
    ActivityResult string `json:"activityResult"`
}

func main() {
    // Open or create a log file
    file, err := os.OpenFile("/var/log/app/timestamp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer file.Close()

    logger := log.New(file, "", 0)

    // Infinite loop to log timestamp every 5 second
    for {
        entry := LogEntry{
            Timestamp: time.Now().Format(time.RFC3339),
            SourceIp: "192.168.0.1",
            SourceIdentity: "user1",
            TargetIdentity: "user2",
            EventType: "login",
            ActivityResult: "success",
        }

        jsonEntry, err := json.Marshal(entry)
        if err != nil {
            log.Fatalf("Failed to marshal JSON: %v", err)
        }

        // Log to the file in JSON format
        logger.Println(string(jsonEntry))

        time.Sleep(5 * time.Second)
    }
}
