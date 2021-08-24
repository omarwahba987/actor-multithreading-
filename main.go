package main

import (
	"fmt"
	"time"
)

var ActorMap = make(map[string]Actor)

type Actor interface {
	Increase()
	Decrease()
	GetName() string
	GetCount() int
	DisplayActorsInformation()
}
type DefaultActor struct {
	Name  string
	Count int
}

func (d DefaultActor) GetCount() int {
	return d.Count
}

func (d DefaultActor) DisplayActorsInformation() {
	time.Sleep(time.Millisecond * 10)
	for _, v := range ActorMap {
		fmt.Printf("actor name: %v actor count : %v \n", v.GetName(),v.GetCount())
	}
}

func (d DefaultActor) GetName() string {
	return d.Name
}

func (d *DefaultActor) Increase() {
	d.Count++
}

func (d *DefaultActor) Decrease() {
	d.Count--
}

type WithValueActor struct {
	Name  string
	Count int
	value int
}



type AutoIncreaseActor struct {
	AutoIncrease WithValueActor
}

func (a AutoIncreaseActor) GetCount() int {
	return a.AutoIncrease.Count
}

func (a AutoIncreaseActor) DisplayActorsInformation() {
	time.Sleep(time.Millisecond * 10)
	for _, v := range ActorMap {
		fmt.Printf("actor name: %v actor count : %v \n", v.GetName(),v.GetCount())
	}
}

func (a AutoIncreaseActor) GetName() string {
	return a.AutoIncrease.Name
}

func (a *AutoIncreaseActor) Increase() {
	a.AutoIncrease.Count += a.AutoIncrease.value
}

func (a *AutoIncreaseActor) Decrease() {
	a.AutoIncrease.Count -= a.AutoIncrease.value
}

type AutoDecreaseActor struct {
	AutoDecrease WithValueActor
}

func (a AutoDecreaseActor) GetCount() int {
	return a.AutoDecrease.Count
}

func (a AutoDecreaseActor) DisplayActorsInformation() {
	time.Sleep(time.Millisecond * 10)
	for _, v := range ActorMap {
		fmt.Printf("actor name: %v actor count : %v \n", v.GetName(),v.GetCount())
	}
}

func (a AutoDecreaseActor) GetName() string {
	return a.AutoDecrease.Name
}

func (a *AutoDecreaseActor) Increase() {
	a.AutoDecrease.Count += a.AutoDecrease.value
}

func (a *AutoDecreaseActor) Decrease() {
	a.AutoDecrease.Count -= a.AutoDecrease.value
}


type Dispatcher struct{}

func (d Dispatcher) DispatchActor(ch chan Actor) {
	for {
		actor := <-ch
		if _, ok := ActorMap[actor.GetName()]; ok {
			fmt.Printf("***actor name - %v - already exist ***\n", actor.GetName())
		} else {
			ActorMap[actor.GetName()] = actor
			fmt.Printf("***actor - %v - added ***\n", actor.GetName())
		}
	}
}

func main() {
	ch := make(chan Actor, 1)
	actorDispatcher := Dispatcher{}
	go actorDispatcher.DispatchActor(ch)


	input:
		time.Sleep(time.Millisecond*20)

		fmt.Println("for dispatch a new actor press 1 \n" +
			"for select an actor press 2\n"+
			"to exit press CTRL+c")
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
		return
	}
	if input==1{

		fmt.Println("for dispatch a new basic actor press 1 \n" +
			"for dispatch a new auto increase actor press 2 \n"+
			"for dispatch a new auto decrease actor press 3 ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
			return
		}
		if input==1 {
			fmt.Println("enter the actor name")
			var  name string
			_, err = fmt.Scanln(&name)
			if err != nil {
				fmt.Println(err)
				return
			}
			ch<-&DefaultActor{name,0}
			goto input
		}else if input==2 {
			fmt.Println("enter the actor name")
			var  name string
			var interval int
			_, err = fmt.Scanln(&name)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("enter the increase interval value")
			_, err = fmt.Scanln(&interval)
			if err != nil {
				fmt.Println(err)
				return
			}
			ch<-&AutoIncreaseActor{WithValueActor{name,0,interval}}
			goto input
		}else if input==3 {
			fmt.Println("enter the actor name")
			var  name string
			var interval int
			_, err = fmt.Scanln(&name)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("enter the decrease interval value")
			_, err = fmt.Scanln(&interval)
			if err != nil {
				fmt.Println(err)
				return
			}
			ch<-&AutoDecreaseActor{WithValueActor{name,0,interval}}
		}else {
			fmt.Println("please enter a valid choice")
			goto input
		}

	}else if input==2{
		fmt.Println("enter the actor name")
		var  name string
		_, err = fmt.Scanln(&name)
		if err != nil {
			fmt.Println(err)
			return
		}
		if actor,ok:= ActorMap[name];ok{
			input2:
			fmt.Println("you now selecting actor :",name)
			fmt.Println("if you want to increase press 1\n"+
				"if you want to decrease press 2\n"+
				"if you want to show all actors count press 3")
			var input int
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println(err)
				return
			}
			if input==1{
				actor.Increase()
				goto input2

			}else if input==2 {
				actor.Decrease()
				goto input2
			}else if input==3 {
				actor.DisplayActorsInformation()
				goto input
			}else {
				fmt.Println("please enter a valid choice")
				goto input
			}
		}else {
			fmt.Println("this actor not exist")
			goto input
		}
	}else {
		fmt.Println("please enter a valid choice")
		goto input
	}




	//x := &DefaultActor{"omar", 0}
	//ch <- x
	//ch <- x
	//y := &AutoDecreaseActor{WithValueActor{"ahmed", 0, 2}}
	//ch <- y
	//time.Sleep(time.Millisecond * 10)
	//ActorMap["ahmed"].Increase()
	//fmt.Println(ActorMap["ahmed"])
	//y.Increase()
	//fmt.Println(ActorMap["ahmed"])

}
