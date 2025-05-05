package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
type Floor struct {
	Content         [][]int
	FullContent     [][]int
	QuadtreeContent quadtree.Quadtree
	
	AlrRegistered   bool
}

// types d'affichage du terrain disponibles
const (
	GridFloor int = iota
	FromFileFloor
	QuadTreeFloor
)

// GetHeight retourne la hauteur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeight() (height int) {
	return len(f.FullContent)
}

// GetWidth retourne la largeur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidth() (width int) {
	if len(f.FullContent) > 0 {
		width = len(f.FullContent[0])
	}
	return
}
