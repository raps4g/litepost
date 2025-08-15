package ui

import "github.com/rivo/tview"

type FocusLayout struct {
    Matrix  [][]tview.Primitive
    X, Y    int
    OnFocus tview.Primitive
    OpenDialog tview.Primitive
}

func (f *FocusLayout) MoveLeft() {
    if f.Y > 0 {
        f.Y--
        f.OnFocus = f.Matrix[f.X][f.Y]
    }
}

func (f *FocusLayout) MoveRight() {
    if f.X < len(f.Matrix) && f.Y < len(f.Matrix[f.X])-1 {
        f.Y++
        f.OnFocus = f.Matrix[f.X][f.Y]
    }
}

func (f *FocusLayout) MoveUp() {
    if f.X > 0 && f.Y < len(f.Matrix[f.X-1]) {
        f.X--
        f.OnFocus = f.Matrix[f.X][f.Y]
    }
}

func (f *FocusLayout) MoveDown() {
    if f.X < len(f.Matrix)-1 && f.Y < len(f.Matrix[f.X+1]) {
        f.X++
        f.OnFocus = f.Matrix[f.X][f.Y]
    }
}

