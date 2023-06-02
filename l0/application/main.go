package main

import (
    "application/cmd/app"
)

const (
    NATS_HOST         = "localhost"
    NATS_PORT         = 4222
    NATS_CLUSTER_ID   = "test-cluster"
    NATS_CLIENT_ID    = "main-client" 
    NATS_CHANNEL      = "chan"        
    NATS_DURABLE_NAME = "durable"      

    POSTGRES_HOST     = "localhost"
    POSTGRES_PORT     = 5432       
    POSTGRES_USER     = "paul"     
    POSTGRES_PASS     = "example"  
    POSTGRES_DBNAME   = "example" 
    
    POSTGRES_TABLE_DELIVERY = "delivery"
    POSTGRES_TABLE_PAYMENT  = "payment"
    POSTGRES_TABLE_ITEMS    = "items"
    POSTGRES_TABLE_MAIN     = "model"
)


func main() {
    var appConsts = app.ApplicationConsts{
        NatsHost       : NATS_HOST,
        NatsPort       : NATS_PORT,
        NatsClusterId  : NATS_CLUSTER_ID,
        NatsClientId   : NATS_CLIENT_ID,
        NatsChannel    : NATS_CHANNEL,
        NatsDurableName: NATS_DURABLE_NAME,
        PostgresHost   : POSTGRES_HOST,
        PostgresPort   : POSTGRES_PORT,
        PostgresUser   : POSTGRES_USER,
        PostgresPass   : POSTGRES_PASS,
        PostgresDbname : POSTGRES_DBNAME,
        TableDelivery  : POSTGRES_TABLE_DELIVERY,
        TablePayment   : POSTGRES_TABLE_PAYMENT,
        TableItem      : POSTGRES_TABLE_ITEMS,
        TableMain      : POSTGRES_TABLE_MAIN,
    }

    var application app.Application

    application.Run(appConsts)
}

