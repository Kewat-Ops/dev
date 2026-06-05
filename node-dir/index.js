const express = require("express");
const app = express();

app.get("/node", (req, res) => {
  res.send("Hello node");
});

app.listen(3000, () => {
  console.log("Node server running on port 3000");
});
