package app

import (
    _"fmt"
    "os"
    "log"
    "application/internal/web"
    "application/internal/database"
    "application/internal/nats"
)

var tableColumns = map[string]map[string]string {
    "delivery": map[string]string { 
        "name"        : "string",
        "phone"       : "string",
        "zip"         : "string",
        "city"        : "string",
        "address"     : "string",
        "region"      : "string",
        "email"       : "string",
    },
    "payment": map[string]string {
        "transaction"  : "string", 
        "request_id"   : "string",
        "currency"     : "string",
        "provider"     : "string",
        "amount"       : "int"   ,
        "payment_dt"   : "int"   ,
        "bank"         : "string",
        "delivery_cost": "float" ,
        "goods_total"  : "float" ,
        "custom_fee"   : "float" ,
    },
    "items": map[string]string {
        "chrt_id"     : "int"   ,  
        "track_number": "string",
        "price"       : "float" ,
        "rid"         : "string",
        "name"        : "string",
        "sale"        : "float" ,
        "size"        : "string",
        "total_price" : "float" ,
        "nm_id"       : "int"   ,
        "brand"       : "string",
        "status"      : "int"   ,  
    },
    "model": map[string]string {
        "order_uid"         : "string",
        "track_number"      : "string",
        "entry"             : "string",
        "delivery"          : "map"   ,
        "payment"           : "map"   ,
        "items"             : "array" ,
        "locale"            : "string",
        "internal_signature": "string",
        "customer_id"       : "string",
        "delivery_service"  : "string",
        "shardkey"          : "string",
        "sm_id"             : "int"   ,   
        "date_created"      : "string",
        "oof_shard"         : "string",
    },
}


type ApplicationConsts struct {
    // Consts
    NatsHost        string
    NatsPort        int
    NatsClusterId   string
    NatsClientId    string
    NatsChannel     string
    NatsDurableName string

    PostgresHost   string
    PostgresPort   int
    PostgresUser   string
    PostgresPass   string
    PostgresDbname string

    TableDelivery string
    TablePayment  string
    TableItem     string
    TableMain     string
}


type Application struct {
    consts ApplicationConsts 
    
    db        database.Database
    nats      nats.Nats
    webServer web.WebServer
    

    cacheList []int 
    cache map[int]map[string]interface{}


    jsonData chan map[string]interface{}

    cacheListRequest chan int 
    cacheListResult chan []int

    cacheValRequest chan int 
    cacheValResult chan map[string]interface{}

	infoLog  *log.Logger
    errorLog *log.Logger
}



func (app * Application) initChannels() {
    app.jsonData = make(chan map[string]interface{})
    
    app.cacheListRequest = make(chan int)
    app.cacheListResult  = make(chan []int)

    app.cacheValRequest  = make(chan int) 
    app.cacheValResult   = make(chan map[string]interface{})
}


func (app *Application) connectToNATS() {
    consts := nats.NatsConsts{
        TableColumns    : tableColumns,
        NatsHost        : app.consts.NatsHost,
        NatsPort        : app.consts.NatsPort,
        NatsClusterId   : app.consts.NatsClusterId,
        NatsClientId    : app.consts.NatsClientId,
        NatsChannel     : app.consts.NatsChannel,
        NatsDurableName : app.consts.NatsDurableName,
    }

    var ok bool
    app.nats, ok = nats.Init(consts, app.jsonData, app.infoLog, app.errorLog)
    if (!ok) {
        app.errorLog.Fatalf("nats: connection failed!")
    }
}


func (app *Application) connectToPostgres() {
    consts := database.DatabaseConsts{
        TableColumns  : tableColumns,
        PostgresHost  : app.consts.PostgresHost,
        PostgresPort  : app.consts.PostgresPort,
        PostgresUser  : app.consts.PostgresUser,
        PostgresPass  : app.consts.PostgresPass,
        PostgresDbname: app.consts.PostgresDbname,
        TableDelivery : app.consts.TableDelivery,
        TablePayment  : app.consts.TablePayment, 
        TableItem     : app.consts.TableItem,
        TableMain     : app.consts.TableMain,
    }

    var ok bool
    app.db, ok = database.Init(consts, app.infoLog, app.errorLog)
    if (!ok) {
        app.errorLog.Fatalf("db: connection failed!")
    }
    app.cacheList, app.cache, ok = app.db.RestoreCache()
    if (!ok) {
        app.errorLog.Fatalf("db: connection failed!")
    }
}


func (app * Application) listenJsonData() {
    for {
        jsonMap := <-app.jsonData

        id, _ := app.db.PushData(jsonMap)

        app.cacheList = append(app.cacheList, id)

        app.cache[id] = jsonMap
    }
}


func (app * Application) listenCacheList() {
    for {
        _ = <-app.cacheListRequest
        app.cacheListResult <- app.cacheList
    }
}


func (app * Application) listenCacheVal() {
    for {
        id := <-app.cacheValRequest
        
        jsonMap, ok := app.cache[id]

        if !ok {
            res := map[string]interface{}{
                "miss": "miss",
            }

            app.cacheValResult<- res
            continue
        }

        app.cacheValResult<- jsonMap
    }
}


func (app *Application) runServices() {
    go app.listenJsonData()
    go app.listenCacheList()
    go app.listenCacheVal()

}


func (app *Application) Run(consts ApplicationConsts) {

	app.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
 
	app.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
 

    app.consts = consts

    app.initChannels()
    
    app.connectToNATS()
    defer app.nats.Close()

    app.connectToPostgres()
    defer app.db.Close()

    app.runServices()

    //app.interactive()
    app.webServer = web.InitWeb(app.cacheListRequest, app.cacheListResult, 
            app.cacheValRequest, app.cacheValResult, app.infoLog, app.errorLog)

    app.webServer.Start()
}
