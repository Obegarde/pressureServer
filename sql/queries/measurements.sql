-- name: CreateMeasurement :one
INSERT INTO measurements(
id,
created_at,
measurement_date,
measurement_time, 
pressure_1,
pressure_2,
temperature_1,
temperature_2)
VALUES (
	gen_random_uuid(),
	NOW(),
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
	)
RETURNING *;

-- name: GetMeasurements :many
SELECT *
FROM measurements
ORDER BY measurement_date;

-- name: GetMeasurementsByDate :many
SELECT *
FROM measurements
WHERE measurement_date BETWEEN $1 AND $2;

-- name: CreateMeasurementsBatch :exec
INSERT INTO measurements(
	id,
	created_at,
	measurement_date,
	measurement_time,
	pressure_1,
	pressure_2,
	temperature_1,
	temperature_2)
VALUES(
	gen_random_uuid(),
	NOW(),
	unnest($1::DATE[]),
    	unnest($2::TIME[]),
    	unnest($3::FLOAT8[]),
    	unnest($4::FLOAT8[]),
    	unnest($5::FLOAT8[]),
    	unnest($6::FLOAT8[])
	);
