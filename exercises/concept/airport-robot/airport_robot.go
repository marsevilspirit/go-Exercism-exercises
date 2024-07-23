package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Italian struct{}

type Portuguese struct{}


func SayHello(name string, language interface{}) string {
    switch language.(type) {
    case Italian:
        return "I can speak Italian: Ciao " + name + "!"
    case Portuguese:
        return "I can speak Portuguese: Olá " + name + "!"
    default:
        return ""
    }
}
