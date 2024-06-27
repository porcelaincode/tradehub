package services

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/porcelaincode/tradehub/internal/database"
	"github.com/porcelaincode/tradehub/internal/models"
)

const instrumentsURL = "http://api.kite.trade/instruments"

func DownloadAndSaveInstruments() error {
    resp, err := http.Get(instrumentsURL)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    reader := csv.NewReader(resp.Body)
    reader.Comma = ','

    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    for _, record := range records[1:] {
        expiry, _ := time.Parse("2006-01-02", record[5])
        instrument := models.Instrument{
            InstrumentToken: atoi(record[0]),
            ExchangeToken:   atoi(record[1]),
            TradingSymbol:   record[2],
            Name:            record[3],
            LastPrice:       atof(record[4]),
            Expiry:          expiry,
            Strike:          atof(record[6]),
            TickSize:        atof(record[7]),
            LotSize:         atoi(record[8]),
            InstrumentType:  record[9],
            Segment:         record[10],
            Exchange:        record[11],
        }

        _, err := database.DB.Exec(`
            INSERT INTO instruments (
                instrument_token, exchange_token, tradingsymbol, name, last_price, expiry, strike, tick_size, lot_size, instrument_type, segment, exchange
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
            instrument.InstrumentToken, instrument.ExchangeToken, instrument.TradingSymbol, instrument.Name, instrument.LastPrice, instrument.Expiry,
            instrument.Strike, instrument.TickSize, instrument.LotSize, instrument.InstrumentType, instrument.Segment, instrument.Exchange)
        if err != nil {
            log.Printf("Error inserting instrument: %v", err)
        }
    }
    return nil
}

func atoi(str string) int {
    i, _ := strconv.Atoi(str)
    return i
}

func atof(str string) float64 {
    f, _ := strconv.ParseFloat(str, 64)
    return f
}