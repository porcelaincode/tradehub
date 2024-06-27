# TradeHub: Mock Broker Server

This project is a mock broker server built with FastAPI, allowing you to paper trade using simulated market data. The server handles instruments, orders, positions, and emits relevant events via WebSockets.

## Features

- Downloads instruments from an external API and stores them in an SQLite database.
- Provides endpoints for subscribing to instrument tokens and placing orders.
- Simulates live market data and emits events for subscribed instruments.
- Handles orders and updates positions, including GTT and regular orders.

## Endpoints

### POST /subscribe

Subscribe to an array of instrument tokens to receive simulated market data.

Request body:

```json
[{ "instrument_token": "11536", "exchange_segment": "nse_cm" }]
```

### POST /orders

Place an order with the specified parameters.

Request body:

```json
{
  "am": "YES",
  "dq": "0",
  "es": "nse_cm",
  "mp": "0",
  "pc": "NRML",
  "pf": "N",
  "pr": "100.5",
  "pt": "L",
  "qt": "10",
  "rt": "DAY",
  "tp": "0",
  "ts": "YESBANK-EQ",
  "tt": "B",
  "ig": "order_tag"
}
```

### GET /positions

Retrieve the current positions.

Response:

```json
{
  "stat": "Ok",
  "stCode": 200,
  "data": [
    {
      "buyAmt": "2625.00",
      "cfSellAmt": "0.00",
      "prod": "NRML",
      "exSeg": "nse_fo",
      "sqrFlg": "Y",
      "actId": "PRS2206",
      "cfBuyQty": "0",
      "cfSellQty": "0",
      "tok": "53179",
      "flBuyQty": "25",
      "flSellQty": "25",
      "sellAmt": "2625.00",
      "posFlg": "true",
      "cfBuyAmt": "0.00",
      "stkPrc": "0.00",
      "trdSym": "BANKNIFTY21JULFUT",
      "sym": "BANKNIFTY",
      "expDt": "29 Jul, 2021",
      "type": "FUTIDX",
      "series": "XX",
      "brdLtQty": "25",
      "exp": "1627569000",
      "optTp": "XX",
      "genNum": "1",
      "genDen": "1",
      "prcNum": "1",
      "prcDen": "1",
      "lotSz": "25",
      "multiplier": "1",
      "precision": "2",
      "hsUpTm": "2021/07/13 18:34:44"
    }
  ]
}
```

## TODO

<input type="checkbox" disabled /> Write better random LTP generation logic

<input type="checkbox" disabled /> Move the database to Aerospike or Redis for better performance

<input type="checkbox" disabled /> <b>Handle concurrency and transaction Handling</b> such that all database operations are performed within proper transactions to maintain data integrity.

<input type="checkbox" disabled /> <b>Add detailed error handling</b> and logging to capture and debug issues effectively.

<input type="checkbox" disabled /> Implement <b>necessary security measures</b> such as authentication and authorization if required.

<input type="checkbox" disabled /> <b>Optimize database queries</b> and WebSocket handling for better performance under load.

<input type="checkbox" disabled checked/> Write the entire thing in Go lang
