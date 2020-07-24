package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	intpl := "./log/access%v.log"
	outpl := "./res/%v.log"
	n := 4
	for i := 0; i < n; i++ {
		fin, err := os.OpenFile(fmt.Sprintf(intpl, i+1), os.O_RDONLY, 0)
		if err != nil {
			panic(err)
		}
		fout, err := os.OpenFile(fmt.Sprintf(outpl, i+1), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			panic(err)
		}

		s := bufio.NewScanner(fin)
		for s.Scan() {
			err := s.Err()
			if err != nil {
				panic(err)
			}
			line := s.Text()
			if line == "" {
				continue
			}
			if match(line) {
				_, err := fout.Write([]byte(line + "\n"))
				if err != nil {
					panic(err)
				}
			}
		}
		if err := s.Err(); err != nil {
			panic(err)
		}

		fin.Close()
		fout.Close()
	}
}

func match(s string) bool {
	pid := "9dd401c2-2680-4922-8a24-2fe972eddbb4"
	fbid := "143694809403459"
	deviceDescr := `Dalvik/2.1.0 (Linux; U; Android 5.1.1; SM-J320H Build/LMY47V)`
	pidPos := strings.Index(s, pid)
	_ = pidPos
	fbidPos := strings.Index(s, fbid)
	_ = fbidPos
	_ = strings.Index(s, deviceDescr)
	devicePos := strings.Index(s, deviceDescr)
	_ = devicePos

	//return pidPos >= 0 || fbidPos >= 0
	//return fbidPos >= 0
	return devicePos >= 0
}
