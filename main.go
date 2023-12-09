package main

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func main() {
	fmt.Println(timestamppb.New(time.Date(2023, time.June, 26, 00, 00, 00, 00, time.UTC)))
}
