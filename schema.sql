CREATE TABLE IF NOT EXISTS department(
    department_id SERIAL PRIMARY KEY,
    department_name VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS employees(
    employee_id SERIAL PRIMARY KEY,
    employee_name VARCHAR(100) NOT NULL,
    department_id INTEGER references test7.department (department_id)
);