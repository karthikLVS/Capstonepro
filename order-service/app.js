const express = require("express");
const app = express();
app.use(express.json());

let orders = [{ id: 1, item: "Laptop", userId: 1 }];

app.get("/orders", (req, res) => res.json(orders));

app.post("/orders", (req, res) => {
  const newOrder = { id: orders.length + 1, ...req.body };
  orders.push(newOrder);
  res.status(201).json(newOrder);
});

app.listen(5001, () => console.log("Order-Service running on port 5001"));
