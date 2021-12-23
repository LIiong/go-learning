package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tkrajina/gpxgo/gpx"
	"go-learning/31-http-request/polyline"
	"io/ioutil"
	"time"

	//"io/ioutil"
	"regexp"
	"strings"
)

var GARMIN_CN_URL_DICT = map[string]string{
	"BASE_URL":       "https://connect.garmin.cn",
	"SSO_URL_ORIGIN": "https://sso.garmin.com",
	"SSO_URL":        "https://sso.garmin.cn/sso",
	"MODERN_URL":     "https://connect.garmin.cn/modern",
	"SIGNIN_URL":     "https://sso.garmin.cn/sso/signin",
	"CSS_URL":        "https://connect.garmin.cn/gauth-custom-cn-v1.2-min.css",
}
var client = resty.New()

type runingPage struct {
	ActivityId   int64  `json:"activityId"`
	ActivityName string `json:"activityName"`
}
type Activitie struct {
	Id               int64   `json:"run_id"`
	Name             string  `json:"name"`
	Distance         float64 `json:"distance"`
	MovingTime       string  `json:"moving_time"`
	Type             string  `json:"type"`
	StartDate        string  `json:"start_date"`
	StartDateLocal   string  `json:"start_date_local"`
	LocationCountry  string  `json:"location_country"`
	SummaryPolyline  string  `json:"summary_polyline"`
	AverageHeartrate string  `json:"average_heartrate"`
	AverageSpeed     float64 `json:"average_speed"`
	Streak           int64   `json:"streak"`
}

func main() {
	result := login()
	setLogin(result)

	ids := getRunningIds(0, 2)()
	fmt.Println("activityIds :", ids)
	activities := make([]Activitie, 0)
	for _, id := range ids {
		url := fmt.Sprintf("%s/proxy/download-service/export/gpx/activity/%d", GARMIN_CN_URL_DICT["MODERN_URL"], id)
		fmt.Println("下载URL：" + url)
		get, _ := client.R().Get(url)
		//fileName := fmt.Sprintf("test-%d.gpx", id)
		//ioutil.WriteFile(fileName, get.Body(), 0644)
		gpxFile, err := gpx.ParseBytes(get.Body())
		if err != nil {
			panic(err)
		}
		summaryPolyline := processPolyline(gpxFile)

		startTime := gpxFile.TimeBounds().StartTime
		data := gpxFile.MovingData()
		cstZone := time.FixedZone("", 0)
		activitie := Activitie{
			Id:              id,
			Name:            "run from gpx",
			SummaryPolyline: summaryPolyline,
			Type:            "Run",
			StartDate:       startTime.In(time.Local).Format("2006-01-02 15:04:05"),
			StartDateLocal:  startTime.In(time.Local).Format("2006-01-02 15:04:05"),
			Distance:        data.MovingDistance,
			MovingTime:      time.Unix(int64(data.MovingTime), 0).In(cstZone).Format("15:04:05"),
			AverageSpeed:    data.MovingDistance / data.MovingTime,
			Streak:          1,
		}
		activities = append(activities, activitie)
	}
	//fmt.Println("---", activities)
	marshal, _ := json.Marshal(activities)
	ioutil.WriteFile("test.txt", marshal, 0644)
}

func processPolyline(file *gpx.GPX) string {
	points := make([][]float64, 0)
	for _, track := range file.Tracks {
		for _, segment := range track.Segments {
			for _, point := range segment.Points {
				points = append(points, []float64{point.Latitude, point.Longitude})
			}
		}
	}
	//coords, _, _ := polyline.DecodeCoords([]byte(summaryPolyline))
	//fmt.Println(coords)
	return string(polyline.EncodeCoords(points))
}

func getIds(result []runingPage) []int64 {
	var activityIds = make([]int64, 0)
	for _, page := range result {
		activityIds = append(activityIds, page.ActivityId)
	}
	return activityIds
}

func getRunningIds(start int, end int) func() []int64 {

	return func() []int64 {
		url := fmt.Sprintf("%s/proxy/activitylist-service/activities/search/activities?start=%d&limit=%d", GARMIN_CN_URL_DICT["MODERN_URL"], start, end)
		fmt.Println(url)
		response, _ := client.R().
			SetResult([]runingPage{}).
			Get(url)
		return getIds(*response.Result().(*[]runingPage))
	}
}

func setLogin(logins string) {
	r := regexp.MustCompile("(https:[^\"]+?ticket=[^\"]+)")
	s := r.FindString(logins)
	s = strings.ReplaceAll(s, "\\", "")
	fmt.Println("====>" + s)
	client.R().Get(s)
}

func login() string {

	resp, _ := client.R().
		SetHeaders(map[string]string{
			"origin":     "https://sso.garmin.cn",
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36",
		}).
		SetQueryParams(map[string]string{
			"webhost":                         GARMIN_CN_URL_DICT["MODERN_URL"],
			"service":                         GARMIN_CN_URL_DICT["MODERN_URL"],
			"source":                          GARMIN_CN_URL_DICT["SIGNIN_URL"],
			"redirectAfterAccountLoginUrl":    GARMIN_CN_URL_DICT["MODERN_URL"],
			"redirectAfterAccountCreationUrl": GARMIN_CN_URL_DICT["MODERN_URL"],
			"gauthHost":                       GARMIN_CN_URL_DICT["SSO_URL"],
			"locale":                          "zh_CN",
			"id":                              "gauth-widget",
			"cssUrl":                          GARMIN_CN_URL_DICT["CSS_URL"],
			"clientId":                        "GarminConnect",
			"rememberMeShown":                 "true",
			"rememberMeChecked":               "false",
			"createAccountShown":              "true",
			"openCreateAccount":               "false",
			"usernameShown":                   "false",
			"displayNameShown":                "false",
			"consumeServiceTicket":            "false",
			"initialFocus":                    "true",
			"embedWidget":                     "false",
			"generateExtraServiceTicket":      "false",
		}).
		SetFormData(map[string]string{
			"username":            "1121013687@qq.com",
			"password":            "Ljh7262556",
			"embed":               "true",
			"lt":                  "e1s1",
			"_eventId":            "submit",
			"displayNameRequired": "false",
		}).
		Post(GARMIN_CN_URL_DICT["SIGNIN_URL"])
	return resp.String()
}
