package xpath

/*
#cgo pkg-config: libxml-2.0
#include <libxml/xpath.h> 
#include <libxml/xpathInternals.h>
*/
import "C"
import "unsafe"
import . "github.com/moovweb/gokogiri/util"
import "runtime"

type Expression struct {
	Ptr *C.xmlXPathCompExpr
	xpath string
}

func Compile(path string) (expr *Expression) {
	if len(path) == 0 {
		return
	}

	xpathBytes := AppendCStringTerminator([]byte(path))
	xpathPtr := unsafe.Pointer(&xpathBytes[0])
	ptr := C.xmlXPathCompile((*C.xmlChar)(xpathPtr))
	if ptr == nil {
		return
	}
	expr = &Expression{Ptr: ptr, xpath: path}
	runtime.SetFinalizer(expr, (*Expression).Free)
	return
}

func (exp *Expression) String() string {
	return exp.xpath
}

func (exp *Expression) Free() {
	if exp.Ptr != nil {
		C.xmlXPathFreeCompExpr(exp.Ptr)
		exp.Ptr = nil
	}
}
