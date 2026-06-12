const fs = require('fs');
const express = require('express');
const client = require('prom-client');

// OpenTelemetry imports
const { NodeSDK } = require('@opentelemetry/sdk-node');
const { getNodeAutoInstrumentations } = require('@opentelemetry/auto-instrumentations-node');
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');

const app = express();

// Read secret at runtime
const db_pass = fs.readFileSync('/run/secrets/db_password', 'utf8').trim();

// Prometheus metrics
const collectDefaultMetrics = client.collectDefaultMetrics;
collectDefaultMetrics();
const httpRequestsTotal = new client.Counter({
  name: 'http_requests_total',
  help: 'Total number of HTTP requests',
});

// OpenTelemetry setup
const sdk = new NodeSDK({
  traceExporter: new JaegerExporter({ host: 'jaeger', port: 6831 }),
  instrumentations: [getNodeAutoInstrumentations()]
});
sdk.start();

app.get('/node', (req, res) => {
  httpRequestsTotal.inc();
  res.send(`Hello from Node with tracing! Secret is: ${db_pass}`);
});

// Expose metrics endpoint
app.get('/metrics', async (req, res) => {
  res.set('Content-Type', client.register.contentType);
  res.end(await client.register.metrics());
});

app.listen(3000, () => {
  console.log('Node service running on port 3000');
});

