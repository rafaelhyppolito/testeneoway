create table base (
    id SERIAL PRIMARY KEY,
    Cpf varchar(14) not null,
	Priv int null,
	Incompleto int null,
	DtUltCompra date null, 
	TicketMedio numeric(9,2),
	TicketUltCompra numeric(9,2),
	LojMaisFrequente varchar(18),
	LojUltCompra varchar(18)
)