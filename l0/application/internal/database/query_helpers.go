package database

import (
    "fmt"
)


func generateInsertColumns(columns map[string]map[string]string, table string) ([]string) {
    var columnsArray []string

    for key, typeValue := range columns[table] {
        if typeValue == "string" || typeValue == "int" || typeValue == "float" {
            columnsArray = append(columnsArray, key)
        }
    }

    return columnsArray
}

func generateInsertData(columns map[string]map[string]string, jsonMap map[string]interface{}, 
        table string, cols []string) ([]string) {
    
    var valuesArray []string

    for _, key := range cols {
        value, ok := jsonMap[key]

        typeValue := columns[table][key]

        if typeValue == "string" {
            if !ok {
                value = ""
            }
            
            valuesArray = append(valuesArray, fmt.Sprintf("'%s'", value))         
        } else if typeValue == "int" {
            if !ok {
                value = 0
            }

            valuesArray = append(valuesArray, fmt.Sprintf("%d", value))         
        } else if typeValue == "float" {
            if !ok {
                value = 0.0
            }

            valuesArray = append(valuesArray, fmt.Sprintf("%f", value))         
        }
    }

   
    return valuesArray
}


func addPrefix(columnsArray []string, prefix string) {
    for i := range(columnsArray) {
        columnsArray[i] = prefix + columnsArray[i]
    }
}

