package web

import (
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func RunStep(client selenium.WebDriver, webCase WebCase) {
	log.Println("------------------------------------------")
	log.Println("caseName =  " + webCase.Name)
	stepBody := webCase.Check
	for _, checkBody := range stepBody {
		// collec, ctx := utils.MongoDb()
		switch checkBody.CheckTpye {
		case "find":
			for _, checkEle := range checkBody.CheckBody {
				log.Println("find----" + checkEle)
				_, err := client.FindElement(selenium.ByCSSSelector, checkEle)
				if err != nil {
					log.Println("find----" + checkEle + "----error")
					// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody, "step": "checkEle", "stepResult": "error"})
				} else {
					log.Println("find-----" + checkEle + "-----success")

					// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody, "step": "checkEle", "stepResult": "succcess"})
				}
			}
		case "click":
			for _, checkEle := range checkBody.CheckBody {
				log.Println("click-----" + checkEle)
				ele, err := client.FindElement(selenium.ByCSSSelector, checkEle)
				if err != nil {
					log.Println("find-----" + checkEle + "-----error")
					// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody, "step": "checkClickEle", "stepResult": "error"})
				} else {
					log.Println("find-----" + checkEle + "-----success")

					// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody, "step": "checkClickEle", "stepResult": "success"})
					err = ele.Click()
					if err != nil {
						log.Println("input-----" + checkEle + "-----error")
						// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody, "step": "clickAction", "stepResult": "error"})
					} else {
						log.Println("input-----" + checkEle + "-----success")
						// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody, "step": "clickAction", "stepResult": "success"})
					}
				}

			}
		case "input":
			for i := 0; i < len(checkBody.CheckBody); i += 2 {
				log.Println("input-----" + checkBody.CheckBody[i] + checkBody.CheckBody[i+1])
				ele, err := client.FindElement(selenium.ByCSSSelector, checkBody.CheckBody[i])
				if err != nil {
					log.Println("find-----" + checkBody.CheckBody[i] + "-----error")
					// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody[i], "step": "findInputEle", "stepResult": "error"})
				} else {
					log.Println("find-----" + checkBody.CheckBody[i] + "-----success")
					// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody, "step": "findInputEle", "stepResult": "success"})
					err := ele.SendKeys(checkBody.CheckBody[i+1])
					if err != nil {
						log.Println("input-----" + checkBody.CheckBody[i+1] + "-----error")
						// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody[i+1], "step": "InputData", "stepResult": "error"})
					} else {
						log.Println("input-----" + checkBody.CheckBody[i+1] + "-----success")
						// collec.InsertOne(*ctx, bson.M{"caseName": webCase.Name, "caseType": webCase.Url, "checkType": checkBody.CheckTpye, "checkBody": checkBody.CheckBody[i+1], "step": "InputData", "stepResult": "success"})
					}
				}

			}
		case "wait":
			time.Sleep(time.Second * time.Duration(5))
		default:
			log.Println("timeout -----err")
		}
	}
}
