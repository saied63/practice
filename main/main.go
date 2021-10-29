package main

import (
	"github.com/saied63/practice/NativeRestServer"
	_ "github.com/saied63/practice/publishedmodule/complexmath"
	_ "github.com/saied63/practice/publishedmodule/restServer"
	_ "github.com/saied63/practice/publishedmodule/useless"
	_ "github.com/saied63practice/publishedmodule/simplemath"
	"go.mongodb.org/mongo-driver/bson"
	_ "strconv"
)

//var _ = complexmath.Devide
//var _ = simplemath.Add

var insertOneTest bson.D = bson.D{{Key: "name", Value: "njjn"}, {Key: "age", Value: 44}, {Key: "city", Value: "japon"}}

func init() {
	//fmt.Println("init func from main")

}

func main() {

	//fmt.Println(complexmath.Devide(17, 3))
	/* 	var a = "d"
	   	fmt.Println(a)
	   	b := "saied"
	   	fmt.Println(b)

	   	fmt.Println("print go from main func")
	   	fmt.Println(simplemath.Add(5, 2))
	   	fmt.Println(complexmath.Devide(17, 3))
	   	forStatementWithLable()
	   	anotherKindOfFor()
	   	anotherKingFor2()
	   	anotherKindOfFo3()
	   	testSwitch(10)
	   	testmultiplecase(3)
	   	testAnotherSwitchCase()
	   	fmt.Println(switchInsideFor())
	   	randomNumber() */
	//testArray()
	//testarrayInitialize()
	//testRangeOnIterator()
	//testMultidecimalIteratorWithRang()
	//testSclice()
	//testSliceLengthAndCapacity()
	//makeSliceByMake()
	//testAppend()
	//nilSlice()
	//testSliceRefrence()
	//useCoppySlice()
	// go testGoRoutine()
	//time.Sleep(10 * time.Second)
	//callWorkGroup()
	//time.Sleep(10 * time.Second)
	//AnymousFunc("")
	//AnymousFuncWithoutAssyncToVariable("")
	//AnymousFuncWithParameter()
	//UseUserDefineFunctionType()
	//UseUserDefineFunctionType_Sum()
	//UseHighOrderFunc()
	// we can run function using () and get returned values
	//i, err := ReturnHigherOrderFunction("22")()
	// or
	/*
		fn := ReturnHigherOrderFunction("22")
		i, err := fn()
		if err != nil {
			fmt.Println("error : ", err)
		} else {
			fmt.Println("converted int is : ", i)
		}
	*/
	// var isConnected bool = true
	// if !Ping() {
	// 	isConnected = false
	// 	ClientConnect()
	// 	if Ping() {
	// 		isConnected = true
	// 	}
	// }
	// var isConnected bool = true
	// if !Ping() {
	// 	isConnected = false
	// 	ClientConnect()
	// 	if Ping() {
	// 		isConnected = true
	// 	}
	// }
	// ti := time.Now()
	// if isConnected {
	// 	for i := 0; i < 100000; i++ {
	// 		InsertOne(insertOneTest)
	// 	}
	// }
	// fmt.Println(time.Since(ti).Seconds())
	//InsertOne(insertOneTest)
	//fmt.Println(time.Since(ti).Seconds())

	//ti = time.Now()
	//SelectAll()
	//fmt.Println(time.Since(ti).Seconds())
	//ti := time.Now()

	// record := bson.M{"name": "ft"}
	// fmt.Println(record)
	// if isConnected {
	// 	FindeOne(record)
	// }

	// targetPack := bson.M{
	// 	"name": "ft",
	// }

	// fmt.Println(targetPack)
	// fmt.Printf("update count %d ", UpdateOne(record, targetPack))
	// fmt.Println(time.Since(ti).Seconds())
	//TestMongo()

	NativeServer.StartServerRest()
}
