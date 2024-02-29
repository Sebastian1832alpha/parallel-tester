func main() {
	// Create a sub-directory filesystem from the embedded files
	subFS, err := fs.Sub(staticAssets, "dist")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	domainsQuery := r.URL.Query().Get("teritory")
	if domainsQuery == "" {
		http.Error(w, "No domains", http.StatusBadRequest)
		return
	}
