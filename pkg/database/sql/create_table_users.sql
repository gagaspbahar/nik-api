create table users (
	id text unique primary key,
	name text,
	district_id text,
	city_id text,
	province_id text,
	gender text,
	date_of_birth text,
	foreign key(district_id) references t_kecamatan(id),
	foreign key (city_id) references t_kota(id),
	foreign key (province_id) references t_provinsi(id)
);