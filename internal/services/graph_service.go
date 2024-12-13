package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/moyigeek/go-web-graph-app/internal/models"
)

type GraphService struct {
	db *sql.DB
}

func NewGraphService(db *sql.DB) *GraphService {
	return &GraphService{db: db}
}

func (gs *GraphService) GetFilteredNodes() ([]models.Node, error) {
	query := `
		SELECT Name, X, Y, Indegree 
		FROM node_info 
		WHERE Indegree > 5 AND X BETWEEN -100 AND 100 AND Y BETWEEN -100 AND 100
	`
	rows, err := gs.db.Query(query)
	if err != nil {
		log.Println("Error querying nodes:", err)
		return nil, err
	}
	defer rows.Close()

	var nodes []models.Node
	for rows.Next() {
		var node models.Node
		if err := rows.Scan(&node.Name, &node.X, &node.Y, &node.Indegree); err != nil {
			log.Println("Error scanning node:", err)
			return nil, err
		}
		nodes = append(nodes, node)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return nodes, nil
}

func (gs *GraphService) GetEdges() ([]models.Edge, error) {
	query := `
		SELECT From, To, X1, Y1, X2, Y2 
		FROM edge_info
	`
	rows, err := gs.db.Query(query)
	if err != nil {
		log.Println("Error querying edges:", err)
		return nil, err
	}
	defer rows.Close()

	var edges []models.Edge
	for rows.Next() {
		var edge models.Edge
		if err := rows.Scan(&edge.From, &edge.To, &edge.X1, &edge.Y1, &edge.X2, &edge.Y2); err != nil {
			log.Println("Error scanning edge:", err)
			return nil, err
		}
		edges = append(edges, edge)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return edges, nil
}