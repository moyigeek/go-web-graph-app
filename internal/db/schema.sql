CREATE TABLE node_info (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    x REAL NOT NULL,
    y REAL NOT NULL,
    indegree INTEGER NOT NULL
);

CREATE TABLE edge_info (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    from_node TEXT NOT NULL,
    to_node TEXT NOT NULL,
    x1 REAL NOT NULL,
    y1 REAL NOT NULL,
    x2 REAL NOT NULL,
    y2 REAL NOT NULL,
    FOREIGN KEY (from_node) REFERENCES node_info(name),
    FOREIGN KEY (to_node) REFERENCES node_info(name)
);