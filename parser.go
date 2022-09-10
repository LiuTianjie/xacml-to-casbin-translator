package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/Hasdcorona/go-xacml/pdp"
)

func getSimpleName(str string) string {
	index := strings.LastIndex(str, ":")
	if index == -1 {
		return str
	}

	return str[index+1:]
}

func ParsePolicy(path string) *pdp.Policy {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()

	var p pdp.Policy
	s, err := xmlFile.Stat()
	if err != nil {
		fmt.Println("Error stating file:", err)
		return nil
	}
	size := s.Size()
	b := make([]byte, size)
	xmlFile.Read(b)
	err = xml.Unmarshal([]byte(b), &p)
	return &p
}

func PrintPolicy(p *pdp.Policy) {
	// fmt.Printf("%+v\n", p)

	for i, rule := range p.Rules {
		sub := ""
		for i, subject := range rule.Target.Subjects.Subjects {
			sub += getSimpleName(subject.SubjectMatch.SubjectAttributeDesignator.AttributeId) + ": " + subject.SubjectMatch.AttributeValue.Value
			//sub += subject.SubjectMatch.AttributeValue.Value
			if i != len(rule.Target.Subjects.Subjects)-1 {
				sub += ", "
			}
		}

		obj := ""
		for i, object := range rule.Target.Resources.Resources {
			obj += getSimpleName(object.ResourceMatch.ResourceAttributeDesignator.AttributeId) + ": " + object.ResourceMatch.AttributeValue.Value
			//obj += object.ResourceMatch.AttributeValue.Value
			if i != len(rule.Target.Resources.Resources)-1 {
				obj += ", "
			}
		}

		act := ""
		for i, action := range rule.Target.Actions.Actions {
			act += getSimpleName(action.ActionMatch.ActionAttributeDesignator.AttributeId) + ": " + action.ActionMatch.AttributeValue.Value
			//act +=  action.ActionMatch.AttributeValue.Value
			if i != len(rule.Target.Actions.Actions)-1 {
				act += ", "
			}
		}

		con := rule.Condition.AttributeValue.Value

		eft := rule.Effect

		fmt.Print("[" + sub + ", " + obj + ", " + act + ", " + con + ", " + eft + "]")

		if i != len(p.Rules)-1 {
			fmt.Print(", ")
		}
	}
}

func ParseRequest(path string) *Request {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()

	var p Request
	s, err := xmlFile.Stat()
	if err != nil {
		fmt.Println("Error stating file:", err)
		return nil
	}
	size := s.Size()
	b := make([]byte, size)
	xmlFile.Read(b)
	err = xml.Unmarshal([]byte(b), &p)
	return &p
}

func PrintRequest(r *Request) {
	sub := "sub: "
	for i, attribute := range r.Subject.Attribute {
		subAttr, _ := json.Marshal(attribute)
		sub += string(subAttr)
		if i != len(r.Subject.Attribute)-1 {
			sub += ", "
		}
	}

	obj := "obj:"
	for i, attribute := range r.Resource.Attribute {
		objAttr, _ := json.Marshal(attribute)
		obj += string(objAttr)
		if i != len(r.Resource.Attribute)-1 {
			obj += ", "
		}
	}

	act := "act: "
	for i, attribute := range r.Action.Attribute {
		actAttr, _ := json.Marshal(attribute)
		act += string(actAttr)
		if i != len(r.Action.Attribute)-1 {
			act += ", "
		}
	}

	env := "["
	for i, attribute := range r.Environment.Attribute {
		env += attribute.AttributeValue
		if i != len(r.Environment.Attribute)-1 {
			env += ", "
		}
	}
	env += "]"

	fmt.Print(sub + ", " + obj + ", " + act + ", " + env)
}

func ParseResponse(path string) *Response {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()

	var p Response
	s, err := xmlFile.Stat()
	if err != nil {
		fmt.Println("Error stating file:", err)
		return nil
	}
	size := s.Size()
	b := make([]byte, size)
	xmlFile.Read(b)
	err = xml.Unmarshal([]byte(b), &p)
	return &p
}

func PrintResponse(r *Response) {
	fmt.Print(r.Result.Decision + ", " + strings.TrimLeft(r.Result.Status.StatusCode.Value, "urn:oasis:names:tc:xacml")[11:])
}
