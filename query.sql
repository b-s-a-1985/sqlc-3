-- name: CreateSchema :exec
CREATE SCHEMA IF NOT EXISTS test8 AUTHORIZATION postgres;

-- name: SetSchema :exec
SET search_path TO test8;

-- name: CreateTable :exec
CREATE TABLE IF NOT EXISTS department(
    department_id SERIAL PRIMARY KEY,
    department_name VARCHAR(20) NOT NULL
);

-- name: InsertDepartment :exec
INSERT INTO department (department_name)
VALUES ($1);

-- name: ListDepartments :many
SELECT * FROM department;