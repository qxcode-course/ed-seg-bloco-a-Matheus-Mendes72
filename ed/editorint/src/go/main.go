package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	textos  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.textos.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move para a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	if e.cursor == e.line.Value.End() {
		novaL := NewList[rune]()
		e.textos.Insert(e.line.Next(), novaL)
		e.line = e.line.Next()
		e.cursor = e.line.Value.Front()
		return
	}

	novaL := NewList[rune]()
	for it := e.cursor; it != e.line.Value.End(); {
		prox := it.Next()

		novaL.PushBack(it.Value)
		e.line.Value.Erase(it)

		it = prox
	}

	e.textos.Insert(e.line.Next(), novaL)

	e.line = e.line.Next()
	e.cursor = novaL.Front()
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.cursor.Next()
		return
	}

	if e.line != e.textos.Back() {
        e.line = e.line.Next()
        e.cursor = e.line.Value.Front()
    }
}

func (e *Editor) KeyUp() {
	if e.line == e.textos.Front() {
		return
	}

	pos := e.line.Value.IndexOf(e.cursor)
	e.line = e.line.Prev()
	inicio := e.line.Value.Front()

	for i := 0; i < pos; i++ {
		inicio = inicio.Next()
	}

	e.cursor = inicio
}

func (e *Editor) KeyDown() {
	if e.line == e.textos.Back() {
		return
	}

	pos := e.line.Value.IndexOf(e.cursor)
	e.line = e.line.Next()
	inicio := e.line.Value.Front()

	for i := 0; i < pos; i++ {
		inicio = inicio.Next()
	}

	e.cursor = inicio
}

func (e *Editor) KeyBackspace() {
	if e.cursor != e.line.Value.Front() {
		e.line.Value.Erase(e.cursor.Prev())
		return
	}

	if e.line != e.textos.Front() {
		nextL := e.line.Prev()

		posF := nextL.Value.End()
		pos := e.line.Value.Front()

		for pos != e.line.Value.End() {
			next := pos.Next()
			nextL.Value.Insert(posF, pos.Value)
			e.line.Value.Erase(pos)
			pos = next
		}

		e.textos.Erase(e.line)
		e.line = nextL
		e.cursor = posF
	}
}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.End() {
		e.line.Value.Erase(e.cursor.Next())
		return
	}

	if e.line != e.textos.Back() {
		nextL := e.line.Next()

		for pos := nextL.Value.Front(); pos != nextL.Value.End(); {
			next := pos.Next()
			e.line.Value.Insert(e.line.Value.End(), pos.Value)
			nextL.Value.Erase(pos)

			pos = next
		}

		e.textos.Erase(nextL)
	}
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.textos = NewList[*List[rune]]()
	e.textos.PushBack(NewList[rune]())
	e.line = e.textos.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.textos.Front(); line != e.textos.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
