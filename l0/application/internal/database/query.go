package database

import (
    "fmt"
    "strings"
)


func generatePushQuery(consts DatabaseConsts, 
        jsonMap map[string]interface{}, freeIds *FreeIds) (string, int, bool) { 
    
    columns := consts.TableColumns

    // First, we insert delivery
    
    jsonDelivery, ok := jsonMap[consts.TableDelivery].(map[string]interface{})

    if (!ok) {
        return "", 0, false
    }

    columnsDeliveryArray := generateInsertColumns(columns, consts.TableDelivery)
    valuesDeliveryArray  := generateInsertData(columns, jsonDelivery, consts.TableDelivery, columnsDeliveryArray)

    deliveryId := freeIds.getDeliveryId()
    columnsDeliveryArray = append(columnsDeliveryArray, consts.TableDelivery + "_id")
    valuesDeliveryArray  = append(valuesDeliveryArray , fmt.Sprintf("%d", deliveryId))

    queryDelivery := fmt.Sprintf("INSERT INTO %s(\n    %s\n)\nVALUES (\n    %s\n);\n\n", 
            consts.TableDelivery, 
            strings.Join(columnsDeliveryArray, ",\n    "), 
            strings.Join(valuesDeliveryArray , ",\n    "))

    // Payment

    jsonPayment, ok := jsonMap[consts.TablePayment].(map[string]interface{})

    if (!ok) {
        return "", 0, false
    }

    columnsPaymentArray := generateInsertColumns(columns, consts.TablePayment)
    valuesPaymentArray  := generateInsertData(columns, jsonPayment, consts.TablePayment, columnsPaymentArray)

    paymentId := freeIds.getPaymentId()
    columnsPaymentArray = append(columnsPaymentArray, consts.TablePayment + "_id")
    valuesPaymentArray  = append(valuesPaymentArray , fmt.Sprintf("%d", paymentId))

    queryPayment := fmt.Sprintf("INSERT INTO %s(\n    %s\n)\nVALUES (\n    %s\n);\n\n", 
            consts.TablePayment, 
            strings.Join(columnsPaymentArray, ",\n    "), 
            strings.Join(valuesPaymentArray , ",\n    "))

    // Add main

    jsonMain := jsonMap

    columnsMainArray := generateInsertColumns(columns, consts.TableMain)
    valuesMainArray  := generateInsertData(columns, jsonMain, consts.TableMain, columnsMainArray)

    mainId := freeIds.getMainId()
    columnsMainArray = append(columnsMainArray, consts.TableMain + "_id")
    valuesMainArray  = append(valuesMainArray , fmt.Sprintf("%d", mainId))

    columnsMainArray = append(columnsMainArray, consts.TableDelivery + "_id")
    valuesMainArray  = append(valuesMainArray , fmt.Sprintf("%d", deliveryId))

    columnsMainArray = append(columnsMainArray, consts.TablePayment + "_id")
    valuesMainArray  = append(valuesMainArray , fmt.Sprintf("%d", paymentId))



    queryMain := fmt.Sprintf("INSERT INTO %s(\n    %s\n)\nVALUES (\n    %s\n);\n\n", 
            consts.TableMain, 
            strings.Join(columnsMainArray, ",\n    "), 
            strings.Join(valuesMainArray , ",\n    "))

    // items 

 
    jsonItemArray := jsonMap[consts.TableItem].([]interface{})

    queryItem := ""

    if len(jsonItemArray) != 0 {
        
        columnsItemArray := generateInsertColumns(columns, consts.TableItem) 

        var queryItemArray []string 

        for _, jsonRawItem := range jsonItemArray {
            jsonItem := jsonRawItem.(map[string]interface{})
            valuesItemArray  := generateInsertData(columns, jsonItem, consts.TableItem, columnsItemArray)
    
            itemId := freeIds.getItemId()
            valuesItemArray  = append(valuesItemArray , fmt.Sprintf("%d", itemId))
            
            valuesItemArray  = append(valuesItemArray , fmt.Sprintf("%d", mainId))

            queryItemArray = append(queryItemArray, fmt.Sprintf("(\n    %s\n)", 
                    strings.Join(valuesItemArray , ",\n    ")))

        }
        columnsItemArray = append(columnsItemArray, consts.TableItem + "_id")

        columnsItemArray = append(columnsItemArray, consts.TableMain + "_id")

        queryItem = fmt.Sprintf("INSERT INTO %s(\n    %s\n)\nVALUES %s;", 
                consts.TableItem, 
                strings.Join(columnsItemArray, ",\n    "),
                strings.Join(queryItemArray, ", "))
    }


    return queryDelivery + queryPayment + queryMain + queryItem, mainId, true
}


func generateRestoreQuery(columns map[string]map[string]string, consts DatabaseConsts) (
        string, string, []string) {
   
    columnsArray := []string{consts.TableMain + "." + consts.TableMain + "_id"}

    // Main

    columnsMainArray := generateInsertColumns(columns, consts.TableMain)

    addPrefix(columnsMainArray, consts.TableMain + ".")

    columnsArray = append(columnsArray, columnsMainArray...)
    
    // Delivery

    columnsDeliveryArray := generateInsertColumns(columns, consts.TableDelivery)
    
    addPrefix(columnsDeliveryArray, consts.TableDelivery + ".")

    columnsArray = append(columnsArray, columnsDeliveryArray...)

    // Payment

    columnsPaymentArray := generateInsertColumns(columns, consts.TablePayment)

    addPrefix(columnsPaymentArray, consts.TablePayment + ".")

    columnsArray = append(columnsArray, columnsPaymentArray...)


    query := strings.Join(columnsArray, ",\n    ")

    // Item

    columnsItemArray := []string{consts.TableMain + "_id"}
    origColumnsItemArray := generateInsertColumns(columns, consts.TableItem)
    columnsItemArray = append(columnsItemArray, origColumnsItemArray...)

    queryItem := fmt.Sprintf("SELECT\n    %s\nFROM %s",
            strings.Join(columnsItemArray, ",\n    "), consts.TableItem)

    // Final query

    query1 := fmt.Sprintf("SELECT\n    %s\nFROM %s\n", 
            query, consts.TableMain)
    query2 := fmt.Sprintf("LEFT JOIN %s ON %s.%s_id=%s.%s_id\n", consts.TableDelivery, 
            consts.TableMain, consts.TableDelivery, consts.TableDelivery, consts.TableDelivery)
    query3 := fmt.Sprintf("LEFT JOIN %s ON %s.%s_id=%s.%s_id;\n", consts.TablePayment, 
            consts.TableMain, consts.TablePayment, consts.TablePayment, consts.TablePayment)

    return query1 + query2 + query3, queryItem, columnsArray
            
}


func generateIdQuery(consts DatabaseConsts) string {
     query := fmt.Sprintf(`SELECT 
            MAX(%s.%s_id), MAX(%s.%s_id), MAX(%s.%s_id), MAX(%s.%s_id) 
            FROM %s, %s, %s, %s;`, 
            consts.TableDelivery, consts.TableDelivery, 
            consts.TablePayment , consts.TablePayment , 
            consts.TableMain    , consts.TableMain    , 
            consts.TableItem    , consts.TableItem    ,
            consts.TableDelivery, 
            consts.TablePayment, 
            consts.TableMain, 
            consts.TableItem)   

    return query
}
