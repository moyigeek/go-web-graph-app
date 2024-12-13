package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
}

func GetNodesWithIndegreeGreaterThan(minIndegree int) ([]Node, error) {
	rows, err := db.Query("SELECT Name, X, Y, Indegree FROM node_info WHERE Indegree > ? AND X BETWEEN -100 AND 100 AND Y BETWEEN -100 AND 100", minIndegree)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []Node
	for rows.Next() {
		var node Node
		if err := rows.Scan(&node.Name, &node.X, &node.Y, &node.Indegree); err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func GetEdges() ([]Edge, error) {
	rows, err := db.Query("SELECT From, To, X1, Y1, X2, Y2 FROM edge_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var edges []Edge
	for rows.Next() {
		var edge Edge
		if err := rows.Scan(&edge.From, &edge.To, &edge.X1, &edge.Y1, &edge.X2, &edge.Y2); err != nil {
			return nil, err
		}
		edges = append(edges, edge)
	}
	return edges, nil
}