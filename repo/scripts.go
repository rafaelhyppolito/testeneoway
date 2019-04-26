package repo

//Funcao com script SQL para inserção final dos dados já higienizados
func InsertFinal() string {
	return "insert into base(cpf,priv,incompleto,dtultcompra,ticketmedio,ticketultcompra,lojmaisfrequente,lojultcompra) select replace(replace(cpf,'.',''),'-',''), CAST(priv AS INTEGER), CAST(incompleto AS INTEGER), CASE WHEN dtultcompra = 'NULL' THEN NULL ELSE CAST(dtultcompra AS DATE) END AS dtultcompra, CASE WHEN ticketmedio = 'NULL' THEN NULL ELSE CAST(ticketmedio AS NUMERIC(9,2)) END AS ticketmedio, CASE WHEN ticketultcompra = 'NULL' THEN NULL ELSE CAST(ticketultcompra AS NUMERIC(9,2)) END AS ticketultcompra, CASE WHEN lojmaisfrequente = 'NULL' THEN NULL ELSE replace(replace(replace(lojmaisfrequente,'.',''),'-',''),'/','') END AS lojmaisfrequente, CASE WHEN lojultcompra = 'NULL' THEN NULL ELSE replace(replace(replace(lojultcompra,'.',''),'-',''),'/','') END AS lojultcompra from basetmp"
}

//Funcao com srcipt SQL para limpar a tabela temporária
func TruncateTmp() string{
	return "TRUNCATE TABLE basetmp"
}