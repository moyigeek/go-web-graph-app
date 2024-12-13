package models

type Node struct {
    Name    string  `json:"name"`
    X       float64 `json:"x"`
    Y       float64 `json:"y"`
    Indegree int     `json:"indegree"`
}