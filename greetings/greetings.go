package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func Hello(name string) (string, error) {

    if name == "" {
        return "", errors.New("empty name")
    }


    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    //making a map in Go   make(map[key-type]value-type)
    messages := make(map[string]string)

    // for ind, val := range array  , to use for loop on arrays in Go
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}


//Go calls init automatically AFTER global vars initialized
func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    //Go's array omitting a size, a "slice", for dynamically sized arrays
    formats := []string {
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    return formats[rand.Intn(len(formats))]
}
