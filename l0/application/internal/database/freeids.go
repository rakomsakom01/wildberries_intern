package database


type FreeIds struct {
    deliveryId int
    paymentId int
    mainId int
    itemId int
}


func (f *FreeIds) setIDs(deliveryId, paymentId, mainId, itemId int) {
    f.deliveryId = deliveryId
    f.paymentId = paymentId
    f.mainId = mainId
    f.itemId = itemId
}


func (f *FreeIds) getDeliveryId() int {
    f.deliveryId++
    return f.deliveryId
}


func (f *FreeIds) getPaymentId() int {
    f.paymentId++
    return f.paymentId
}


func (f *FreeIds) getMainId() int {
    f.mainId++
    return f.mainId
}


func (f *FreeIds) getItemId() int {
    f.itemId++
    return f.itemId
}

