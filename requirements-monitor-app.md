# Project Requirements: Monitor Application with Pipe Communication

## Overview

This project aims to develop a monitoring application in Go that handles logs from multiple processes through pipe communication. The application will also feature a frontend interface to manage processes and display logs in a user-friendly manner. Logs will be stored in daily files, and the application will support different log levels (e.g., DEBUG, INFO, WARN, ERROR, CRITICAL), each styled differently based on importance.

## Functional Requirements

1. **Log Handling**:
   - Processes will write logs to named pipes.
   - The monitoring application will read logs from these pipes.
   - Logs will be categorized by type: DEBUG, INFO, WARN, ERROR, CRITICAL.
   - Logs will be stored in daily log files.

2. **Named Pipes Communication**:
   - Each process creates a named pipe for log communication.
   - The process notifies the monitoring application to start listening on the pipe.
   - The monitoring application dynamically adds the new pipe and starts reading logs.

3. **Log Storage**:
   - Logs will be written to daily files, named in the format `logs/YYYY-MM-DD.log`.
   - Each log entry will include a timestamp, log level, process identifier, and message.

4. **Frontend Interface**:
   - A web-based frontend to display process status and logs.
   - Buttons for each process to indicate if the process is connected and selected.
   - When a button is selected, logs from the corresponding process will be displayed in the monitor and written to the file.
   - If a button is not selected, logs will only be written to the file.

5. **Log Styling**:
   - Different styles for each log level in the monitor.
     - DEBUG: Less prominent styling (e.g., gray text).
     - INFO: Normal styling (e.g., black text).
     - WARN: Attention-grabbing styling (e.g., yellow text).
     - ERROR: Important styling (e.g., red text).
     - CRITICAL: Most prominent styling (e.g., bold red text with background).

## Non-Functional Requirements

1. **Performance**:
   - The application should handle high log input efficiently.
   - The frontend should update log displays in real-time with minimal latency.

2. **Scalability**:
   - The system should support monitoring a large number of processes concurrently.

3. **Reliability**:
   - The application should ensure log messages are not lost.
   - It should handle pipe communication errors gracefully.

4. **Usability**:
   - The frontend should be intuitive and user-friendly.
   - Buttons and log display should be responsive and clearly indicate process status and log levels.

## Technical Requirements

1. **Backend**:
   - Language: Go
   - Libraries: 
     - `github.com/labstack/echo/v4` for the web server and frontend interface.
2. **Frontend**:
   - Framework: React.js (for dynamic and responsive UI).

3. **File System**:
   - Daily log files will be stored in a specified directory.

## System Architecture

1. **Processes**:
   - Write logs to named pipes.
   - Notify the monitoring application to listen on new pipes via Unix sockets.

2. **Monitoring Application**:
   - Unix socket server/Master Pipe to receive notifications of new pipes.
   - Goroutines to listen on multiple pipes concurrently.
   - Parse log messages and categorize by log level.
   - Write logs to daily files.
   - Serve logs to the frontend via a web server (Echo).

3. **Frontend Application**:
   - Display process status and logs.
   - Buttons to select/unselect processes.
   - Real-time updates of logs with different styles based on log levels.

## Example Log Message Format
    [INFO|15:04:05|ProcessName] This is an info log message
    [ERROR|15:04:05|ProcessName] This is an error log message
    [DEBUG|15:04:05|ProcessName] This is a debug log message
    [WARN|15:04:05|ProcessName] This is a warning log message
    [CRITICAL|15:04:05|ProcessName] This is a critical log message



