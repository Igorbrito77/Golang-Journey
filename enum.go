package main;

import(
	"fmt"
);

type Status uint8;

const (
	Created Status = iota
	Pending
	Approved
	Rejected
  );

  const (
	Barbaro =  "Bárbaro"
	Arqueiro = "Arqueiro"
  );


func main(){



	fmt.Println(Rejected);

	fmt.Println(Barbaro);


}