package grifts

import (
	"github.com/markbates/grift/grift"
	"github.com/balmanrawat/kopila-api/models"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here

		err := models.DB.TruncateAll()
		if err != nil {
			log.Fatal(err)
			return nil
		}
		
		csvFile, err := os.Open("grifts/data.csv")
		if err != nil {
			log.Fatal(err)
			return nil
		}
		
		reader := csv.NewReader(bufio.NewReader(csvFile))
		for {
			line, error := reader.Read()
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}
			fmt.Println(line)
			u := &models.User{
				Name: line[0],
				PhoneNumber: line[1],
				Tole: line[2],
				Ward: line[3],
			}

			fmt.Println(u)
			models.DB.Create(u)
		}
		return nil
	})

})
