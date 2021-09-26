package meta

import (
	"errors"
	"math"
	"time"
)

func Meta(dataInicial, dataFinal time.Time, valorInicial, valorFinal, valorAtual float64) (percentualRealizado float64, err error) {
	if valorAtual == 0 {
		valorAtual = valorInicial
	}

	periodoDias := dataFinal.Sub(dataInicial).Hours() / 24
	if periodoDias <= 0 {
		err = errors.New("Datas inválidas")
	}

	crescimento := valorFinal / valorInicial
	if crescimento <= 0 {
		err = errors.New("Valores inválidas")
	}

	taxaCrescimentoDiaria := math.Pow(crescimento, float64(1/periodoDias))

	diasPercorridos := time.Now().Sub(dataInicial).Hours() / 24

	valorIdealHoje := valorInicial * math.Pow(taxaCrescimentoDiaria, diasPercorridos)

	percentualRealizado = valorInicial / valorIdealHoje
	return percentualRealizado, err
}
