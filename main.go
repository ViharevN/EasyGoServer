package main

import (
	"flag"
	"fmt"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

// Реализация метода String для удовлетворения интерфейса flag.Value
func (na *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", na.Host, na.Port)
}

// Реализация метода Set для удовлетворения интерфейса flag.Value
func (na *NetAddress) Set(value string) error {
	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid address format: %s", value)
	}

	na.Host = parts[0]

	_, err := fmt.Sscanf(parts[1], "%d", &na.Port)
	if err != nil {
		return fmt.Errorf("invalid port format: %s", parts[1])
	}

	return nil
}

func main() {
	addr := new(NetAddress)
	// Если интерфейс не реализован,
	// будет ошибка компиляции
	_ = flag.Value(addr)

	// Проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()

	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
