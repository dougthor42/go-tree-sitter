package kotlin

//#include "parser.h"
//TSLanguage *tree_sitter_kotlin();
import "C"
import (
	"unsafe"

	sitter "github.com/dougthor42/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_kotlin())
	return sitter.NewLanguage(ptr)
}
