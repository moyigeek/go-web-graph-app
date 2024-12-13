const svgWidth = 800;
const svgHeight = 600;
const zoomScaleFactor = 1.2;

let svg, g, zoom;

function initGraph() {
    svg = d3.select("#graph")
        .append("svg")
        .attr("width", svgWidth)
        .attr("height", svgHeight)
        .call(d3.zoom().scaleExtent([0.5, 5]).on("zoom", zoomed));

    g = svg.append("g");

    fetchNodesAndEdges();
}

function zoomed(event) {
    g.attr("transform", event.transform);
}

function fetchNodesAndEdges() {
    fetch('/api/graph')
        .then(response => response.json())
        .then(data => {
            drawGraph(data.nodes, data.edges);
        })
        .catch(error => console.error('Error fetching graph data:', error));
}

function drawGraph(nodes, edges) {
    const link = g.selectAll(".link")
        .data(edges)
        .enter().append("line")
        .attr("class", "link")
        .attr("stroke", "#999")
        .attr("stroke-opacity", 0.6);

    const node = g.selectAll(".node")
        .data(nodes)
        .enter().append("g")
        .attr("class", "node")
        .attr("transform", d => `translate(${d.X},${d.Y})`)
        .call(d3.drag()
            .on("start", dragstarted)
            .on("drag", dragged)
            .on("end", dragended));

    node.append("circle")
        .attr("r", 5)
        .attr("fill", "#69b3a2");

    node.append("text")
        .attr("dx", 12)
        .attr("dy", ".35em")
        .text(d => d.Name);

    simulation(nodes, edges);
}

function simulation(nodes, edges) {
    const simulation = d3.forceSimulation(nodes)
        .force("link", d3.forceLink().id(d => d.Name).distance(50))
        .force("charge", d3.forceManyBody().strength(-100))
        .force("center", d3.forceCenter(svgWidth / 2, svgHeight / 2));

    simulation.on("tick", () => {
        g.selectAll(".link")
            .attr("x1", d => d.source.X)
            .attr("y1", d => d.source.Y)
            .attr("x2", d => d.target.X)
            .attr("y2", d => d.target.Y);

        g.selectAll(".node")
            .attr("transform", d => `translate(${d.X},${d.Y})`);
    });
}

function dragstarted(event, d) {
    if (!event.active) simulation.alphaTarget(0.3).restart();
    d.fx = d.x;
    d.fy = d.y;
}

function dragged(event, d) {
    d.fx = event.x;
    d.fy = event.y;
}

function dragended(event, d) {
    if (!event.active) simulation.alphaTarget(0);
    d.fx = null;
    d.fy = null;
}

document.addEventListener("DOMContentLoaded", initGraph);