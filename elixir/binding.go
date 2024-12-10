package elixir

//#include "parser.h"
//TSLanguage *tree_sitter_elixir();
import "C"
import (
	"unsafe"

	sitter "github.com/dougthor42/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_elixir())
	return sitter.NewLanguage(ptr)
}
