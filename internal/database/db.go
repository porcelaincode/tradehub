package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("sqlite3", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }

    createTables()
}

func createTables() {
    createInstrumentsTable := `
    CREATE TABLE IF NOT EXISTS instruments (
        instrument_token INTEGER PRIMARY KEY,
        exchange_token INTEGER,
        tradingsymbol TEXT,
        name TEXT,
        last_price REAL,
        expiry DATE,
        strike REAL,
        tick_size REAL,
        lot_size INTEGER,
        instrument_type TEXT,
        segment TEXT,
        exchange TEXT
    );`

    createOrdersTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        am TEXT,
        disclosed_quantity TEXT,
        exchange_segment TEXT,
        market_protection TEXT,
        product TEXT,
        pf TEXT,
        price TEXT,
        order_type TEXT,
        quantity TEXT,
        validity TEXT,
        trigger_price TEXT,
        trading_symbol TEXT,
        transaction_type TEXT,
        tag TEXT
    );`

    createPositionsTable := `
    CREATE TABLE IF NOT EXISTS positions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        buyAmt TEXT,
        cfSellAmt TEXT,
        prod TEXT,
        exSeg TEXT,
        sqrFlg TEXT,
        actId TEXT,
        cfBuyQty TEXT,
        cfSellQty TEXT,
        tok TEXT,
        flBuyQty TEXT,
        flSellQty TEXT,
        sellAmt TEXT,
        posFlg TEXT,
        cfBuyAmt TEXT,
        stkPrc TEXT,
        trdSym TEXT,
        sym TEXT,
        expDt TEXT,
        type TEXT,
        series TEXT,
        brdLtQty TEXT,
        exp TEXT,
        optTp TEXT,
        genNum TEXT,
        genDen TEXT,
        prcNum TEXT,
        prcDen TEXT,
        lotSz TEXT,
        multiplier TEXT,
        precision TEXT,
        hsUpTm TEXT
    );`

    _, err := DB.Exec(createInstrumentsTable)
    if err != nil {
        log.Fatalf("Error creating instruments table: %v", err)
    }

    _, err = DB.Exec(createOrdersTable)
    if err != nil {
        log.Fatalf("Error creating orders table: %v", err)
    }

    _, err = DB.Exec(createPositionsTable)
    if err != nil {
        log.Fatalf("Error creating positions table: %v", err)
    }
}