package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"hill-cypher/cipher"
	"hill-cypher/util"
)

type App struct {
	keyFileLabel        *widget.Label
	keyFileBtn          *widget.Button
	selectedKeyFilePath *string

	textFileLabel        *widget.Label
	textFileBtn          *widget.Button
	selectedTextFilePath *string

	encryptBtn *widget.Button
	decryptBtn *widget.Button
}

func (a *App) checkFileSelected() bool {
	return a.selectedTextFilePath != nil && a.selectedKeyFilePath != nil
}

func (a *App) setSelectedKeyFilePath(keyFilePath *string) {
	a.selectedKeyFilePath = keyFilePath
	if a.selectedKeyFilePath != nil {
		a.keyFileLabel.SetText(*a.selectedKeyFilePath)
	} else {
		a.keyFileLabel.SetText("No key file selected")
	}

	a.updateBtnState()
}

func (a *App) setSelectedTextFilePath(textFilePath *string) {
	a.selectedTextFilePath = textFilePath
	if a.selectedTextFilePath != nil {
		a.textFileLabel.SetText(*a.selectedTextFilePath)
	} else {
		a.textFileLabel.SetText("No text file selected")
	}

	a.updateBtnState()
}

func (a *App) updateBtnState() {
	selected := a.checkFileSelected()
	if selected {
		a.decryptBtn.Enable()
		a.encryptBtn.Enable()
		return
	}

	a.decryptBtn.Disable()
	a.encryptBtn.Disable()
}

func (a *App) getFilesContent() ([][]int, *string, error) {
	keyStr, keyErr := util.ReadFileString(*a.selectedKeyFilePath)
	if keyErr != nil {
		return nil, nil, keyErr
	}

	key, keyParseError := util.ParseKey(keyStr)
	if keyParseError != nil {
		return nil, nil, keyParseError
	}

	textStr, textErr := util.ReadFileString(*a.selectedTextFilePath)
	if textErr != nil {
		return nil, nil, textErr
	}

	return key, textStr, nil
}

func ShowGui() {
	myApp := App{}
	a := app.New()
	w := a.NewWindow("Hello")

	myApp.textFileLabel = widget.NewLabel("No text file selected")
	myApp.keyFileLabel = widget.NewLabel("No key file selected")

	encryptBtn := widget.NewButton("Encrypt", func() {
		fmt.Println("Encrypting...")
		key, content, err := myApp.getFilesContent()
		if err != nil {
			fmt.Println(err)
			dialog.ShowError(err, w)
			return
		}
		encrypted := cipher.Encrypt(key, *content)
		saveErr := util.SaveFile("./encrypted_output.txt", encrypted)
		if saveErr != nil {
			dialog.ShowError(err, w)
			return
		}
	})
	encryptBtn.Disable()
	myApp.encryptBtn = encryptBtn

	decryptBtn := widget.NewButton("Decrypt", func() {
		fmt.Println("Decrypting...")
		key, content, err := myApp.getFilesContent()
		if err != nil {
			fmt.Println(err)
			dialog.ShowError(err, w)
			return
		}
		encrypted := cipher.Decrypt(key, *content)
		saveErr := util.SaveFile("./decrypted_output.txt", encrypted)
		if saveErr != nil {
			dialog.ShowError(err, w)
			return
		}
	})
	decryptBtn.Disable()
	myApp.decryptBtn = decryptBtn

	w.SetContent(container.NewVBox(
		myApp.textFileLabel,
		widget.NewButton("Select text file", func() {
			d := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, w)
					myApp.setSelectedTextFilePath(nil)
					return
				}
				if reader == nil {
					myApp.setSelectedTextFilePath(nil)
					return
				}
				path := reader.URI().Path()
				myApp.setSelectedTextFilePath(&path)

				defer reader.Close()
			}, w)
			d.Show()
		}),

		myApp.keyFileLabel,
		widget.NewButton("Select key file", func() {
			d := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, w)
					myApp.setSelectedKeyFilePath(nil)
					return
				}
				if reader == nil {
					myApp.setSelectedKeyFilePath(nil)
					return
				}

				path := reader.URI().Path()
				myApp.setSelectedKeyFilePath(&path)

				defer reader.Close()
			}, w)
			d.Show()
		}),

		encryptBtn,
		decryptBtn,
	))

	w.ShowAndRun()
}
