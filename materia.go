package materia

import (
    "errors"
    "math/rand"
)

func AsignarCali(name string) ([]string, error) {
    if name == "" {
        return nil, errors.New("El nombre esta vacio")
    }

    var calie []string

    for i := 0; i < 4; i++ {
        calie = append(calie, randomCali())
    }

    return calie, nil
}

func NomCali(names []string) (map[string][]string, error) {
    cali := make(map[string][]string)

    for _, name := range names {
        califica, err := AsignarCali(name)

        if err != nil {
            return nil, err
        }

        cali[name] = califica
    }

    return cali, nil
}

func randomCali() string {
    options := []string{"70", "80", "90","100","NA"}
    return options[rand.Intn(len(options))]
}
