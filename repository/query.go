package repository

const (
	Insert = `insert into rfb_data (
		cnpj,
		cnpj_basico,
		cnpj_ordem,
		cnpj_dv,
		identificador,
		nome_fantasia,
		situacao_cadastral,
		data_situacao_cadastral,
		motivo_situacao_cadastral,
		nome_cidade_exterior,
		pais,
		data_inicio,
		cnae_principal,
		cnae_secundario,
		tipo_logradouro,
		logradouro,
		numero,
		complemento,
		bairro,
		cep,
		uf,
		municipio,
		ddd_1,
		telefone_1,
		ddd_2,
		telefone_2,
		ddd_fax,
		fax,
		email,
		situacao_especial,
		data_situacao_especial,
		create_at
	) values (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13,
		$14,
		$15,
		$16,
		$17,
		$18,
		$19,
		$20,
		$21,
		$22,
		$23,
		$24,
		$25,
		$26,
		$27,
		$28,
		$29,
		$30,
		$31,
		$32
	) 
	on conflict (cnpj)
	do update set 
		cnpj = $1,
		cnpj_basico = $2,
		cnpj_ordem = $3,
		cnpj_dv = $4,
		identificador = $5,
		nome_fantasia = $6,
		situacao_cadastral = $7,
		data_situacao_cadastral = $8,
		motivo_situacao_cadastral = $9,
		nome_cidade_exterior = $10,
		pais = $11,
		data_inicio = $12,
		cnae_principal = $13,
		cnae_secundario = $14,
		tipo_logradouro = $15,
		logradouro = $16,
		numero = $17,
		complemento = $18,
		bairro = $19,
		cep = $20,
		uf = $21,
		municipio = $22,
		ddd_1 = $23,
		telefone_1 = $24,
		ddd_2 = $25,
		telefone_2 = $26,
		ddd_fax = $27,
		fax = $28,
		email = $29,
		situacao_especial = $30,
		data_situacao_especial = $31,
		update_at = $32`

	InsertError = `insert into errors (
		cnpj,
		field,
		value,
		error,
		file,
		create_at
	) values (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)`	
)
