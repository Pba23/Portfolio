package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    err := tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    // Serve les fichiers statiques
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // Page d'accueil
    http.HandleFunc("/", homeHandler)
    
    // Récupère le port depuis la variable d'environnement PORT
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Port par défaut pour le développement local
    }
    
    fmt.Printf("Serveur en cours sur le port %s\n", port)
    http.ListenAndServe(":"+port, nil)
}