# GoBoard

GoBoard is a lightweight and native Go framework built from scratch.  
It is designed for learning Go fundamentals and can also serve as a minimal project template.

## Features
- Pure Go, no heavy frameworks
- Web server (HTTP APIs)
- Database integration (SQLite/MySQL/PostgreSQL)
- Local file operations
- Modular design for easy extension

## Getting Started

### Prerequisites
- Go 1.22 or above
- Git

### Installation
```bash
git clone https://github.com/yourname/GoBoard.git
cd GoBoard
go run main.go

## Project Structure
GoBoard/
├── main.go          # Entry point
├── handlers/        # API handlers
├── middlewares/     # Middleware
├── tools/           # Utility functions
├── templates/       # web templates
├── db/        # DB related code
└── README.md

## Example
Run the project and visit:
http://localhost:8080/hello

## Roadmap
 ### Add user authentication
 ### RESTful API examples
 ### More database drivers support
 ### Docker deployment

## License
MIT License
