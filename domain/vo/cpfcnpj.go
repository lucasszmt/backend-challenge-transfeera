package vo

import (
	"bytes"
	"regexp"
	"unicode"
)

var (
	CPFRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

type DocType string

const (
	CPF  DocType = "cpf"
	CNPJ DocType = "cnpj"
)

type CpfCnpj struct {
	value   string
	docType DocType
}

// TODO arrumar, está aceitando criar uma chave CNPJ on de é passado um cpf
func NewCpfCnpj(doc string) (*CpfCnpj, error) {
	//TODO improve CPF/CNPJ validation with sum of digits
	cpfcpnj := new(CpfCnpj)
	if CPFRegexp.MatchString(doc) {
		cpfcpnj.SetValue(doc)
		cpfcpnj.SetType(CPF)
		return cpfcpnj, nil
	}
	if CNPJRegexp.MatchString(doc) {
		cpfcpnj.SetValue(doc)
		cpfcpnj.SetType(CNPJ)
		return cpfcpnj, nil
	}
	return nil, ErrInvalidCPFCNPJ
}

func (c *CpfCnpj) GetType() DocType {
	return c.docType
}

func (c *CpfCnpj) GetValue() string {
	return c.value
}

func (c *CpfCnpj) SetType(docType DocType) {
	c.docType = docType
}

func (c *CpfCnpj) SetValue(doc string) {
	c.value = removeNonDigits(doc)
}

func removeNonDigits(doc string) string {
	buf := bytes.NewBufferString("")
	for _, r := range doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
