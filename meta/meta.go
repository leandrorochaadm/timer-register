package meta

import (
	"errors"
	"math"
	"time"
)

func Meta(dataInicial, dataFinal, dataAfericao time.Time, valorInicial, valorFinal, valorAtual float64) (percentualRealizadoDesejado, percentualRealizadoTempo, percentualRealizadoDoValor float64, err error) {
	if dataAfericao.IsZero() {
		dataAfericao = time.Now()
	}
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

	diasPercorridos := dataAfericao.Sub(dataInicial).Hours() / 24

	valorIdealHoje := valorInicial * math.Pow(taxaCrescimentoDiaria, diasPercorridos)

	percentualRealizadoDesejado = (valorAtual / valorIdealHoje) * 100

	diasMeta := dataFinal.Sub(dataInicial).Hours() / 24

	percentualRealizadoTempo = (diasPercorridos / diasMeta) * 100

	percentualRealizadoDoValor = (valorAtual / valorFinal) * 100
	return
}
