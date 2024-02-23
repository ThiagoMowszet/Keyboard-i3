package main

import (
    "flag"
    "fmt"
    "os/exec"
)

func main() {
    layoutPtr := flag.String("layout", "us", "Diseño del teclado: es para español, us para inglés")

    flag.Parse()

    layout := *layoutPtr

    cmd := exec.Command("setxkbmap", layout)
    err := cmd.Run()

    if err != nil {
        fmt.Println("Error al ejecutar setxkbmap:", err)
        return
    }

    fmt.Printf("Configuración de teclado cambiada a %s\n", layout)
}
