package orders

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type OrderHttpEndpoints interface {
	CreateOrder() func(w http.ResponseWriter, r *http.Request)
	ListOrder() func(w http.ResponseWriter, r *http.Request)
}

type orderHttpEndpoints struct {
	store OrdersStore
}

func NewOrderHttpEndpoints(store OrdersStore) OrderHttpEndpoints {
	return &orderHttpEndpoints{store: store}
}

func (o *orderHttpEndpoints) CreateOrder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		contextData := r.Context().Value("user_id")
		userId := contextData.(string)
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		cmd := &CreateOrderCommand{}
		err = json.Unmarshal(jsonData, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		cmd.UserId = userId
		newOrder, err := o.store.Create(&Order{
			ProductId: cmd.ProductId,
			UserId:    cmd.UserId,
		})
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusCreated, newOrder)
		return
	}
}

func (o *orderHttpEndpoints) ListOrder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		contextData := r.Context().Value("user_id")
		userId := contextData.(string)
		orders, err := o.store.ListOrdersByUserId(userId)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusCreated, orders)
		return
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
