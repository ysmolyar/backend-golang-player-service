# ⚾ Player Service

Player Service is a backend application that serves baseball player data. In addition, Player service integrates with [Ollama](https://github.com/ollama/ollama/blob/main/docs/api.md), which allows us to run the [tinyllama LLM](https://ollama.com/library/tinyllama) locally.

## Dependencies

- [Go 1.21+](https://golang.org/dl/)
- [Make](https://www.gnu.org/software/make/)
- [Docker](https://www.docker.com/)

## 🛠️ Setup Instructions

1. Verify system dependencies
   1. Go
      - Verify installation: `go version`
   2. Make
      - Verify installation: `make --version`
   3. Docker
      - Download and install from [docker.com](https://www.docker.com/)
      - Verify installation: `docker --version`

2. Clone this repository or Download the code as zip
   - run `git clone https://github.com/your-org/backend-golang-player-service.git`

## Run the application

### Part 1: Application Dependencies

1. Install application dependencies
    - From the project's root directory, run: `make deps`

### Part 2: Run Player Service

1. Start the Player service
   ```bash
   make run
   ```

2. Verify the Player service is running
   1. Open your browser and visit `http://localhost:8080/v1/players`
   2. If the application is running successfully, you will see player data appear in the browser

### Part 3: Start LLM Docker Container

Player service integrates with Ollama 🦙, which allows us to run LLMs locally. This app runs [tinyllama](https://ollama.com/library/tinyllama) model.

1. Pull and run Ollama docker image
   ```bash
   docker pull ollama/ollama
   ```

2. Run Ollama docker image
   ```bash
   docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama
   ```

3. Download and run tinyllama model
   ```bash
   docker exec -it ollama ollama run tinyllama
   ```

4. Test Ollama API server
   ```bash
   curl -v --location 'http://localhost:11434/api/generate' \
   --header 'Content-Type: application/json' \
   --data '{
       "model": "tinyllama",
       "prompt": "why is the sky blue?",
       "stream": false
   }'
   ```

### Part 4: Verify Player Service and LLM Integration

1. Ensure Player Service is running
2. Visit `http://localhost:8080/v1/chat/list-models`
   - If the application is running successfully, you will see information about tinyllama

## API Endpoints

### Player Endpoints
- `GET /v1/players` - List all players
- `GET /v1/players/:id` - Get a specific player

### Chat Endpoints
- `GET /v1/chat/list-models` - List available LLM models
- `POST /v1/chat/generate` - Generate text using the LLM