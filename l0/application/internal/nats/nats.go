package nats

import (
    stan "github.com/nats-io/stan.go"
    "encoding/json"
    "strconv"
    "time"
    "log"
)


type NatsConsts struct {
    // Consts
    NatsHost        string
    NatsPort        int
    NatsClusterId   string
    NatsClientId    string
    NatsChannel     string
    NatsDurableName string

    TableColumns    map[string]map[string]string
}


type Nats struct {
    consts NatsConsts 

    sc stan.Conn

    sub stan.Subscription

    jsonData chan<- map[string]interface{}

    infoLog  *log.Logger
    errorLog *log.Logger
}



func verifyJsonMap(columns map[string]map[string]string, jsonMap map[string]interface{}, tableName string) (bool) {
    table, ok := columns[tableName]

    if !ok {
        return false
    }

    for key, value := range jsonMap {
        trueValueType, ok := table[key]
        
        // JSON Contains extra data
        if !ok {
            return false
        }


        if trueValueType == "string" {
            _, ok := value.(string)

            if !ok {
                return false
            }
        } else if trueValueType == "int" {
            trueValue, ok := value.(float64)

            if !ok {
                return false
            }

            delete(jsonMap, key)
            jsonMap[key] = int(trueValue)
        } else if trueValueType == "float" {
            _, ok := value.(float64)

            if !ok {
                return false
            }
        } else if trueValueType == "map" {
            trueValue, ok := value.(map[string]interface{})

            if !ok {
                return false
            }

            if !verifyJsonMap(columns, trueValue, key) {
                return false
            }
        } else if trueValueType == "array" {
            trueValue, ok := value.([]interface{})

            if !ok {
                return false
            }
            
            for i := range trueValue {
                arrValue, ok := trueValue[i].(map[string]interface{})

                if !ok {
                    return false
                }
                
                if !verifyJsonMap(columns, arrValue, key) {
                    return false
                }
            }
        }
    }

    return true
}


func (nats *Nats) resetConnection() {
    for {
        var ok bool
    
        time.Sleep(10 * time.Second)

        ok = nats.connectToNATS()
        if (!ok) {
            continue
        }
    
        ok = nats.subscribe()
        if (!ok) {
            continue
        }
        break
    }
}


func (nats *Nats) connectToNATS() (bool) {
    sc, err := stan.Connect(
            nats.consts.NatsClusterId, 
            nats.consts.NatsClientId, 
            stan.NatsURL(
                "nats://" + nats.consts.NatsHost + ":" + strconv.Itoa(nats.consts.NatsPort)),
            stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
                nats.errorLog.Println("nats: connection failed. Trying to reconect.")
                go nats.resetConnection()
            }))

    if err != nil {
        return false
    }

    nats.sc = sc

    return true
}


func (nats *Nats) parseMessage(m *stan.Msg) {
    var jsonMap map[string]interface{}
    err := json.Unmarshal(m.Data, &jsonMap)
    if err != nil {
        nats.infoLog.Println("nats: incorrectData.")

        return
    }

    if !verifyJsonMap(nats.consts.TableColumns, jsonMap, "model") {
        nats.infoLog.Println("nats: incorrectData.")
        
        return
    }

    nats.jsonData <- jsonMap
}




func (nats *Nats) subscribe() (bool) {
    sub, err := nats.sc.Subscribe(nats.consts.NatsChannel, nats.parseMessage, 
            stan.DurableName(nats.consts.NatsDurableName))

    if err != nil {
        return false
    }
    
    nats.sub = sub
    return true
}


func (nats *Nats) Close() {
    nats.sub.Unsubscribe()

    nats.sub.Close()
}


func Init(consts NatsConsts, jsonData chan<- map[string]interface{},
	    infoLog *log.Logger, errorLog  *log.Logger) (Nats, bool) {

    nats := Nats {
        consts  : consts,
        jsonData: jsonData,
        infoLog : infoLog,
        errorLog: errorLog,
    }

    var ok bool
    
    ok = nats.connectToNATS()
    if (!ok) {
        return nats, false
    }
    
    ok = nats.subscribe()
    if (!ok) {
        return nats, false
    }

    return nats, true
}
