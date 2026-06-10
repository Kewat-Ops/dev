const fs = require('fs');
const express = require('express');
const app = express();

// Read secret at runtime
const db_pass = fs.readFileSync('/run/secrets/db_password', 'utf8').trim();
console.log("DB password:", db_pass);  // demo only

app.get('/node', (req, res) => {
  res.send(`Hello from Node! Secret is: ${db_pass}`);
});

app.listen(3000, () => {
  console.log('Server running on port 3000');
});

