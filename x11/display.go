package x11

// #cgo LDFLAGS: -lX11 -lXtst
// #include <X11/Xlib.h>
// #include <X11/keysym.h>
// #include <X11/keysymdef.h>
// #include <X11/extensions/XTest.h>
import "C"
import "errors"

var ErrOpeningDisplay = errors.New("Could not open display")

type Display interface {
	Press(key string)
	KeyEvent(key string, press bool)
}

func NewDisplay(identifier string) (Display, error) {
	disp := C.XOpenDisplay(C.CString(identifier))
	if disp == nil {
		return nil, ErrOpeningDisplay
	}

	return &display{
		id:       identifier,
		xDisplay: disp,
	}, nil
}

type display struct {
	id       string
	xDisplay *C.struct__XDisplay
}

func (d *display) Press(key string) {
	d.KeyEvent(key, true)
	d.KeyEvent(key, false)
}

func (d *display) KeyEvent(key string, press bool) {
	keycode := d.keyCode(key)
	var pressEvent C.int
	if press {
		pressEvent = C.int(1)
	}
	C.XTestFakeKeyEvent(d.xDisplay, keycode, pressEvent, 0)
	C.XFlush(d.xDisplay)
}

func (d *display) keyCode(key string) C.uint {
	char := C.CString(key)
	keysym := C.XStringToKeysym(char)
	keycode := C.XKeysymToKeycode(d.xDisplay, keysym)
	return C.uint(keycode)
}
