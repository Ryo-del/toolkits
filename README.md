# devops-toolkit

## Overview
The DevOps Toolkit is a command-line and graphical user interface (GUI) application designed to assist with various DevOps tasks, including log parsing, system monitoring, and port scanning. This toolkit aims to provide users with efficient tools to manage and analyze system resources and logs.

## Features
- **CLI Commands**:
  - Log Parser: Processes log files and extracts relevant information.
  - Monitor: Monitors system resources and outputs the current status.
  - Port Scanner: Scans specified ports and reports their status.

- **GUI Application**:
  - User-friendly interface built with Fyne.
  - Components for log parsing, monitoring, and port scanning.

## Installation
1. Clone the repository:
   ```
   git clone https://github.com/yourusername/devops-toolkit.git
   ```
2. Navigate to the project directory:
   ```
   cd devops-toolkit
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```

## Usage
### CLI
To run the CLI application, execute the following command:
```
go run cmd/main.go
```
You can use the following commands:
- `logparser`: To parse log files.
- `monitor`: To monitor system resources.
- `portscanner`: To scan ports.

### GUI
To run the GUI application, execute:
```
go run ui/main.go
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.