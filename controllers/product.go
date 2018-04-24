package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	// "google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
	"om/models"
	"om/utilities"
)

func MapProductUrl(w http.ResponseWriter, r *http.Request) utilities.ResponseJSON {

	returnData := utilities.ResponseJSON{}
	returnData.Msg = "Failure"
	returnData.Code = 400
	returnData.Model = nil
	ID := r.URL.Query().Get("id")
	inputobj := models.Product{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &inputobj)

	switch true {
	case cast.ToString(r.Method) == "GET" && cast.ToString(ID) == "":
		allData, err := models.GetAllProduct()
		if err == nil {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = allData
		} else {
			returnData.Code = 401
			returnData.Msg = "Failure:PRODUCT GET REQUEST"
		}
		break

	case cast.ToString(r.Method) == "GET" && cast.ToString(ID) != "":
		allData, err := models.GetProductById(cast.ToInt(ID))
		if err == nil {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = allData
		} else {
			returnData.Code = 402
			returnData.Msg = "Failure:PRODUCT doesn't exists"
		}
		break

	case cast.ToString(r.Method) == "POST":
		err := models.UpdateProductById(&inputobj)
		if err == nil {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = nil
		} else {
			returnData.Code = 403
			returnData.Msg = "Failure:PRODUCT  Update error"
		}
		break

	case cast.ToString(r.Method) == "PUT":
		_, err := models.AddProduct(&inputobj)
		if err == nil {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = nil
		} else {
			returnData.Code = 404
			returnData.Msg = "Failure:PRODUCT Creation error"
		}
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure:Authorization error"
		returnData.Model = nil

	}
	return returnData
}
