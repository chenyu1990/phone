package area

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"phone-area/app/http"
	"phone-area/schema"
	"regexp"
	"sync"
)

const API = "https://cx.shouji.360.cn/phonearea.php?number="

type Area struct {
	File string
}

type Data struct {
	Province        string `json:"province"`
	City            string `json:"city"`
	ServiceProvider string `json:"sp"`
}

type Response struct {
	Code int
	Data Data
}

func NewArea(file string) *Area {
	return &Area{
		File: file,
	}
}

func (a *Area) Run() error {
	phonesInfo, err := a.getPhones()
	if err != nil {
		return err
	}

	group := sync.WaitGroup{}
	for _, info := range phonesInfo {
		group.Add(1)
		go func(pi *schema.PhoneInfo) {
			_ = a.getInfo(pi)
			group.Done()
		}(info)
	}
	group.Wait()

	return nil
}

func (a *Area) getInfo(info *schema.PhoneInfo) error {

	body, err := http.Get(API + info.Number)
	if err != nil {
		return err
	}
	fmt.Printf("%s %v\n", info.Number, body)
	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	info.Province = res.Data.Province
	info.City = res.Data.City
	if info.City != "" {
		info.Area = fmt.Sprintf("%s省%s市", info.Province, info.City)
	} else {
		info.Area = fmt.Sprintf("%s市", info.Province)
	}
	info.ServiceProvider = res.Data.ServiceProvider

	return nil
}

func (a *Area) getPhones() (schema.PhoneInfos, error) {
	fi, err := os.Open(a.File)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	phoneInfos := schema.PhoneInfos{}

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		number := string(a)
		if number == "" {
			continue
		}

		compile := regexp.MustCompile("[0-9]{11}")
		subMatch := compile.FindStringSubmatch(number)
		for _, match := range subMatch {
			phoneInfo := &schema.PhoneInfo{
				Number: match,
			}
			phoneInfos = append(phoneInfos, phoneInfo)
		}
	}

	return phoneInfos, nil
}
