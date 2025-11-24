package handlers

func (h *Handler) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, "cannot get notes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notes)
}

func (h *Handler) GetNoteById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	note, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(note)
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := h.Service.Create(&note); err != nil {
		http.Error(w, "cannot create note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)

	if err := h.Service.Update(&note); err != nil {
		http.Error(w, "cannot update", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(note)
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "cannot delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
