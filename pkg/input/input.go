package input


import (
		"fmt"
		"net/http"
		"os"
		"io"
	)

func UrlForDay(day int) string {
	return fmt.Sprintf("https://adventofcode.com/2020/day/%v", day)
}

func UrlForInput(day int) string {
 return	fmt.Sprintf("%v/input", UrlForDay(day))
}


func GetInput(session string, dir string, day int) error{
	transport := http.Transport {
		DisableCompression: false,
	}
	client := http.Client {
		Transport: &transport,
	}
	url := UrlForInput(day)
	data, err := client.Get(url)
	if err != nil {
		return err
	}
	out, err := os.Create(fmt.Sprintf("%v/input_day-0%v", dir, day))
	if err != nil {
		return err
	}
	defer out.Close()
	defer data.Body.Close()
	_, err1 := io.Copy(out, data.Body)
	return err1
}



