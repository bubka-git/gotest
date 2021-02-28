package main

import (
	"fmt"
	"io/ioutil"
	"log"
	complexpb "protobuf/complex"
	enumpb "protobuf/enum_example"
	simplepb "protobuf/simple"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	//sm := doSimple()

	//readAndWrite("simple.bin", sm)
	//readAndWriteJSON(sm)

	doEnum()

	doComplex()

}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third Message",
			},
			&complexpb.DummyMessage{
				Id:   4,
				Name: "Fourth Message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           32,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_FRIDAY

	fmt.Println(em)
}

func readAndWriteJSON(pmsg proto.Message) {
	smAsString := protoToJSON(pmsg)

	fmt.Println("Marshaled to string:\n", smAsString)

	sm2 := &simplepb.SimpleMessage{}
	protoFromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct: ", sm2)
}

func protoToJSON(pmsg proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pmsg)
	if err != nil {
		log.Fatalln("Unable to marshal to string")
		return ""
	}
	return out
}

func protoFromJSON(inStr string, pmsg proto.Message) {
	if err := jsonpb.UnmarshalString(inStr, pmsg); err != nil {
		log.Fatalln("Cannot unmarshal the string to pb Proto: ", err)
	}
}

func readAndWrite(fileName string, pmsg proto.Message) {
	//writeToFile()
	writeToFile(fileName, pmsg)

	sm2 := &simplepb.SimpleMessage{}

	//readFromFile()
	readFromFile(fileName, sm2)
	fmt.Println("Reading the content: ", sm2)
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 9},
	}
	fmt.Println(sm)
	fmt.Println("The ID is: ", sm.GetId())
	return &sm
}

func writeToFile(fileName string, pmsg proto.Message) error {

	out, err := proto.Marshal(pmsg)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalln("Can't write to file ", err)
		return err
	}
	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fileName string, pmsg proto.Message) error {

	out, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Unable to read file", err)
		return err
	}

	if err := proto.Unmarshal(out, pmsg); err != nil {
		log.Fatalln("Couldn't put bytes into proto struct ", err)
		return err
	}
	return nil
}
