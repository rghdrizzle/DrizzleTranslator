package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"rghdrizzle/drizzletranslator/helper"
	"rghdrizzle/drizzletranslator/response"

	//"strings"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func DetectLanguage(payload string)response.DetectResponse{
	client := resty.New()
	// payload := `{"q":"سَلَام"}`
	var detectResponse response.DetectResponse
	resp, err := client.R().
	SetHeader("x-rapidapi-key", "491d9b6e75msh35b56ae3d1c2b01p199039jsn53d6861791f8").
	SetHeader("x-rapidapi-host", "google-translator9.p.rapidapi.com").
	SetHeader("Content-Type", "application/json").
	SetBody(payload).
	Post("https://google-translator9.p.rapidapi.com/v2/detect")


	if err != nil {
        log.Fatal("Error:",err)
    }

	jsonerr := json.Unmarshal([]byte(resp.String()),&detectResponse)
	if jsonerr!=nil{
		log.Fatal("Error encoding",jsonerr)
	}
	
    return detectResponse

}

func Translate(c *fiber.Ctx)error{
	var reqBody helper.ReqBody
	var detectResponse response.DetectResponse

	client := resty.New()
	inputText := c.FormValue("source")
	targetLangSelected := "en"//c.FormValue("target")
	payloadToDetect := fmt.Sprintf("{\"q\":\"%s\"}", inputText)
	detectResponse = DetectLanguage(payloadToDetect)
	reqBody.SourceLang = detectResponse.Data.Detections[0][0].Language
	reqBody.TargetLang = targetLangSelected
	reqBody.SourceText = inputText
	fmt.Println("Source lang:"+reqBody.SourceLang)

	req := response.TranslateRequest{
		Q: reqBody.SourceText,
		Source: reqBody.SourceLang,
		Target: reqBody.TargetLang,
		Format: "text",
	}

	payload ,err:= json.Marshal(req)
	if err!=nil{
		log.Fatal("Error:",err)
	}

	resp,err := client.R().
	SetHeader("x-rapidapi-key", "491d9b6e75msh35b56ae3d1c2b01p199039jsn53d6861791f8").
	SetHeader("x-rapidapi-host", "google-translator9.p.rapidapi.com").
	SetHeader("Content-Type", "application/json").
	SetBody(payload).
	Post("https://google-translator9.p.rapidapi.com/v2")

	if err != nil {
        log.Fatal("Error:",err)
    }



	return c.JSON(resp.String())
	
}