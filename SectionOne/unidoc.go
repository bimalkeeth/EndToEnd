package main

import (
	"errors"
	"fmt"
	pdfcontent "github.com/unidoc/unipdf/contentstream"
	pdfcore "github.com/unidoc/unipdf/core"
	pdf "github.com/unidoc/unipdf/model"
	"os"
)

func main() {

	err := detectSignatureInput("/Users/bimalkeeth/Downloads/26.pdf")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func detectSignatureInput(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		if pageNum == 1 {
			found, x, y, err := locateSignatureLine(page)
			if err != nil {
				return err
			}
			// Happens if did not find the "___.." line or if there was no Tm position marker before.
			if !found || (x == 0 && y == 0) {
				return errors.New("Unable to find the signature line")
			}

			fmt.Printf("Position: x: %f, y: %f\n", x, y)
		}

	}

	return nil
}

func locateSignatureLine(page *pdf.PdfPage) (bool, float64, float64, error) {
	found := false
	x := float64(0)
	y := float64(0)

	pageContentStr, err := page.GetAllContentStreams()
	if err != nil {
		return found, x, y, err
	}

	cstreamParser := pdfcontent.NewContentStreamParser(pageContentStr)
	if err != nil {
		return found, x, y, err
	}

	operations, err := cstreamParser.Parse()
	if err != nil {
		return found, x, y, err
	}

	for _, op := range *operations {



		if op.Operand == "Tm" && len(op.Params) == 6 {
			if val, ok := op.Params[4].(*pdfcore.PdfObjectFloat); ok {
				x = float64(*val)
			}

			if val, ok := op.Params[5].(*pdfcore.PdfObjectFloat); ok {
				y = float64(*val)
			}
		} else if op.Operand == "Tj" && len(op.Params) == 1 {
			val, ok := op.Params[0].(*pdfcore.PdfObjectString)
			if ok {
				str := *val

				fmt.Println(str)
				//if strings.Contains(str, "________________") {
				//	fmt.Printf("Tj: %s\n", str)
				//	found = true
				//	break
				//}
			}
		}
	}

	return found, x, y, nil
}