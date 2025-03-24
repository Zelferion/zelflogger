package zelflogger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var RepeatedString = Style{}.Dim(strings.Repeat("-", 145))

func FatalLogger(a ...interface{}) {
	var log string
	TimeStamp := Color{}.Cyan(fmt.Sprintf("[%v]", GetTime()))
	Type := Style{}.Bold(Color{}.FatalRed("\"Fatal\""))
	log = fmt.Sprintf("%s %s - %s", TimeStamp, Type, Color{}.FatalRed(fmt.Sprint(a...)))
	fmt.Println(log)
}

func ErrorLogger(a ...interface{}) {
	var log string
	TimeStamp := Color{}.Cyan(fmt.Sprintf("[%v]", GetTime()))
	Type := Color{}.Red("\"Error\"")
	log = fmt.Sprintf("%s %s - %s", TimeStamp, Type, Color{}.Red(fmt.Sprint(a...)))
	fmt.Println(log)
}

func InfoLogger(a ...interface{}) {
	var log string
	TimeStamp := Color{}.Cyan(fmt.Sprintf("[%v]", GetTime()))
	Type := Color{}.BrightBlack("\"Info\"")
	log = fmt.Sprintf("%s %s - %s", TimeStamp, Type, Color{}.BrightBlack(fmt.Sprint(a...)))
	fmt.Println(log)
}


func WarnLogger(a ...interface{}) {
	var log string
	TimeStamp := Color{}.Cyan(fmt.Sprintf("[%v]", GetTime()))
	Type := Color{}.Yellow("\"Warn\"")
	log = fmt.Sprintf("%s %s - %s", TimeStamp, Type, Color{}.Yellow(fmt.Sprint(a...)))
	fmt.Println(log)
}

func MessageLogger(ype string, msg string) {
  Type := Color{}.BrightGreen(fmt.Sprintf("\"%v\"", ype))
  TimeStamp := Color{}.Cyan(fmt.Sprintf("[%v]" ,GetTime()))
  fmt.Printf("%s %s - %s\n", TimeStamp, Type, Color{}.BrightGreen(msg))
}

func GetTime() string {
	var timestamp string
	if Conf.TimestampsLogging {
		timestamp = time.Now().Format("2006/01/02 15:04:05")
	}
	return timestamp
}

func ReqUrl(r *http.Request) string {
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	} else if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		protocol = proto
	}

	fullURL := protocol + "://" + r.Host + r.URL.Path
	if r.URL.RawQuery != "" {
		fullURL += "?" + r.URL.RawQuery
	}
	return fullURL
}

func HeaderLogger(r *http.Request) {
	var Data string
	Header := r.Header
	user, password, ok := r.BasicAuth()
	if ok {
		Data = fmt.Sprintf("%s: %s %s: %s", Color{}.Cyan("username"), user, Color{}.Yellow("password"), password)
	}
	fmt.Print(Style{}.Dim(Color{}.Green(fmt.Sprintf("%s\n", Header))))
	if ok {
		fmt.Print(Style{}.Dim(fmt.Sprintf("%s\n", Data)))
	}
}

func BodyLogger(r *http.Request) {
	Body := r.Body
	defer Body.Close()
	bodyBytes, _ := io.ReadAll(Body)
	if string(bodyBytes) == "" {
		return
	}
	fmt.Print(Style{}.Dim(Color{}.Yellow("Request Body: ", string(bodyBytes))))
	fmt.Print("\n")
}

func TimeLogger(r *http.Request, Time time.Time) {
	timeTaken := Style{}.Italic(Color{}.Cyan(time.Since(Time)))
	fmt.Println(Style{}.Dim("Time taken to resolve the request of "), r.RemoteAddr, " : ", timeTaken)
}

func RequestLogger(r *http.Request) {
	Type := Color{}.Yellow(fmt.Sprintf("[%v]", "Incoming Request"))
	UserIp := Style{}.Italic(Color{}.Magenta(r.RemoteAddr))
	Method := Color{}.Blue(r.Method)
	URL := Color{}.Green(ReqUrl(r))
	DataLength := Color{}.Yellow(r.ContentLength)
	TimeStamp := Color{}.Cyan(GetTime())
	fmt.Printf("[%s] %s \"%s - %s\" from %s. %v bytes\n", TimeStamp, Type, URL, Method, UserIp, DataLength)
}


func ResolveJsonLog(b interface{}) error {
	if b == nil {
		return nil
	}
	if Conf.BodyLogging {
		jsonData, err := json.MarshalIndent(b, "", "  ")
		if err != nil {
			return err
		}

		fmt.Print(Style{}.Dim(fmt.Sprintf("Response Body: %s\n", string(jsonData))))
		return nil
	}
	return nil
}

func ResolveResp(t time.Time, r *http.Request) {
	if Conf.HttpRespLogging {
		Type := Color{}.Green(fmt.Sprintf("[%v]", "Resolved Request"))
		var TimeTaken string
		if Conf.TimestampsLogging {
			TimeTaken = Style{}.Italic(Color{}.Cyan(time.Since(t)))
		}
		UserIp := Style{}.Italic(Color{}.Magenta(r.RemoteAddr))
		TimeStamp := Color{}.Cyan(GetTime())
		fmt.Printf("[%s] %s from %s. %v\n", TimeStamp, Type, UserIp, TimeTaken)
		fmt.Printf("%s\n\n\n\n", RepeatedString)
	}
}

func ColorText(col string, symb string) string {
  return fmt.Sprintf("%v%s%v", col, symb, Reset)
}
