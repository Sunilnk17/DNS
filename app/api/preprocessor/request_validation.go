package preprocessor

import (
	"context"
	"drone-navigation-service-master/app/config"
	"encoding/json"
	"github.com/leebenson/conform"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

func DecodeAndValidateRequestParams(r *http.Request, structPointer interface{}) error {
	err := BindRequest(r, structPointer)

	if err != nil {
		log.Print(err)
		return err
	}
	log.Print(r.Context(), "Decoded Params: ", structPointer)
	err = conform.Strings(structPointer) //Sanitizes input. Removes extra space or converts to uppercase based on the struct tags.
	if err != nil {
		log.Print(r.Context(), "Error: ", err.Error())
		return err
	}

	err = validateParams(r.Context(), structPointer)

	log.Print(r.Context(), "Validated Params: ", structPointer)
	return err
}

func validateParams(ctx context.Context, structPointer interface{}) error {
	paramsValidationErr := config.GetReqParamsValidator().Struct(structPointer)
	if paramsValidationErr != nil {
		//TODO: construct response for invalid error message

		if _, ok := paramsValidationErr.(*validator.InvalidValidationError); ok {
			log.Fatal(ctx, paramsValidationErr)
			return paramsValidationErr
		}
	}
	return paramsValidationErr

}

func BindRequest(r *http.Request, structPointer interface{}) error {
	var err error
	err = r.ParseForm()
	if err != nil {
		return err
	}
	var decoderError error
	if r.Method == http.MethodGet {
		decoderError = config.GetReqParamsDecoder().Decode(structPointer, r.Form)
	}
	return decoderError
}

func BindAndUnMarshallRequest(r *http.Request, structPointer interface{}, validationPointer interface{}) ([]byte, error) {

	err := BindRequest(r, structPointer)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(structPointer)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, validationPointer)

	return bytes, err
}
