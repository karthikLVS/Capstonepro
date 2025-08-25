const request = require("supertest");
const express = require("express");

let app = express();
app.use(express.json());

let orders = [{ id: 1, item: "Laptop", userId: 1 }];
app.get("/orders", (req, res) => res.json(orders));

test("GET /orders should return list", async () => {
  const res = await request(app).get("/orders");
  expect(res.statusCode).toBe(200);
  expect(Array.isArray(res.body)).toBe(true);
});
