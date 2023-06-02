package database

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "strings"
    "strconv"
    "log"
)


type DatabaseConsts struct {
    TableColumns   map[string]map[string]string 

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


type Database struct {
    db *sql.DB

    freeIds FreeIds

    consts DatabaseConsts

    unpushedQuery map[int]string 

    infoLog  *log.Logger
    errorLog *log.Logger
}


func (db *Database) connectToPostgres() bool {
    var postgresInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
            db.consts.PostgresHost, 
            db.consts.PostgresPort, 
            db.consts.PostgresUser, 
            db.consts.PostgresPass, 
            db.consts.PostgresDbname)

    database, err := sql.Open("postgres", postgresInfo)

    if err != nil {
        return false
    }

    db.db = database

    return true
}


func (db *Database) getFreeIds() bool {

    query := generateIdQuery(db.consts)

    rows, err := db.db.Query(query)
    if err != nil {
        return false
    }

    defer rows.Close()

    var deliveryId int
    var paymentId  int
    var mainId     int
    var itemId     int

    for rows.Next() {
        rows.Scan(&deliveryId, &paymentId, &mainId, &itemId)
    
        if err != nil {
            return false
        } 
    }

    db.freeIds.setIDs(deliveryId, paymentId, mainId, itemId)
    
    err = rows.Err()

    if err != nil {
        return false
    }

    return true
}


func (db *Database) RestoreCache() ([]int, map[int]map[string]interface{}, bool) {
    cacheList := []int{}
    cache := map[int]map[string]interface{}{}

    queryMain, queryItem, columnsMainArray := generateRestoreQuery(db.consts.TableColumns, db.consts)

    mainIdColumn := db.consts.TableMain + "_id"
    // Query get main

    rowsMain, err := db.db.Query(queryMain)

    if err != nil {
        return cacheList, cache, false
    }

    defer rowsMain.Close()

    // Parse and cache main info

    valsMain := make([]interface{}, len(columnsMainArray))
    
    for i := range valsMain {
	    valsMain[i] = new(sql.RawBytes)
    }

    for rowsMain.Next() {
	    err = rowsMain.Scan(valsMain...)

        mainId := 0
        jsonMap := map[string]interface{}{}

        jsonMap[db.consts.TableDelivery] = map[string]interface{}{}
        jsonMap[db.consts.TablePayment] = map[string]interface{}{}
        jsonMap[db.consts.TableItem] = []interface{}{}
        
        for i, v := range columnsMainArray {
            parts := strings.Split(v, ".")

            data := string(*(valsMain[i].(*sql.RawBytes)))
            
            if parts[0] == db.consts.TableMain {
                if parts[1] == mainIdColumn {
                    mainId, _ = strconv.Atoi(data)
                    continue
                }

                if db.consts.TableColumns[parts[0]][parts[1]] == "string" {
                    jsonMap[parts[1]] = data
                } else if db.consts.TableColumns[parts[0]][parts[1]] == "int" {
                    jsonMap[parts[1]], _ = strconv.Atoi(data)
                } else if db.consts.TableColumns[parts[0]][parts[1]] == "float" {
                    jsonMap[parts[1]], _ = strconv.ParseFloat(data, 64)
                }
            } else {
                jsonSubMap := jsonMap[parts[0]].(map[string]interface{})
                if db.consts.TableColumns[parts[0]][parts[1]] == "string" {
                    jsonSubMap[parts[1]] = data
                } else if db.consts.TableColumns[parts[0]][parts[1]] == "int" {
                    jsonSubMap[parts[1]], _ = strconv.Atoi(data)
                } else if db.consts.TableColumns[parts[0]][parts[1]] == "float" {
                    jsonSubMap[parts[1]], _ = strconv.ParseFloat(data, 64)
                }
            }
        }

        cache[mainId] = jsonMap
        cacheList = append(cacheList, mainId)
    }


    // Query get items info

    rowsItem, err := db.db.Query(queryItem)

    if err != nil {
        return cacheList, cache, false
    }

    defer rowsItem.Close()

    // Parse and cache items info

    columnsItemArray, _ := rowsItem.Columns()

    valsItem := make([]interface{}, len(columnsItemArray))
    
    for i := range valsItem {
	    valsItem[i] = new(sql.RawBytes)
    }

    for rowsItem.Next() {
	    err = rowsItem.Scan(valsItem...)

	    if err != nil {
            return cacheList, cache, false
        }

        mainId := 0
        jsonMap := map[string]interface{}{}
    
        for i, v := range columnsItemArray {

            parts := []string{db.consts.TableItem, v}

            data := string(*(valsItem[i].(*sql.RawBytes)))

            if parts[1] == mainIdColumn {
                mainId, _ = strconv.Atoi(data)
                continue
            }

            if db.consts.TableColumns[parts[0]][parts[1]] == "string" {
                jsonMap[parts[1]] = data
            } else if db.consts.TableColumns[parts[0]][parts[1]] == "int" {
                jsonMap[parts[1]], _ = strconv.Atoi(data)
            } else if db.consts.TableColumns[parts[0]][parts[1]] == "float" {
                jsonMap[parts[1]], _ = strconv.ParseFloat(data, 64)
            }
        }
        jsonMapArray := cache[mainId][db.consts.TableItem].([]interface{})
    
        cache[mainId][db.consts.TableItem] = append(jsonMapArray, jsonMap)
    }    


    return cacheList, cache, true
}


func (db *Database) pushUnpushed() (bool) {
    if !db.Ping() {
        return false
    }

    if len(db.unpushedQuery) != 0 {
        db.infoLog.Println("db: trying to push unpushed queries")
    }

    for id, query := range db.unpushedQuery {
 
        _, err := db.db.Exec(query)
        if err != nil {
            return false
        }
        delete(db.unpushedQuery, id)
    }
    return true
}


func (db *Database) PushData(jsonMap map[string]interface{}) (int, bool) {
    query, id, ok := generatePushQuery(db.consts, jsonMap, &db.freeIds)

    if (!ok) {
        return 0, false
    }

    db.pushUnpushed()
    _, err := db.db.Exec(query)
    if err != nil {
        db.infoLog.Println("db: failure to load data into database. Data saved in cache.")
        db.unpushedQuery[id] = query
    }
    return id, true
}



func (db *Database) Ping() bool {
    err := db.db.Ping()

    return err != nil
}


func (db *Database) Close() {
    db.pushUnpushed()

    db.Close()
}


func Init(consts DatabaseConsts, infoLog *log.Logger, errorLog *log.Logger) (Database, bool) {
    db := Database {
        consts       : consts, 
        unpushedQuery: map[int]string{},    
        infoLog      : infoLog,
        errorLog     : errorLog,
    }

    if !db.connectToPostgres() {
        return db, false
    }
    if !db.getFreeIds() {
        return db, false
    }

    return db, true
}
