const express = require('express');
const client = require('prom-client');
const fs = require('fs');

const app = express();

// Read secret at runtime
const db_pass = fs.readFileSync('/run/secrets/db_password', 'utf8').trim();

// Default metrics (CPU, memory, event loop lag, etc.)
const collectDefaultMetrics = client.collectDefaultMetrics;
collectDefaultMetrics();

// Custom counter example
const httpRequestsTotal = new client.Counter({
  name: 'http_requests_total',
  help: 'Total number of HTTP requests',
});

app.get('/node', (req, res) => {
  httpRequestsTotal.inc();
  res.send(`Hello from Node! Secret is: ${db_pass}`);
});

// Expose metrics endpoint
app.get('/metrics', async (req, res) => {
  res.set('Content-Type', client.register.contentType);
  res.end(await client.register.metrics());
});

app.listen(3000, () => {
  console.log('Node service running on port 3000');
});

