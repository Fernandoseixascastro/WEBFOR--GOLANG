package handlers

import (
	"PageWeb/controller"
	"fmt"
	"net/http"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "erro ao fazer o parse do form: %v", err)

			return

		}

		err := controller.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))

		if err != nil {
			fmt.Fprintf(w, "erro ao criar o usuário: %v", err)

			return
		}
	}

	http.ServeFile(w, r, "handlers/templates/subscription.html")
}
