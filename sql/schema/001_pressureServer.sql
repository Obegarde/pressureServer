-- +goose Up
CREATE TABLE measurements(
	id UUID PRIMARY KEY NOT NULL,
	created_at TIMESTAMP NOT NULL,
	measurement_date DATE NOT NULL,
	measurement_time TIME NOT NULL,
	pressure_1 DECIMAL NOT NULL,
	pressure_2 DECIMAL NOT NULL,
	temperature_1 DECIMAL NOT NULL,
	temperature_2 DECIMAL NOT NULL
);

-- +goose Down
DROP TABLE measurements;
