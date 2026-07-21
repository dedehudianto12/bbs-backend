package auth

import (
	"encoding/json"
	"net/http"

	httphelper "github.com/dedehudianto12/bbs-backend/internal/shared/http"
)

type Handler struct {
	usecase *Usecase
}

func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
	Admin Admin  `json:"admin"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	token, admin, err := h.usecase.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		httphelper.Error(w, http.StatusUnauthorized, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // true di production
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400,
	})

	httphelper.Success(w, http.StatusOK, loginResponse{
		Token: token,
		Admin: *admin,
	})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	httphelper.Success(w, http.StatusOK, nil)
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	admin := r.Context().Value(adminCtxKey).(*Admin)
	httphelper.Success(w, http.StatusOK, admin)
}
