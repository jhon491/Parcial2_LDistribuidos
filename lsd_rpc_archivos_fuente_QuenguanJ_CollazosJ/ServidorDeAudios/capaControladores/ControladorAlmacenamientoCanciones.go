package controlador

import (
	dtos "almacenamiento/capaFachadaServices/DTOs"
	capafachada "almacenamiento/capaFachadaServices/fachada"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var soloDigitos = regexp.MustCompile(`^[0-9]+$`)

type ControladorAlmacenamientoCanciones struct {
	fachada *capafachada.FachadaAlmacenamiento
}

// Constructor del controlador
func NuevoControladorAlmacenamientoCanciones() *ControladorAlmacenamientoCanciones {
	return &ControladorAlmacenamientoCanciones{
		fachada: capafachada.NuevaFachadaAlmacenamiento(),
	}
}

// AlmacenarAudioCancion
func (thisC *ControladorAlmacenamientoCanciones) AlmacenarAudioCancion(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Almacenando canción...\n")
	fmt.Printf("[REST] %s %s recibido desde %s\n", r.Method, r.URL.Path, r.RemoteAddr)

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(50 << 20) // Límite de 50MB

	file, _, err := r.FormFile("archivo")
	if err != nil {
		http.Error(w, "Error leyendo el archivo", http.StatusBadRequest)
		return
	}

	defer file.Close()
	data, _ := io.ReadAll(file)

	// Leer los campos del DTO
	dto := dtos.CancionAlmacenarDTOInput{
		IdTipo: r.FormValue("idTipo"),
		Id:     r.FormValue("id"),
	}

	if !soloDigitos.MatchString(dto.IdTipo) || !soloDigitos.MatchString(dto.Id) {
		http.Error(w, "idTipo e id son obligatorios y deben ser numéricos", http.StatusBadRequest)
		return
	}

	if err := thisC.fachada.GuardarCancion(dto, data); err != nil {
		fmt.Println("Error almacenando canción:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Canción almacenada: tipo=%s id=%s\n", dto.IdTipo, dto.Id)
}
