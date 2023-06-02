package main

import (
    _ "bufio"
    _ "os"
    stan "github.com/nats-io/stan.go"
)


const (
	STAN_CLUSTER_ID = "test-cluster"
	STAN_CLIENT_ID  = "send"
    STAN_CHANNEL    = "chan"

    testData = `{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`


 testData2 = `{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WEE WEE",
  "entry": "WEE WEE",
  "delivery": {
    "name": "Van Darkholme",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Gatchina",
    "region": "Leningrad",
    "email": "vanasama@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "300$",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 300,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Leather pants",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`


    testData3 = `
    {
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Pudge",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Middle",
    "address": "Mid",
    "region": "Kraiot",
    "email": "pudge@freshmeat.com"
  },
  "payment": {
    "transaction": "b563feb7b2meatb84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 3635,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 2135,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 8888,
      "track_number": "WBILMTESTTRACK",
      "price": 110,
      "rid": "qwer19087dsadad0test",
      "name": "Meat Hook",
      "sale": 0,
      "size": "0",
      "total_price": 110,
      "nm_id": 2389212,
      "brand": "Pudge Shop",
      "status": 202
    },
    {
      "chrt_id": 9999,
      "track_number": "WBILMTESTTRACK",
      "price": 1232,
      "rid": "zxc219087a764ae0btest",
      "name": "Rot",
      "sale": 100,
      "size": "0",
      "total_price": 0,
      "nm_id": 2382332,
      "brand": "Pudge Shop",
      "status": 202
    },
    {
      "chrt_id": 5555,
      "track_number": "WBILMTESTTRACK",
      "price": 80,
      "rid": "a19087a764ae0btest",
      "name": "Fresh Heap",
      "sale": 0,
      "size": "0",
      "total_price": 80,
      "nm_id": 23892143,
      "brand": "Pudge Shop",
      "status": 202
    },
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 170,
      "rid": "trfd9087a764ae0btest",
      "name": "Dismember",
      "sale": 0,
      "size": "0",
      "total_price": 170,
      "nm_id": 2389212,
      "brand": "Pudge Shop",
      "status": 202
    },
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 1775,
      "rid": "trfd9087a764ae0btest",
      "name": "Mask of Madness",
      "sale": 0,
      "size": "0",
      "total_price": 1775,
      "nm_id": 2389212,
      "brand": "Pudge Shop",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "fresh meat",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`


testData4 = `dsadhfdsjhjajijhjjpjefnejrfneirjfeijfeijfow
ijwdqiewjirfiejieajgvirehisoaj9ieruhfiepovbrtujjwcfruzehfjvpierhvuoiephviervw
dewonfijerfioaehvuehvuiejvie
fakfei9jviaenvipajirejiahu94hfuihaiuvhieuvheiahovhiuhgaihruafafaenjiabha
ejkOE;JOPHDIEJWDBFIUWFUIWHUIWHUIhiHIFJEAOHAIHAFPIORJFIAOHIPAHWIUFEWHFAUjijajfiehfuiawfiw`

testData5 = `{"coord":{"lon":10.99,"lat":44.34},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"base":"stations","main":{"temp":288.75,"feels_like":288.48,"temp_min":286.1,"temp_max":289.86,"pressure":1013,"humidity":81,"sea_level":1013,"grnd_level":929},"visibility":10000,"wind":{"speed":0.91,"deg":254,"gust":1.19},"clouds":{"all":63},"dt":1685685504,"sys":{"type":2,"id":2004688,"country":"IT","sunrise":1685676895,"sunset":1685732001},"timezone":7200,"id":3163858,"name":"Zocca","cod":200}`

testData6 = `{
   "coord":{
      "lon":10.99,
      "lat":44.34
   },
   "weather":[
      {
         "id":803,
         "main":"Clouds",
         "description":"broken clouds",
         "icon":"04d"
      }
   ],
   "base":"stations",
   "main":{
      "temp":288.75,
      "feels_like":288.48,
      "temp_min":286.1,
      "temp_max":289.86,
      "pressure":1013,
      "humidity":81,
      "sea_level":1013,
      "grnd_level":929
   },
   "visibility":10000,
   "wind":{
      "speed":0.91,
      "deg":254,
      "gust":1.19
   },
   "clouds":{
      "all":63
   },
   "dt":1685685504,
   "sys":{
      "type":2,
      "id":2004688,
      "country":"IT",
      "sunrise":1685676895,
      "sunset":1685732001
   },
   "timezone":7200,
   "id":3163858,
   "name":"Zocca",
   "cod":200
}`

  testData7 = `{
  "order_uid": "b563feb7b2b84b6test",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`
)





func main() {


    sendClient := STAN_CLIENT_ID
    sc, err := stan.Connect(STAN_CLUSTER_ID, sendClient, stan.NatsURL(stan.DefaultNatsURL))

    if err != nil {
        panic(err)
    }

    defer sc.Close()

    /*
     Первые 3 - корректные данные
     Далее идёт мусор, 
     после 2 jsonа с openweathermap
     Ласт - корректный json, но с отсутствующими полями
     */

    err = sc.Publish(STAN_CHANNEL, []byte(testData)) // does not return until an ack has been received from NATS Streaming
    err = sc.Publish(STAN_CHANNEL, []byte(testData2)) // does not return until an ack has been received from NATS Streaming
    err = sc.Publish(STAN_CHANNEL, []byte(testData3)) // does not return until an ack has been received from NATS Streaming
    err = sc.Publish(STAN_CHANNEL, []byte(testData4)) // does not return until an ack has been received from NATS Streaming
    err = sc.Publish(STAN_CHANNEL, []byte(testData5)) // does not return until an ack has been received from NATS Streaming
    err = sc.Publish(STAN_CHANNEL, []byte(testData6)) // does not return until an ack has been received from NATS Streaming
    err = sc.Publish(STAN_CHANNEL, []byte(testData7)) // does not return until an ack has been received from NATS Streaming

    if err != nil {
        panic(err)
    }
}
