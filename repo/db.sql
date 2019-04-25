create table base (
    id int not null,
    Cpf int not null,
	Priv int null,
	Incompleto int null,
	DtUltCompra datetime null, 
	TicketMedio numeric(9,2),
	TicketUltCompra numeric(9,2),
	LojMaisFrequente varchar(18),
	LojUltCompra varchar(18),
    PRIMARY KEY (id)
)