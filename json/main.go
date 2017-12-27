package main 
import (
	"encoding/json"
	"fmt"
	"os"
)
type Person struct {
	Name  Name
	Email []Email
}
type Name struct {
	Family string
	Personal string
}
type Email struct {
	Kind string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
	s += "\n" + v.Kind + ": " + v.Address
	}
	return s
	}
func saveJson(fileName string,key interface{}){
	outFile,err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}
func loadJson(fileName string,key interface{}) {
	inFile,err := os.Open(fileName)
	checkError(err)
	
	decoder := json.NewDecoder(inFile)

	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()

}
func checkError(err error) {
	if err != nil {
	fmt.Println("Fatal error ", err.Error())
	os.Exit(1)
	}
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
		Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}
	saveJson("person1.json", person)

	var person1 Person
	loadJson("person.json", &person1)

	fmt.Println("Person", person1.String())
	fmt.Println("name",person1.Name.Family);
}