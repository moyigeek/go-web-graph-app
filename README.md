# Go Web Graph Application

This project is a web application that visualizes a directed graph of nodes with specific criteria. The application is built using Go for the backend and utilizes an SQLite database to store and retrieve graph data.

## Project Structure

```
go-web-graph-app
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── db
│   │   ├── db.go           # Database connection and interaction logic
│   │   └── schema.sql      # SQL schema for the SQLite database
│   ├── handlers
│   │   └── graph.go        # HTTP handlers for graph-related endpoints
│   ├── models
│   │   ├── node.go         # Node struct definition
│   │   └── edge.go         # Edge struct definition
│   └── services
│       └── graph_service.go # Business logic for processing graph data
├── web
│   ├── css
│   │   └── styles.css      # CSS styles for the web application
│   ├── js
│   │   └── graph.js        # JavaScript for rendering the graph
│   └── index.html          # Main HTML file for the web application
├── go.mod                   # Go module file
├── go.sum                   # Checksums for module dependencies
└── README.md                # Project documentation
```

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd go-web-graph-app
   ```

2. **Install Dependencies**
   Ensure you have Go installed. Run the following command to download the necessary dependencies:
   ```bash
   go mod tidy
   ```

3. **Set Up the Database**
   - Create an SQLite database and run the SQL schema defined in `internal/db/schema.sql` to set up the required tables.

4. **Run the Application**
   Start the server by running:
   ```bash
   go run cmd/server/main.go
   ```

5. **Access the Web Application**
   Open your web browser and navigate to `http://localhost:8080` to view the graph visualization.

## Usage

The application retrieves nodes with an indegree greater than 5 and within the range of -100 to 100 on both axes. Users can interact with the graph by moving and zooming to explore the relationships between nodes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.