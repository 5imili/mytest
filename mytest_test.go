package mytest

import (
	"testing"
	//"fmt"
	"math/rand"
	"time"
	"fmt"
)

//func TestSum(t *testing.T){
//	t.Parallel()
//	if Sum(1,2) != 3{
//		t.FailNow()
//	}
//	if Sum(2,3) == 5{
//		t.Log("Success")
//	}
//
//	var tests = []struct{
//		x,y , expect int
//	}{
//		{1,2,3},
//		{1,3,4},
//		{0,0,0},
//	}
//
//	for _, u := range tests{
//		result := Sum(u.x,u.y)
//		if result != u.expect{
//			t.Errorf("%d + %d expect %d,actual result is :%d\n", u.x, u.y, u.expect,result)
//		}
//	}
//}
//
func TestDiv(t *testing.T){
	t.Parallel()
	if ret, err := Div(1,2); err == nil{
		if ret == 0{
			t.Log("Success")
		}
	}
		//
		//if _, err := Div(1,0);err != nil{
		//	t.Log(err.Error())
		//}
}



func TestBubbleSort(t *testing.T){
	t.Parallel()
	//items :=[]int{55,33,11,99,77}
	for i:= 0;i < 100 ;i++{
		items := make([]int,0)
		for j:=0;j < 50;j++{
			items = append(items, rand.Intn(100))
		}
		//t.Log(t.Name(),"before: ",items)
		BubbleSort(items)
		//t.Log(t.Name(),"after: ", items)
	}
}



func TestQuickSort(t *testing.T){
	t.Parallel()
	//items :=[]int{55,33,11,99,77}
	//t.Log("before: ",items)
	//QuickSort(items)
	//t.Log("after: ", items)

	for i:= 0;i < 1 ;i++{
		//time.Sleep(time.Second *1)
		items := make([]int,0)
		for j:=0;j < 50;j++{
			items = append(items, rand.Intn(100))
		}
		//t.Log(t.Name() , "before: ",items)
		QuickSort(items)
		//t.Log(t.Name(),"after: ", items)
	}
}

func TestMain(m *testing.M){
	println("set up")
	rand.Seed(time.Now().Unix())
	m.Run()
	println("tear down")
}

func BenchmarkBubbleSort(b *testing.B) {
	for i:= 0;i < 1 ;i++{
		//time.Sleep(time.Second *1)
		items := make([]int,0)
		for j:=0;j < 5000;j++{
			items = append(items, rand.Intn(10000))
		}
		//t.Log(t.Name() , "before: ",items)
		BubbleSort(items)
		//t.Log(t.Name(),"after: ", items)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	//items :=[]int{55,33,11,99,77}
		for i:= 0;i < b.N ;i++{
			b.StopTimer()
			fmt.Println(i)
			b.StartTimer()
			//time.Sleep(time.Second *1)
			items := make([]int,0)
			for j:=0;j < 5000;j++{
				items = append(items, rand.Intn(10000))
			}
			//t.Log(t.Name() , "before: ",items)
			QuickSort(items)
			//t.Log(t.Name(),"after: ", items)
		}
	//QuickSort(items)
}