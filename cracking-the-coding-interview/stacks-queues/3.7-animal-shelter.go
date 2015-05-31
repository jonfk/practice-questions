package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
)

type Queue struct {
	Value []Animal
}

func (a *Queue) Enqueue(x Animal) error {
	if a == nil {
		return fmt.Errorf("enqueuing on nil queue")
	}
	a.Value = append(a.Value, x)
	return nil
}
func (a *Queue) Dequeue() (Animal, error) {
	if a == nil || len(a.Value) < 1 {
		return Cat{"false Cat"}, fmt.Errorf("Dequeuing on nil or empty queue")
	}
	var x Animal
	x, a.Value = a.Value[0], a.Value[1:]
	return x, nil
}

type Shelter struct {
	Cat *Queue
	Dog *Queue
	Any *Queue
}

func NewShelter() *Shelter {
	return &Shelter{Cat: &Queue{}, Dog: &Queue{}, Any: &Queue{}}
}

func (shelter *Shelter) Enqueue(x Animal) error {
	switch t := x.(type) {
	case Cat:
		shelter.Cat.Enqueue(t)
	case Dog:
		shelter.Dog.Enqueue(t)
	default:
		return fmt.Errorf("Enqueue :: Unknown Animal")
	}
	shelter.Any.Enqueue(x)
	return nil
}

func (shelter *Shelter) DequeueAny() (Animal, error) {
	x, err := shelter.Any.Dequeue()
	if err != nil {
		return Cat{"False Cat"}, fmt.Errorf("DequeueAny :: %s", err.Error())
	}
	switch x.(type) {
	case Cat:
		_, err = shelter.Cat.Dequeue()
		if err != nil {
			return Cat{"False Cat"}, fmt.Errorf("DequeueAny :: %s", err.Error())
		}
	case Dog:
		_, err = shelter.Dog.Dequeue()
		if err != nil {
			return Cat{"False Cat"}, fmt.Errorf("DequeueAny :: %s", err.Error())
		}
	default:
		return Cat{"False Cat"}, fmt.Errorf("DequeueAny :: Unknown Animal")
	}
	return x, nil
}

func (shelter *Shelter) DequeueCat() (Animal, error) {
	var buffer []Animal
	cat, err := shelter.Cat.Dequeue()
	if err != nil {
		return Cat{"False Cat"}, fmt.Errorf("DequeueCat :: %s", err.Error())
	}

	var xAsCat Animal
	isCat := false
	for !isCat {
		x, err := shelter.Any.Dequeue()
		if err != nil {
			return Cat{"False Cat"}, fmt.Errorf("DequeueCat :: %s", err.Error())
		}
		xAsCat, isCat = x.(Cat)
		buffer = append(buffer, x)
	}
	if xAsCat.Name() != cat.Name() {
		return Cat{"False Cat"}, fmt.Errorf("Mismatched cats")
	}
	buffer = buffer[:len(buffer)-1]
	// put back
	for i := range buffer {
		err = shelter.Any.Enqueue(buffer[i])
		if err != nil {
			return Cat{"False Cat"}, fmt.Errorf("DequeueCat :: %s", err.Error())
		}
	}
	return cat, nil
}

func (shelter *Shelter) DequeueDog() (Animal, error) {
	var buffer []Animal
	dog, err := shelter.Dog.Dequeue()
	if err != nil {
		return Cat{"False Cat"}, fmt.Errorf("DequeueDog :: %s", err.Error())
	}

	var xAsDog Animal
	isDog := false
	for !isDog {
		x, err := shelter.Any.Dequeue()
		if err != nil {
			return Cat{"False Cat"}, fmt.Errorf("DequeueDog :: %s", err.Error())
		}
		xAsDog, isDog = x.(Dog)
		buffer = append(buffer, x)
	}
	if xAsDog.Name() != dog.Name() {
		return Cat{"False Cat"}, fmt.Errorf("Mismatched dogs")
	}
	buffer = buffer[:len(buffer)-1]
	// put back
	for i := range buffer {
		err = shelter.Any.Enqueue(buffer[i])
		if err != nil {
			return Cat{"False Cat"}, fmt.Errorf("DequeueDog :: %s", err.Error())
		}
	}
	return dog, nil
}

type Animal interface {
	Name() string
}

type Cat struct {
	name string
}

func (a Cat) Name() string {
	return a.name
}

type Dog struct {
	name string
}

func (a Dog) Name() string {
	return a.name
}

func main() {
	shelter := NewShelter()
	shelter.Enqueue(Cat{"a"})
	shelter.Enqueue(Cat{"b"})
	shelter.Enqueue(Dog{"c"})
	x, err := shelter.DequeueDog()
	if err != nil {
		log.Fatal(err)
	}
	spew.Printf("%v\n", x)
	spew.Printf("%v\n", shelter)

	// x, err = shelter.DequeueAny()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// spew.Printf("%v\n", x)
	// spew.Printf("%v\n", shelter)

}
