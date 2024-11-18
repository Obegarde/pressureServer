-- +goose Up
CREATE TABLE measurements(
	id UUID PRIMARY KEY NOT NULL,
	created_at TIMESTAMP NOT NULL,
	measurement_date DATE,
	measurement_time TIME,
	pressure_1 DECIMAL,
	pressure_2 DECIMAL,
	temperature_1 DECIMAL,
	temperature_2 DECIMAL
);

-- +goose Down
DROP TABLE measurements;
