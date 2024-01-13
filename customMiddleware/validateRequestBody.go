package customMiddleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ValidateRequestBody[T interface{}]() func(http.Handler) http.Handler {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var unmarshalTypeError *json.UnmarshalTypeError

			r.Body = http.MaxBytesReader(w, r.Body, 1048576)
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			defer r.Body.Close()

			var b T

			if err := dec.Decode(&b); err != nil {
				fmt.Println(err)
				msg := fmt.Sprintf(
					"Unexpected error (decoding request body): %s",
					err.Error(),
				)
				if errors.As(err, &unmarshalTypeError) {
					msg = fmt.Sprintf(
						"The %q field value is invalid",
						unmarshalTypeError.Field,
					)
					http.Error(w, msg, http.StatusBadRequest)
					return
				}
				if strings.HasPrefix(err.Error(), "json: unknown field ") {
					fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
					msg = fmt.Sprintf(
						"Unknown field %s in request body",
						fieldName,
					)
					http.Error(w, msg, http.StatusBadRequest)
					return
				}
			}

			err := dec.Decode(&struct{}{})
			if !errors.Is(err, io.EOF) {
				msg := "Request body must only contain a single JSON object"
				http.Error(w, msg, http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(r.Context(), "body", b)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
