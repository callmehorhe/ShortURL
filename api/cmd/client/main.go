package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/callmehorhe/shorturl/api/pkg/handler"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() > 1 {
		log.Fatal("too many args")
	}

	log.Println("Enter \"create\\get <URL>\"")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	if len(line) != 2 {
		log.Fatal("wrong input")
	}
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := handler.NewURLClient(conn)
	switch strings.ToLower(line[0]) {
	case "create":
		res, err := c.Create(context.Background(), &handler.UrlMessage{Url: line[1]})
		if err != nil {
			log.Fatal(err)
		}
		log.Print(res.GetUrl())
	case "get":
		res, err := c.Get(context.Background(), &handler.UrlMessage{Url: line[1]})
		if err != nil {
			log.Fatal(err)
		}
		log.Print(res.GetUrl())
	default:
		log.Fatal("wrong input")
	}

}
