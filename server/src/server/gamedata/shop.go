package gamedata 

type Shop struct {
    ProductID int
    ID int
    Name string
    Icon string
    TypeID int
    Producer string
    Repertory int
    Price float32
    Kwh float32
    Ghs float32
    Electric float32
    ManageCost float32
    Output float32
    Minimum int

}

var (
    ShopData = make(map[int]Shop)
)

func  Shopinit() {
	rf := readRf(Shop{})
	for i := 0; i < rf.NumRecord(); i++ {
		r := rf.Record(i).(*Shop)
        ShopData[r.ProductID] = *r
    }
}

func GetShopByID(id int) (Shop, bool) {
	record,exists:=  ShopData[id]
	return record,exists
}
