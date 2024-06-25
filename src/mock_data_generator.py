import asyncio
import random
from websocket_manager import websocket_manager

async def generate_mock_tick_data(instrument_token):
    while True:
        if instrument_token in websocket_manager.subscribed_tokens:
            tick_data = {
                "message": [{
                    "last_traded_time": "28/06/2023 15:59:47",
                    "volume": str(random.randint(1, 1000000)),
                    "last_traded_price": str(round(random.uniform(10, 100), 2)),
                    "last_traded_quantity": str(random.randint(1, 100)),
                    "total_buy_quantity": str(random.randint(1, 1000000)),
                    "total_sell_quantity": str(random.randint(1, 1000000)),
                    "buy_price": str(round(random.uniform(10, 100), 2)),
                    "sell_price": str(round(random.uniform(10, 100), 2)),
                    "buy_quantity": str(random.randint(1, 1000000)),
                    "average_price": str(round(random.uniform(10, 100), 2)),
                    "lower_circuit_limit": "13.00",
                    "upper_circuit_limit": "19.40",
                    "52week_high": "24.75",
                    "52week_low": "12.55",
                    "open_interest": str(random.randint(1, 1000000)),
                    "multiplier": "1",
                    "precision": "2",
                    "change": str(round(random.uniform(0, 1), 2)),
                    "net_change_percentage": str(round(random.uniform(0, 1), 2)),
                    "total_traded_value": str(round(random.uniform(1000000, 100000000), 2)),
                    "instrument_token": instrument_token,
                    "exchange_segment": "nse_cm",
                    "trading_symbol": "YESBANK-EQ",
                    "ohlc": {
                        "open": str(round(random.uniform(10, 100), 2)),
                        "high": str(round(random.uniform(10, 100), 2)),
                        "low": str(round(random.uniform(10, 100), 2)),
                        "close": str(round(random.uniform(10, 100), 2))
                    }
                }]
            }
            await websocket_manager.broadcast(str(tick_data))
        await asyncio.sleep(1)

async def start_mock_data_generator():
    while True:
        tasks = []
        for token in websocket_manager.subscribed_tokens:
            tasks.append(generate_mock_tick_data(token))
        await asyncio.gather(*tasks)
        await asyncio.sleep(1)

