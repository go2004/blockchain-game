package gamedata 

type MineLocation struct {
    Location int
    Name string
    Icon string
    Electric float32
    ManageCost float32
    Output float32
    Coord string
    Num int

}

var (
    MineLocationData = make(map[int]MineLocation)
)

func  MineLocationinit() {
	rf := readRf(MineLocation{})
	for i := 0; i < rf.NumRecord(); i++ {
		r := rf.Record(i).(*MineLocation)
        MineLocationData[r.Location] = *r
    }
}

func GetMineLocationByID(id int) (MineLocation, bool) {
	record,exists:=  MineLocationData[id]
	return record,exists
}
