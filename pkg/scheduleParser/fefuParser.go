package scheduleParser

import (
	"encoding/json"
	"homework_bot/internal/domain"
	"io/ioutil"
	"net/http"
)

type FefuParser struct {
}

func NewFefuParser() *FefuParser {
	return &FefuParser{}
}

func ByteToSchedule(body []byte) (domain.Schedule, error) {
	var schedule domain.Schedule
	err := json.Unmarshal(body, &schedule)
	return schedule, err
}

func (p *FefuParser) ParseSchedule(link string) (domain.Schedule, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return domain.Schedule{}, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.7.0.0 Safari/537.36")
	req.Header.Add("Referer", "https://univer.dvfu.ru/schedule")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", "_ym_uid=1725560544711277599; _ym_d=1725560544; enter=true; __ddg1_=DW6OiImjGIwhRgMg18H1; tmr_lvid=91dffdce1bc8fe16cb17f1e2928582c4; tmr_lvidTS=1726746541398; _univer_identity=f411d5c8da71239c0f23d3e3b36cab0be33dcb32a8758116ff760938a76ce228a%3A2%3A%7Bi%3A0%3Bs%3A16%3A%22_univer_identity%22%3Bi%3A1%3Bs%3A50%3A%22%5B50348%2C%22tIb4_L3cgmRQm-2yqk4He78sJcbDmzIc%22%2C2592000%5D%22%3B%7D; _jwts=cdfebf9d65c4865f324df04b5e2b6920f69181404b0e4c3a6017830ac00fad22a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_jwts%22%3Bi%3A1%3Bs%3A199%3A%22eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE3MjU0NDY3NDEsImV4cCI6MTcyODAzODc0MSwidWlkIjo1MDM0OCwidXNlcm5hbWUiOiJrZXJ6aGFrb3YubmEiLCJyb2xlIjoiYWRtaW4ifQ.B7cK7_1fsxOtEOOfWkULFNOC8la8kdLbCJWvle9Q5Lo%22%3B%7D; LtpaToken2=+h5qReQeJD0HnxJ1aLi0t2uXxDl0qoZR0N8LoINiqOWOYEbdC9ixDQM86FL9h4mT7QkLTRNDouzHQjWfLvbCCHrNi23QK4p6MBPOpnfedlM0rAutkbjf50zAj3ko8gaPC8PSDnA4FU3UL4AE2GgdpxIiCcKLeT6wVSi1Brxx68xz90xSpUkNXd6gOA1YtybtN9AwDKhSANcT7S4osE7BJw74IHKxs58sOTD9I42TLQ9cPZ32mSHBxobqcCK6kFXulWd/Z/sXzY9D+W8WNEUyCDze6eKTSc2bguuODUrmpCrUcYcVUwD4+eJxt06StqO9B52T8Ub2Bss7sOpWY6zFIkCCnx6MSY8kmgY+f8f12pQJBMW3DfMtYWv8KjPbK3h1wNr42gy6vAD45n7t6MjI/g==; _ym_isad=1; session-cookie=17f8cd02162c257bb1b9a252b68b8c5b8d780692c23b471f280b932be03606b05f7e1b6ba05330e23ce8499761f323e8; _csrf_univer=afb9b74c94e71cde13c046337b05e315e32fc92c630f012e129937c0df69bb95a%3A2%3A%7Bi%3A0%3Bs%3A12%3A%22_csrf_univer%22%3Bi%3A1%3Bs%3A32%3A%22xSUNwnsf-0yl02nDugvseEQFDcAQUk1v%22%3B%7D; _ym_visorc=w; _univer_session=d5thfu16i9nmvner2jgpavj2ve")
	resp, err := client.Do(req)
	if err != nil {
		return domain.Schedule{}, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.Schedule{}, err
	}

	return ByteToSchedule(body)
}
